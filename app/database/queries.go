package database

import (
	"augustin/structs"
	"context"

	"go.uber.org/zap"
)

// GetHelloWorld returns the string "Hello, world!" from the database and should be used as a template for other queries
func (db *Database) GetHelloWorld() (string, error) {
	var greeting string
	err := db.Dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Error("QueryRow failed: %v\n", zap.Error(err))
		return "", err
	}
	return greeting, err
}

// GetPayments returns the payments from the database
func (db *Database) GetPayments() ([]structs.Payment, error) {
	var payments []structs.Payment
	rows, err := db.Dbpool.Query(context.Background(), "select * from payment")
	if err != nil {
		log.Error("GetPayments failed", zap.Error(err))
		return payments, err
	}
	for rows.Next() {
		var payment structs.Payment
		err = rows.Scan(&payment.ID, &payment.Timestamp, &payment.Sender, &payment.Receiver, &payment.Type, &payment.Amount, &payment.AuthorizedBy, &payment.Item, &payment.PaymentBatch)
		if err != nil {
			return payments, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

// Create payment type
func (db *Database) CreatePaymentType(pt structs.PaymentType) (id int32, err error) {
	err = db.Dbpool.QueryRow(context.Background(), "insert into PaymentType (Name) values ($1) RETURNING ID", pt.Name).Scan(&id)
	return id, err
}

// Create account
func (db *Database) CreateAccount(account structs.Account) (id int32, err error) {
	err = db.Dbpool.QueryRow(context.Background(), "insert into Account (Name) values ($1) RETURNING ID", account.Name).Scan(&id)
	return id, err
}

// Create multiple payments
func (db *Database) CreatePayments(payments []structs.Payment) (err error) {

	log.Info("CreatePayments called")

	// Create a transaction to insert all payments at once
	tx, err := db.Dbpool.Begin(context.Background())
	if err != nil {
		return err
	}

	// Handle transaction after function returns
	defer func() {
		if p := recover(); p != nil {
			// Rollback the transaction if a panic occurred
			_ = tx.Rollback(context.Background())
			// Re-throw the panic
			panic(p)
		} else if err != nil {
			// Rollback the transaction if an error occurred
			_ = tx.Rollback(context.Background())
		} else {
			// Commit the transaction if everything is successful
			err = tx.Commit(context.Background())
		}
	}()

	// Insert payments within the transaction
	for _, payment := range payments {
		_, err := tx.Exec(context.Background(), "INSERT INTO Payment ( Sender, Receiver, Type, Amount) values ($1, $2, $3, $4)", payment.Sender, payment.Receiver, payment.Type, payment.Amount)
		if err != nil {
			return err
		}
	}
	return nil
}



func (db *Database) GetItems() ([]structs.Item, error) {
	var items []structs.Item
	rows, err := db.Dbpool.Query(context.Background(), "select * from items")
	if err != nil {
		return items, err
	}
	for rows.Next() {
		var item structs.Item
		err = rows.Scan(&item.ID, &item.Name, &item.Price)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (db *Database) UpdateItem(item structs.Item) (err error) {

	// TODO: Throw error if is_editable = False
	_, err = db.Dbpool.Exec(context.Background(), `
	UPDATE Item
	SET Name = $2, Price = $3, Image = $4
	WHERE ID = $1
	`, item.ID, item.Name, item.Price, item.Image)

	if err != nil {
		panic(err)
	}

	// Set items
	// for _, item := range settings.Items {
	// 	_, err = db.Dbpool.Query(context.Background(), `
	// 	INSERT INTO Items (Name, Price) VALUES ($1, $2)
	// 	ON CONFLICT (Name)
	// 	DO UPDATE SET Name = $1, Price = $2
	// 	`, item.Name, item.Price)
	// 	if err != nil {
	// 		log.Errorf("SetSettings failed: %v\n", err)
	// 	}
	// }
	return err
}

func (db *Database) GetSettings() (structs.Settings, error) {
	var settings structs.Settings
	err := db.Dbpool.QueryRow(context.Background(), `select * from Settings LIMIT 1`).Scan(&settings.ID, &settings.Color, &settings.Logo)
	if err != nil {
		log.Error("GetSettings failed", zap.Error(err))
	}
	return settings, err
}

func (db *Database) UpdateSettings(settings structs.Settings) (err error) {

	_, err = db.Dbpool.Query(context.Background(), `
	INSERT INTO Settings (Color, Logo) VALUES ($1, $2)
	ON CONFLICT (ID)
	DO UPDATE SET Color = $1, Logo = $2
	`, settings.Color, settings.Logo)

	if err != nil {
		log.Error("SetSettings failed:", zap.Error(err))
	}

	// Set items
	// for _, item := range settings.Items {
	// 	_, err = db.Dbpool.Query(context.Background(), `
	// 	INSERT INTO Items (Name, Price) VALUES ($1, $2)
	// 	ON CONFLICT (Name)
	// 	DO UPDATE SET Name = $1, Price = $2
	// 	`, item.Name, item.Price)
	// 	if err != nil {
	// 		log.Errorf("SetSettings failed: %v\n", err)
	// 	}
	// }
	return err
}

func (db *Database) GetVendorSettings() (string, error) {
	var settings string
	err := db.Dbpool.QueryRow(context.Background(), `select '{"credit":1.61,"qrcode":"/img/Augustin-QR-Code.png","idnumber":"123456789"}'`).Scan(&settings)
	if err != nil {
		log.Error("QueryRow failed:", zap.Error(err))
		return "", err
	}
	return settings, nil
}
