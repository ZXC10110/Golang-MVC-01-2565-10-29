package View

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC Structure",
    "version": "1.0.11"
  },
  "tags": [
    {
      "name": "Plagiarism",
      "description": "Testing of MVC's structure"
    }
  ],
  "paths": {
    "/getPlagiarism": {
      "get": {
        "tags": [
          "Plagiarism"
        ],
        "summary": "get Plagiarism",
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Plagiarism"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "Plagiarism": {
        "type": "object",
        "properties": {
          "pic_no": {
            "type": "integer",
            "format": "int64"
          }
        },
        "xml": {
          "name": "Plagiarism"
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
	Title:            "Golang MVC",
	Description:      "This is a sample server todo server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
