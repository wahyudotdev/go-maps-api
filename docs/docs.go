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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/maps/route": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Maps"
                ],
                "summary": "Get routes between 2 points",
                "parameters": [
                    {
                        "type": "number",
                        "name": "end_lat",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "name": "end_lng",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "name": "start_lat",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "name": "start_lng",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Maps API",
	Description:      "Public API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
