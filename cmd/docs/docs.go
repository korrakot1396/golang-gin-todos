// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/tasks": {
            "get": {
                "summary": "Retrieve all tasks",
                "description": "Retrieves a list of all tasks in the system",
                "produces": ["application/json"],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Task"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "summary": "Create a new task",
                "description": "Create a new task in the system",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "parameters": [
                    {
                        "in": "body",
                        "name": "task",
                        "description": "Task object to be created",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Task"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Task created successfully",
                        "schema": {
                            "$ref": "#/definitions/Task"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
			"delete": {
				"summary": "Delete all tasks",
				"description": "Deletes all tasks in the system",
				"responses": {
					"200": {
						"description": "All tasks deleted successfully"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			}
        },
        "/tasks/{id}": {
            "get": {
                "summary": "Retrieve a task by ID",
                "description": "Retrieves a task by its unique ID",
                "produces": ["application/json"],
                "parameters": [
                    {
                        "in": "path",
                        "name": "id",
                        "description": "Task ID",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/Task"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "404": {
                        "description": "Task not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
			"delete": {
				"summary": "Delete a task by ID",
				"description": "Deletes a task by its unique ID",
				"parameters": [
					{
						"in": "path",
						"name": "id",
						"description": "Task ID",
						"required": true,
						"type": "string"
					}
				],
				"responses": {
					"204": {
						"description": "Task deleted successfully"
					},
					"400": {
						"description": "Bad request"
					},
					"404": {
						"description": "Task not found"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			},
			"put": {
				"summary": "Update the title of a task by ID",
				"description": "Updates the title of a task by its unique ID",
				"parameters": [
					{
						"in": "path",
						"name": "id",
						"description": "Task ID",
						"required": true,
						"type": "string"
					},
					{
						"in": "body",
						"name": "title",
						"description": "New title for the task",
						"required": true,
						"schema": {
							"type": "string"
						}
					}
				],
				"responses": {
					"200": {
						"description": "Task title updated successfully"
					},
					"400": {
						"description": "Bad request"
					},
					"404": {
						"description": "Task not found"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			}
        },
		"/tasks/done/{id}": {
			"put": {
				"summary": "Mark a task as done",
				"description": "Marks a task as done by its unique ID",
				"parameters": [
					{
						"in": "path",
						"name": "id",
						"description": "Task ID",
						"required": true,
						"type": "string"
					}
				],
				"responses": {
					"200": {
						"description": "Task marked as done successfully"
					},
					"400": {
						"description": "Bad request"
					},
					"404": {
						"description": "Task not found"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			}
		}		
    },
    "definitions": {
        "Task": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "description": "Unique identifier for the task"
                },
                "title": {
                    "type": "string",
                    "description": "Title of the task"
                },
                "date": {
                    "type": "string",
                    "description": "Date of the task"
                },
                "status": {
                    "type": "string",
                    "description": "Status of the task"
                },
                "description": {
                    "type": "string",
                    "description": "Description of the task"
                },
                "image": {
                    "type": "string",
                    "description": "Base64 encoded image of the task, if available"
                }
			}
		}
    }
}`


// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "TaskTodos Service API",
	Description:      "A Task service API in Go using Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
