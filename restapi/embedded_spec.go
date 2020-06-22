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
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "For the Kabanero REST service",
    "title": "Kabanero REST API",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "message"
        ],
        "responses": {
          "200": {
            "description": "standard message response",
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        }
      }
    },
    "/test": {
      "get": {
        "tags": [
          "message"
        ],
        "responses": {
          "200": {
            "description": "standard message response",
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        }
      }
    },
    "/v1/describe/stacks/{stackName}/versions/{version}": {
      "get": {
        "operationId": "describe",
        "parameters": [
          {
            "type": "string",
            "name": "stackName",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "version",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "describe stack",
            "schema": {
              "$ref": "#/definitions/DescribeStack"
            }
          },
          "500": {
            "description": "describe stack error",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/stacks": {
      "get": {
        "operationId": "list",
        "responses": {
          "200": {
            "description": "list successful",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/KabaneroStack"
              }
            }
          },
          "500": {
            "description": "list stack error",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "message"
        ],
        "responses": {
          "200": {
            "description": "standard message response",
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "DescribeStack": {
      "type": "object",
      "properties": {
        "apps": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "instance": {
                "type": "string"
              },
              "managedby": {
                "type": "string"
              },
              "name": {
                "type": "string"
              },
              "partof": {
                "type": "string"
              },
              "version": {
                "type": "string"
              }
            }
          }
        },
        "digest check": {
          "type": "string"
        },
        "image": {
          "type": "string"
        },
        "image digest": {
          "type": "string"
        },
        "image name": {
          "type": "string"
        },
        "kabanero digest": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "project": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "KabaneroStack": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "status": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "Image name": {
                "type": "string"
              },
              "digest check": {
                "type": "string"
              },
              "image digest": {
                "type": "string"
              },
              "kabanero digest": {
                "type": "string"
              },
              "status": {
                "type": "string"
              },
              "version": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "message": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "For the Kabanero REST service",
    "title": "Kabanero REST API",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "message"
        ],
        "responses": {
          "200": {
            "description": "standard message response",
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        }
      }
    },
    "/test": {
      "get": {
        "tags": [
          "message"
        ],
        "responses": {
          "200": {
            "description": "standard message response",
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        }
      }
    },
    "/v1/describe/stacks/{stackName}/versions/{version}": {
      "get": {
        "operationId": "describe",
        "parameters": [
          {
            "type": "string",
            "name": "stackName",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "version",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "describe stack",
            "schema": {
              "$ref": "#/definitions/DescribeStack"
            }
          },
          "500": {
            "description": "describe stack error",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/stacks": {
      "get": {
        "operationId": "list",
        "responses": {
          "200": {
            "description": "list successful",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/KabaneroStack"
              }
            }
          },
          "500": {
            "description": "list stack error",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "message"
        ],
        "responses": {
          "200": {
            "description": "standard message response",
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "DescribeStack": {
      "type": "object",
      "properties": {
        "apps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/DescribeStackAppsItems0"
          }
        },
        "digest check": {
          "type": "string"
        },
        "image": {
          "type": "string"
        },
        "image digest": {
          "type": "string"
        },
        "image name": {
          "type": "string"
        },
        "kabanero digest": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "project": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "DescribeStackAppsItems0": {
      "type": "object",
      "properties": {
        "instance": {
          "type": "string"
        },
        "managedby": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "partof": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "KabaneroStack": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "status": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/KabaneroStackStatusItems0"
          }
        }
      }
    },
    "KabaneroStackStatusItems0": {
      "type": "object",
      "properties": {
        "Image name": {
          "type": "string"
        },
        "digest check": {
          "type": "string"
        },
        "image digest": {
          "type": "string"
        },
        "kabanero digest": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "message": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
}
