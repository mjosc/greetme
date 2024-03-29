// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Mock Server",
    "version": "0.1.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/users/{name}": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "userByName",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "The name of the user to be returned",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A success response describing the user if it exists and an error if not",
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string",
                  "minLength": 1,
                  "x-nullable": false
                },
                "id": {
                  "description": "The user ID",
                  "type": "integer",
                  "minimum": 1,
                  "x-nullable": false
                },
                "message": {
                  "type": "string",
                  "minLength": 1,
                  "x-nullable": false
                },
                "name": {
                  "description": "The name of the user",
                  "type": "string",
                  "minLength": 1,
                  "x-nullable": false
                },
                "valid": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Mock Server",
    "version": "0.1.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/users/{name}": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "userByName",
        "parameters": [
          {
            "minLength": 1,
            "type": "string",
            "description": "The name of the user to be returned",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A success response describing the user if it exists and an error if not",
            "schema": {
              "type": "object",
              "properties": {
                "error": {
                  "type": "string",
                  "minLength": 1,
                  "x-nullable": false
                },
                "id": {
                  "description": "The user ID",
                  "type": "integer",
                  "minimum": 1,
                  "x-nullable": false
                },
                "message": {
                  "type": "string",
                  "minLength": 1,
                  "x-nullable": false
                },
                "name": {
                  "description": "The name of the user",
                  "type": "string",
                  "minLength": 1,
                  "x-nullable": false
                },
                "valid": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
}
