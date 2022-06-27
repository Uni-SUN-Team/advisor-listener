// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/advisor-listener/api/advisors": {
            "get": {
                "description": "Advisor Listener for the service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advisor"
                ],
                "summary": "Advisor Listener",
                "operationId": "AdvisorListenerHandler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseAdvisors"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseFail"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseFail"
                        }
                    }
                }
            }
        },
        "/advisor-listener/api/advisors/:id": {
            "get": {
                "description": "Advisor Listener for the service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advisor"
                ],
                "summary": "Advisor Listener",
                "operationId": "AdvisorListenerIdHandler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseAdvisor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseFail"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseFail"
                        }
                    }
                }
            }
        },
        "/advisor-listener/api/advisors/slug/:slug": {
            "get": {
                "description": "Advisor Listener for the service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advisor"
                ],
                "summary": "Advisor Listener",
                "operationId": "AdvisorListenerSlugHandler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseAdvisors"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseFail"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseFail"
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Health checking for the service",
                "produces": [
                    "text/plain"
                ],
                "summary": "Health Check",
                "operationId": "HealthCheckHandler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ResponseAdvisor": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.advisorData"
                },
                "meta": {
                    "$ref": "#/definitions/model.pagination"
                }
            }
        },
        "model.ResponseAdvisors": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.advisorData"
                    }
                },
                "meta": {
                    "$ref": "#/definitions/model.pagination"
                }
            }
        },
        "model.ResponseFail": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "$ref": "#/definitions/model.error"
                }
            }
        },
        "model.advisorData": {
            "type": "object",
            "properties": {
                "SEO": {
                    "$ref": "#/definitions/model.seo"
                },
                "active": {
                    "type": "boolean"
                },
                "attachment": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.thumbnailLarge"
                    }
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.categories"
                    }
                },
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "dob": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "job_title": {
                    "type": "string"
                },
                "locale": {
                    "type": "string"
                },
                "publishedAt": {
                    "type": "string"
                },
                "quip": {
                    "type": "string"
                },
                "short_content": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "telephone": {
                    "type": "string"
                },
                "thumbnail": {
                    "$ref": "#/definitions/model.thumbnailLarge"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.categories": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "locale": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "publishedAt": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.error": {
            "type": "object",
            "properties": {
                "details": {},
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.formatAttribute": {
            "type": "object",
            "properties": {
                "ext": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "mime": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "size": {
                    "type": "number"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "model.pagination": {
            "type": "object",
            "properties": {
                "pagination": {
                    "$ref": "#/definitions/model.paginationContent"
                }
            }
        },
        "model.paginationContent": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageCount": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.seo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "keywords": {
                    "type": "string"
                },
                "metaDescription": {
                    "type": "string"
                },
                "metaTitle": {
                    "type": "string"
                },
                "preventIndexing": {
                    "type": "boolean"
                },
                "shareImage": {
                    "$ref": "#/definitions/model.seoShareImage"
                }
            }
        },
        "model.seoShareImage": {
            "type": "object",
            "properties": {
                "alt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "media": {
                    "$ref": "#/definitions/model.thumbnailLarge"
                }
            }
        },
        "model.thumbnailLarge": {
            "type": "object",
            "properties": {
                "alternativeText": {
                    "type": "string"
                },
                "caption": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "ext": {
                    "type": "string"
                },
                "formats": {
                    "$ref": "#/definitions/model.thumbnailLargeFormat"
                },
                "hash": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "mime": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "previewUrl": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "provider_metadata": {
                    "type": "string"
                },
                "size": {
                    "type": "number"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "model.thumbnailLargeFormat": {
            "type": "object",
            "properties": {
                "large": {
                    "$ref": "#/definitions/model.formatAttribute"
                },
                "medium": {
                    "$ref": "#/definitions/model.formatAttribute"
                },
                "small": {
                    "$ref": "#/definitions/model.formatAttribute"
                },
                "thumbnail": {
                    "$ref": "#/definitions/model.formatAttribute"
                },
                "xsmall": {
                    "$ref": "#/definitions/model.formatAttribute"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
