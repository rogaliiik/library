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
        "/v1/auth/sign-in": {
            "post": {
                "description": "user sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user-auth"
                ],
                "summary": "User SignIn",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.signInOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    }
                }
            }
        },
        "/v1/auth/sign-up": {
            "post": {
                "description": "user sign up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user-auth"
                ],
                "summary": "User SignUp",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.signUpOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    }
                }
            }
        },
        "/v1/book": {
            "get": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "get all books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookApi"
                ],
                "summary": "Book Get all",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Book"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    }
                }
            }
        },
        "/v1/book/create": {
            "post": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "book create",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookApi"
                ],
                "summary": "Book Create",
                "parameters": [
                    {
                        "description": "book info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.bookCreateOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    }
                }
            }
        },
        "/v1/book/{bookId}": {
            "get": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "book get by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookApi"
                ],
                "summary": "Book Get by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "book id",
                        "name": "bookId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "update book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookApi"
                ],
                "summary": "Book Update",
                "parameters": [
                    {
                        "description": "book update info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.BookUpdateInput"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "book id",
                        "name": "bookId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.statusOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "delete book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookApi"
                ],
                "summary": "Book Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "book id",
                        "name": "bookId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.statusOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Book": {
            "type": "object",
            "required": [
                "author"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 1
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "domain.BookUpdateInput": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@gmail.com"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "v1.bookCreateOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "v1.errorMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.signInOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "v1.signUpOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "v1.statusOutput": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "UserAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "LibraryAPI",
	Description:      "API for Library app.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}