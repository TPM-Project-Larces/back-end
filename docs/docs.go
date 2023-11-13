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
        "/auth/login": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "bad_request",
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
        "/encryption/decrypt_file": {
            "post": {
                "description": "Provide the filename to decrypt",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Encryption"
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
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/encryption/saved_file": {
            "post": {
                "description": "Save a file to encrypt",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Encryption"
                ],
                "summary": "Save file",
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
                        "description": "file_saved",
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
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/encryption/upload_file": {
            "post": {
                "description": "Upload a file to encrypt",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Encryption"
                ],
                "summary": "Upload file",
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
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/encryption/upload_key": {
            "post": {
                "description": "Uploads a public key",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Encryption"
                ],
                "summary": "Upload key",
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
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get a list of all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ListUsersResponse"
                        }
                    },
                    "500": {
                        "description": "internal_server_error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "User data to Update",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "bad_request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "user_not_found",
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
            },
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "bad_request",
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
            },
            "delete": {
                "description": "Deletes a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.DeleteUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.DeleteUserResponse"
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
        "/user/username": {
            "get": {
                "description": "Provide the user data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Find user by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User` + "`" + `s username to find",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ShowUserResponse"
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
        "model.Address": {
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
        "model.Contact": {
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
        "schemas.AuthResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.CreateUserRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/model.Address"
                },
                "contact": {
                    "$ref": "#/definitions/model.Contact"
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
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.CreateUserResponse": {
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
        "schemas.DeleteUserRequest": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.DeleteUserResponse": {
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
        "schemas.ListUsersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.UserResponse"
                    }
                }
            }
        },
        "schemas.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "schemas.ShowUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.UserResponse"
                }
            }
        },
        "schemas.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/model.Address"
                },
                "contact": {
                    "$ref": "#/definitions/model.Contact"
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
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.UpdateUserResponse": {
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
        "schemas.UserResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/model.Address"
                },
                "contact": {
                    "$ref": "#/definitions/model.Contact"
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
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Server API",
	Description:      "Server Operations",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
