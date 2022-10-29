package View

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC 2565 ข้อที่ 2 วันเสาร์ที่ 29 ตุลาคม 2565",
    "version": "1.0.11"
  },
  "tags": [
    {
      "name": "feedback",
      "description": "MVC VIEW FOR FEEDBACK 63050096_2565_1"
    }
  ],
  "paths": {
    "/createFeedback": {
      "post": {
        "tags": [
          "feedback"
        ],
        "summary": "create feedback",
        "requestBody": {
          "description": "create feedback",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/feedbackReq"
              }
            },
            "application/xml": {
              "schema": {
                "$ref": "#/components/schemas/feedbackReq"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "create Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/feedbackRes"
                }
              }
            }
          }
        }
      }
    },
    "/getFeedback": {
      "get": {
        "tags": [
          "feedback"
        ],
        "summary": "get feedback",
        "responses": {
          "201": {
            "description": "create Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/feedbackRes"
                }
              }
            }
          }
        }
      }
    },
    "/updateFeedback": {
      "post": {
        "tags": [
          "feedback"
        ],
        "summary": "update feedback",
        "requestBody": {
          "description": "update feedback",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/UpdateReq"
              }
            },
            "application/xml": {
              "schema": {
                "$ref": "#/components/schemas/UpdateReq"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "update Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/UpdateRes"
                }
              }
            }
          }
        }
      }
    },
    "/adminUpdate": {
      "post": {
        "tags": [
          "feedback"
        ],
        "summary": "update feedback by admin",
        "requestBody": {
          "description": "update feedback by admin",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/UpdateReq"
              }
            },
            "application/xml": {
              "schema": {
                "$ref": "#/components/schemas/UpdateReq"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "update Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/UpdateRes"
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
      "feedbackReq": {
        "type": "object",
        "properties": {
          "first_name": {
            "type": "string"
          },
          "last_name": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "feedback": {
            "type": "string"
          }
        },
        "xml": {
          "name": "feedbackReq"
        }
      },
      "feedbackRes": {
        "type": "object",
        "properties": {
          "ref_no": {
            "type": "integer"
          },
          "first_name": {
            "type": "string"
          },
          "last_name": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "feedback": {
            "type": "string"
          },
          "time_stamp": {
            "type": "string"
          }
        },
        "xml": {
          "name": "feedbackRes"
        }
      },
      "UpdateReq": {
        "type": "object",
        "properties": {
          "ref_no": {
            "type": "integer"
          },
          "first_name": {
            "type": "string"
          },
          "last_name": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "feedback": {
            "type": "string"
          },
          "time_stamp": {
            "type": "string"
          }
        },
        "xml": {
          "name": "UpdateReq"
        }
      },
      "UpdateRes": {
        "type": "object",
        "properties": {
          "ref_no": {
            "type": "integer"
          },
          "first_name": {
            "type": "string"
          },
          "last_name": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "feedback": {
            "type": "string"
          },
          "time_stamp": {
            "type": "string"
          }
        },
        "xml": {
          "name": "UpdateRes"
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
	Description:      "63050096_2565_1",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
