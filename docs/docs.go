// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/v1/books": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books_v1"
                ],
                "summary": "get all books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional header 1",
                        "name": "OptionalHeader1",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.book"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
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
                    "books_v1"
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
                    },
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional header 1",
                        "name": "OptionalHeader1",
                        "in": "header"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/books/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books_v1"
                ],
                "summary": "get a book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional header 1",
                        "name": "OptionalHeader1",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
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
                    "books_v1"
                ],
                "summary": "delete a book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional header 1",
                        "name": "OptionalHeader1",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
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
        },
        "/v1/queries": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queries"
                ],
                "summary": "queries API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional header 1",
                        "name": "OptionalHeader1",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Mandatory query 1",
                        "name": "MandateQuery1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional query 1",
                        "name": "OptionalQuery1",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.QueryParamsv1"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/books": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books_v2"
                ],
                "summary": "get all books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 2",
                        "name": "MandateHeader2",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 3",
                        "name": "MandateHeader3",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional header 1",
                        "name": "OptionalHeader1",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optinal header 2",
                        "name": "OptionalHeader2",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optinal header 3",
                        "name": "OptionalHeader3",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.book2"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "books_v2"
                ],
                "summary": "create a new book",
                "parameters": [
                    {
                        "description": "Book data",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.book2"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 2",
                        "name": "MandateHeader2",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 3",
                        "name": "MandateHeader3",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional header 1",
                        "name": "OptionalHeader1",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optinal header 2",
                        "name": "OptionalHeader2",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optinal header 3",
                        "name": "OptionalHeader3",
                        "in": "header"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.book2"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/books/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "books_v2"
                ],
                "summary": "get a book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 2",
                        "name": "MandateHeader2",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 3",
                        "name": "MandateHeader3",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional header 1",
                        "name": "OptionalHeader1",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optinal header 2",
                        "name": "OptionalHeader2",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optinal header 3",
                        "name": "OptionalHeader3",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.book2"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
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
                    "books_v2"
                ],
                "summary": "delete a book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 2",
                        "name": "MandateHeader2",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 3",
                        "name": "MandateHeader3",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional header 1",
                        "name": "OptionalHeader1",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optinal header 2",
                        "name": "OptionalHeader2",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Optinal header 3",
                        "name": "OptionalHeader3",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.book2"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
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
        },
        "/v2/queries": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queries"
                ],
                "summary": "queries API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "AuthToken",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 1",
                        "name": "MandateHeader1",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 2",
                        "name": "MandateHeader2",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory header 3",
                        "name": "MandateHeader3",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory query 1",
                        "name": "MandateQuery1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory query 2",
                        "name": "MandateQuery2",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Mandatory query 3",
                        "name": "MandateQuery3",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional query 1",
                        "name": "OptionalQuery1",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Optional query 2",
                        "name": "OptionalQuery2",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Optional query 3",
                        "name": "OptionalQuery3",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.QueryParamsv2"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.QueryParamsv1": {
            "type": "object",
            "required": [
                "mandateQuery1"
            ],
            "properties": {
                "mandateQuery1": {
                    "type": "string"
                },
                "optionalQuery1": {
                    "type": "string"
                }
            }
        },
        "main.QueryParamsv2": {
            "type": "object",
            "required": [
                "mandateQuery1",
                "mandateQuery2",
                "mandateQuery3"
            ],
            "properties": {
                "mandateQuery1": {
                    "type": "string"
                },
                "mandateQuery2": {
                    "type": "string"
                },
                "mandateQuery3": {
                    "type": "string"
                },
                "optionalQuery1": {
                    "type": "string"
                },
                "optionalQuery2": {
                    "type": "string"
                },
                "optionalQuery3": {
                    "type": "string"
                }
            }
        },
        "main.book": {
            "type": "object",
            "required": [
                "author",
                "id",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "main.book2": {
            "type": "object",
            "required": [
                "author",
                "id",
                "isbn",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isbn": {
                    "type": "string"
                },
                "title": {
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
	Title:            "Books API",
	Description:      "Books API documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
