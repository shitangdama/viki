// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-02 05:20:06.4386613 +0800 CST m=+0.050000301

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/aticles": {
            "post": {
                "description": "新增文章",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "新增文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/aticles/": {
            "put": {
                "description": "修改文章",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "修改文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Tags",
                        "name": "tags",
                        "in": "body",
                        "schema": {
                            "type": "string",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    {
                        "description": "Title",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Source",
                        "name": "source",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Info",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/map%5Bstring%5Dinterface%7B%7D"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
