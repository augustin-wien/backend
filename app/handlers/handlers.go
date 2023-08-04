//	@title			Augustin Swagger
//	@version		0.0.1
//	@description	This swagger describes every endpoint of this project.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	GNU Affero General Public License
//	@license.url	https://www.gnu.org/licenses/agpl-3.0.txt

//	@host		localhost:3000
//	@BasePath	/api/

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

package handlers

import (
	"augustin/utils"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mitchellh/mapstructure"
	_ "github.com/swaggo/files"        // swagger embed files
	_ "github.com/swaggo/http-swagger" // http-swagger middleware

	_ "github.com/swaggo/files" // swagger embed files

	"augustin/database"
)

var log = utils.GetLogger()

func respond(w http.ResponseWriter, err error, payload interface{}) {
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, payload)
}


// ReturnHelloWorld godoc
//
//	 	@Summary 		Return HelloWorld
//		@Description	Return HelloWorld as sample API call
//		@Tags			core
//		@Accept			json
//		@Produce		json
//		@Router			/hello/ [get]
//
// HelloWorld API Handler fetching data from database
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	greeting, err := database.Db.GetHelloWorld()
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)

		return
	}
	utils.WriteJSON(w, http.StatusOK, greeting)
}



// Users ----------------------------------------------------------------------

// ListVendors godoc
//
//	 	@Summary 		List Vendors
//		@Tags			vendors
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	database.Vendor
//		@Router			/api/vendors/ [get]
//
func ListVendors(w http.ResponseWriter, r *http.Request) {
	users, err := database.Db.ListVendors()
	respond(w, err, users)
}

// CreateVendor godoc
//
//	 	@Summary 		Create Vendor
//		@Tags			vendors
//		@Accept			json
//		@Produce		json
//		@Success		200
//	    @Param		    data body database.Vendor true "Vendor Representation"
//		@Router			/api/vendors/ [post]
//
func CreateVendor(w http.ResponseWriter, r *http.Request) {
	var vendor database.Vendor
	err := utils.ReadJSON(w, r, &vendor)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	log.Info("id ", vendor.ID, "fn ", vendor.FirstName, "bl ", vendor.Balance)

	id, err := database.Db.CreateVendor(vendor)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, id)
}

// UpdateVendor godoc
//
//	 	@Summary 		Update Vendor
//		@Description	Warning: Unfilled fields will be set to default values
//		@Tags			vendors
//		@Accept			json
//		@Produce		json
//		@Success		200
//      @Param          id   path int  true  "Vendor ID"
//	    @Param		    data body database.Vendor true "Vendor Representation"
//		@Router			/api/vendors/{id} [put]
//
func UpdateVendor(w http.ResponseWriter, r *http.Request) {
	vendorID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var vendor database.Vendor
	err = utils.ReadJSON(w, r, &vendor)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	err = database.Db.UpdateVendor(vendorID, vendor)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, vendor)
}

// DeleteVendor godoc
//
//	 	@Summary 		Delete Vendor
//		@Tags			vendors
//		@Accept			json
//		@Produce		json
//		@Success		200
//      @Param          id   path int  true  "Vendor ID"
//		@Router			/api/vendors/{id} [delete]
//
func DeleteVendor(w http.ResponseWriter, r *http.Request) {
	vendorID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = database.Db.DeleteVendor(vendorID)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


// Items (that can be sold) ---------------------------------------------------


// ListItems godoc
//
//	 	@Summary 		List Items
//		@Tags			Items
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	database.Item
//		@Router			/api/items/ [get]
//
func ListItems(w http.ResponseWriter, r *http.Request) {
	items, err := database.Db.ListItems()
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, items)
}

// CreateItem godoc
//
//	 	@Summary 		Create Item
//		@Tags			Items
//		@Accept			json
//		@Produce		json
//	    @Param		    data body database.Item true "Item Representation"
//		@Success		200	 {int}	id
//		@Router			/api/items/ [post]
//
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item database.Item
	err := utils.ReadJSON(w, r, &item)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	id, err := database.Db.CreateItem(item)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, id)
}


// UpdateItem godoc
//
//	 	@Summary 		Update Item
//		@Description	Requires multipart form (for image)
//		@Tags			Items
//		@Accept			json
//		@Produce		json
//	    @Param		    data body database.Item true "Item Representation"
//		@Success		200
//		@Router			/api/items/{id}/ [put]
//
// UpdateItem requires a multipart form
// https://www.sobyte.net/post/2022-03/go-multipart-form-data/
func UpdateItem(w http.ResponseWriter, r *http.Request) {

	// Read multipart form
	r.ParseMultipartForm(32 << 20)
	mForm := r.MultipartForm

	// Handle normal fields
	var item database.Item
	fields := mForm.Value
	err := mapstructure.Decode(fields, &item)
    if err != nil {
        log.Error(err)
    }

    // Get file from image field
    file, header, err := r.FormFile("Image")
    if err != nil {
        log.Error(err)
    }
    defer file.Close()

	// Debugging
    name := strings.Split(header.Filename, ".")
	log.Infof("Uploading %s\n", name[0])

	// Save file
	path := "/img/"+header.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Error(err)
	}
	io.Copy(f, file)
	item.Image = path

	// Save item to database
	err = database.Db.UpdateItem(item)
	if err != nil {
		log.Error(err)
	}

}

// DeleteItem godoc
//
//	 	@Summary 		Delete Item
//		@Tags			Items
//		@Accept			json
//		@Produce		json
//		@Success		200
//      @Param          id   path int  true  "Item ID"
//		@Router			/api/Items/{id} [delete]
//
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	ItemID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = database.Db.DeleteItem(ItemID)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


// Payments (from one account to another account) -----------------------------

// ListPayments godoc
//
//	 	@Summary 		Get list of all payments
//		@Tags			core
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	database.Payment
//		@Router			/payments [get]
func ListPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := database.Db.ListPayments()
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, payments)
}


// CreatePayments godoc
//
//	 	@Summary 		Create a set of payments
//		@Tags			core
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	database.PaymentType
//		@Router			/payments [post]
func CreatePayments(w http.ResponseWriter, r *http.Request) {
	var paymentBatch database.PaymentBatch
	err := utils.ReadJSON(w, r, &paymentBatch)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = database.Db.CreatePayments(paymentBatch.Payments)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}

// Settings -------------------------------------------------------------------

// getSettings godoc
//
//	 	@Summary 		Return settings
//		@Description	Return configuration data of the system
//		@Tags			core
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	database.Settings
//		@Router			/settings/ [get]
//
func getSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := database.Db.GetSettings()
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, http.StatusOK, settings)
}
