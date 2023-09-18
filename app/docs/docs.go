// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "GNU Affero General Public License",
            "url": "https://www.gnu.org/licenses/agpl-3.0.txt"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hello/": {
            "get": {
                "description": "Return HelloWorld as sample API call",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Return HelloWorld",
                "responses": {}
            }
        },
        "/items/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "List Items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.Item"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Create Item",
                "parameters": [
                    {
                        "description": "Item Representation",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/items/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Delete Item",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/items/{id}/": {
            "put": {
                "description": "Requires multipart form (for image)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Update Item",
                "parameters": [
                    {
                        "description": "Item Representation",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/orders/": {
            "post": {
                "description": "Submits payment order to provider \u0026 saves it to database. Entries need to have an item id and a quantity (for entries without a price like tips, the quantity is the amount of cents). If no user is given, the order is anonymous.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Create Payment Order",
                "parameters": [
                    {
                        "description": "Payment Order",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateOrderResponse"
                        }
                    }
                }
            }
        },
        "/orders/verify/": {
            "post": {
                "description": "Verifies order and creates payments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Verify Payment Order",
                "parameters": [
                    {
                        "type": "string",
                        "format": "3043685539722561",
                        "description": "Order Code",
                        "name": "s",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "882d641c-01cc-442f-b894-2b51250340b5",
                        "description": "Transaction ID",
                        "name": "t",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.VerifyOrderResponse"
                        }
                    }
                }
            }
        },
        "/payments": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Get list of all payments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.Payment"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Create a set of payments",
                "parameters": [
                    {
                        "description": " Create Payment",
                        "name": "amount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreatePaymentsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/settings/": {
            "get": {
                "description": "Return configuration data of the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Return settings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.Settings"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update configuration data of the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Update settings",
                "parameters": [
                    {
                        "description": "Settings Representation",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.Settings"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/vendors/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "List Vendors",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.Vendor"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "Create Vendor",
                "parameters": [
                    {
                        "description": "Vendor Representation",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.Vendor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/vendors/{id}/": {
            "put": {
                "description": "Warning: Unfilled fields will be set to default values",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "Update Vendor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Vendor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Vendor Representation",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.Vendor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "Delete Vendor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Vendor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/vivawallet/transaction_order/": {
            "post": {
                "description": "Post your amount like {\"Amount\":100}, which equals 100 cents",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Create a transaction order",
                "parameters": [
                    {
                        "description": "Amount in cents",
                        "name": "amount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.TransactionOrder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.TransactionOrderResponse"
                            }
                        }
                    }
                }
            }
        },
        "/vivawallet/transaction_verification/": {
            "post": {
                "description": "Accepts {\"OrderCode\":\"1234567890\"} and returns {\"Verification\":true}, if successful",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Verify a transaction",
                "parameters": [
                    {
                        "description": "Transaction ID",
                        "name": "OrderCode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.TransactionVerification"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.TransactionVerificationResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.Item": {
            "type": "object",
            "properties": {
                "archived": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "licenseItem": {
                    "description": "License has to be bought before item",
                    "allOf": [
                        {
                            "$ref": "#/definitions/null.Int"
                        }
                    ]
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "description": "Price in cents",
                    "type": "integer"
                }
            }
        },
        "database.OrderEntry": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "item": {
                    "type": "integer"
                },
                "price": {
                    "description": "Price at time of purchase in cents",
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "receiver": {
                    "type": "integer"
                },
                "sender": {
                    "type": "integer"
                }
            }
        },
        "database.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "authorizedBy": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "order": {
                    "type": "integer"
                },
                "orderEntry": {
                    "type": "integer"
                },
                "receiver": {
                    "type": "integer"
                },
                "sender": {
                    "type": "integer"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "database.Settings": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logo": {
                    "type": "string"
                },
                "mainItem": {
                    "type": "integer"
                },
                "refundFees": {
                    "type": "boolean"
                }
            }
        },
        "database.Vendor": {
            "type": "object",
            "properties": {
                "balance": {
                    "description": "This is joined in from the account",
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "keycloakID": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "lastPayout": {
                    "type": "string",
                    "format": "date-time"
                },
                "licenseID": {
                    "type": "string"
                },
                "urlID": {
                    "description": "This is used for the QR code",
                    "type": "string"
                }
            }
        },
        "handlers.CreateOrderRequest": {
            "type": "object",
            "properties": {
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.CreateOrderRequestEntry"
                    }
                },
                "user": {
                    "type": "string"
                },
                "vendor": {
                    "type": "integer"
                }
            }
        },
        "handlers.CreateOrderRequestEntry": {
            "type": "object",
            "properties": {
                "item": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "handlers.CreateOrderResponse": {
            "type": "object",
            "properties": {
                "smartCheckoutURL": {
                    "type": "string"
                }
            }
        },
        "handlers.CreatePaymentsRequest": {
            "type": "object",
            "properties": {
                "payments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.Payment"
                    }
                }
            }
        },
        "handlers.TransactionOrder": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                }
            }
        },
        "handlers.TransactionOrderResponse": {
            "type": "object",
            "properties": {
                "smartCheckoutURL": {
                    "type": "string"
                }
            }
        },
        "handlers.TransactionVerification": {
            "type": "object",
            "properties": {
                "orderCode": {
                    "type": "integer"
                }
            }
        },
        "handlers.TransactionVerificationResponse": {
            "type": "object",
            "properties": {
                "verification": {
                    "type": "boolean"
                }
            }
        },
        "handlers.VerifyOrderResponse": {
            "type": "object",
            "properties": {
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/database.OrderEntry"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "orderCode": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "transactionID": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                },
                "vendor": {
                    "type": "integer"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        },
        "null.Int": {
            "type": "object",
            "properties": {
                "int64": {
                    "type": "integer"
                },
                "valid": {
                    "description": "Valid is true if Int64 is not NULL",
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Augustin Swagger",
	Description:      "This swagger describes every endpoint of this project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
