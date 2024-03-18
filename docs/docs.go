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
        "/actor": {
            "get": {
                "description": "Get a list of all actors available in the film library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actor"
                ],
                "summary": "GetList",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.ActorWithMovie"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update details of an existing actor in the film library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actor"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Actor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an actor in the film library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actor"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateActor"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an actor from the film library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actor"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.DeleteActor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Log in with user credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CRUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            }
        },
        "/movie": {
            "get": {
                "description": "Delete a movie from the film library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Actor name",
                        "name": "actor",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.MovieWithoudID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update details of an existing movie in the film library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new movie in the film library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateMovie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a movie from the film library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            }
        },
        "/movie/all": {
            "get": {
                "description": "Delete a movie from the film library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movie"
                ],
                "summary": "GetList",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.GetOrderedMovie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CRUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/sender.JSONResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Actor": {
            "type": "object",
            "properties": {
                "actor_id": {
                    "type": "integer"
                },
                "actor_name": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                }
            }
        },
        "domain.ActorWithMovie": {
            "type": "object",
            "properties": {
                "actor_name": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "domain.CRUser": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.CreateActor": {
            "type": "object",
            "properties": {
                "actor_name": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                }
            }
        },
        "domain.CreateMovie": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "description": {
                    "type": "string"
                },
                "movie_title": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "release_date": {
                    "type": "string"
                }
            }
        },
        "domain.DeleteActor": {
            "type": "object",
            "properties": {
                "actor_id": {
                    "type": "integer"
                }
            }
        },
        "domain.GetOrderedMovie": {
            "type": "object",
            "properties": {
                "order": {
                    "type": "string"
                }
            }
        },
        "domain.Movie": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "description": {
                    "type": "string"
                },
                "movie_id": {
                    "type": "integer"
                },
                "movie_title": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "release_date": {
                    "type": "string"
                }
            }
        },
        "domain.MovieWithoudID": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "releaseDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "sender.JSONResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "FilmLibrary",
	Description:      "API server for FilmLibrary",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
