// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/create_user": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "bad_request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not_found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/decrypt_file": {
            "post": {
                "description": "Provide the filename to decrypt",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server operations"
                ],
                "summary": "Decrypt a file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filename to decrypt",
                        "name": "filename",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "file_decrypted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad_request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not_found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/saved_file": {
            "post": {
                "description": "upload a file to encrypt",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server operations"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "File",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "file_uploaded",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad_request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not_found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/upload_file": {
            "post": {
                "description": "upload a file to encrypt",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server operations"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "File",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "file_uploaded",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad_request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not_found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/upload_key": {
            "post": {
                "description": "uploads a public key",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server operations"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "File",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "key_uploaded",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad_request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not_found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateUserRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/schemas.Address"
                },
                "contact": {
                    "$ref": "#/definitions/schemas.Contact"
                },
                "cpf": {
                    "type": "string"
                },
                "dateOfBirth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handler.CreateUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.UserResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.Address": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "zipcode": {
                    "type": "string"
                }
            }
        },
        "schemas.Contact": {
            "type": "object",
            "properties": {
                "celphone": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "schemas.UserResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/schemas.Address"
                },
                "contact": {
                    "$ref": "#/definitions/schemas.Contact"
                },
                "cpf": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "dateOfBirth": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
