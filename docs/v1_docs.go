// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplatev1 = `{
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
        "/books": {
            "get": {
                "tags": [
                    "books"
                ],
                "summary": "get all books",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.book"
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
                    "books"
                ],
                "summary": "create a new book",
                "parameters": [
                    {
                        "description": "Book data",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.book"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.book"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "get a book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.book"
                        }
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
                    "books"
                ],
                "summary": "delete a book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.book"
                        }
                    },
                    "404": {
                        "description": "Book not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Unable to delete book",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfov1 holds exported Swagger Info so clients can modify it
var SwaggerInfov1 = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Books API",
	Description:      "Books API documentation",
	InfoInstanceName: "v1",
	SwaggerTemplate:  docTemplatev1,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfov1.InstanceName(), SwaggerInfov1)
}