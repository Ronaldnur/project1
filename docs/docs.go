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
        "/todos": {
            "get": {
                "description": "Get Todos Data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "operationId": "get-Todo-data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ReadDataTodoResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Todo item.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Create a new Todo",
                "parameters": [
                    {
                        "description": "Todo request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewTodosRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewTodosResponse"
                        }
                    }
                }
            }
        },
        "/todos/{todoId}": {
            "get": {
                "description": "Get a Todo by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "operationId": "get-Todo-by-ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "todoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ReadDataByIdTodoResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a Todo by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "operationId": "update-Todo-by-ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "todoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Todo request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewTodosRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewTodosResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Todo by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "operationId": "delete-Todo-by-ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "todoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.NewTodosRequest": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.NewTodosResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "dto.ReadDataByIdTodoResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dto.TodoWithId"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "dto.ReadDataTodoResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.TodoWithId"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "dto.TodoWithId": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
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
