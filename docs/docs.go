// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://google.com",
        "contact": {
            "name": "API Support",
            "email": "admin@mail.me"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/login": {
            "post": {
                "description": "Login User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "Login User",
                        "name": "dto.LoginReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRes"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Register new User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "description": "Register User",
                        "name": "dto.RegisterReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        },
        "/api/v1/social-media": {
            "get": {
                "description": "Get All new SocialMedia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Get All SocialMedia",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pagination Get All Social Media",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Pagination Get All Social Media",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.SocialMediaListRes"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new SocialMedia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "CreateSocialMedia",
                "parameters": [
                    {
                        "description": "Create SocialMedia",
                        "name": "dto.SocialMediaCreateReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SocialMediaCreateReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.SocialMedia"
                        }
                    }
                }
            }
        },
        "/api/v1/social-media/{id}": {
            "put": {
                "description": "Delete SocialMedia by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "DeleteSocialMedia",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.ResponseDefault"
                        }
                    }
                }
            }
        },
        "/api/v1/social-media/{{id}}": {
            "get": {
                "description": "GetOne SocialMedia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "GetOne SocialMedia",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.SocialMedia"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.LoginReq": {
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
        "dto.LoginRes": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterReq": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
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
        "dto.SocialMediaCreateReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.SocialMediaListRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.SocialMedia"
                    }
                },
                "last_page": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total_data": {
                    "type": "integer"
                }
            }
        },
        "dto.SocialMediaUpdateReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                }
            }
        },
        "entity.SocialMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/entity.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "helpers.ResponseDefault": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
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
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "My Gram API",
	Description:      "Service to manage MyGram data",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
