// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {
            "name": "LiuYong",
            "email": "ly_yong@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ok/abb": {
            "post": {
                "description": "服务器对支付进行验证",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "告知服务器通过支付宝支付完成",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "userID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "需要接受的信息",
                        "name": "wantPayR",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.payOKAbbRecv"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回成功",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONResult"
                        }
                    }
                }
            }
        },
        "/refund/abb": {
            "post": {
                "description": "给订单号然后进行退款操作",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "请求退款",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "userID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "订单信息",
                        "name": "refundR",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.refundAbbRecv"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功信息",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONResult"
                        }
                    }
                }
            }
        },
        "/want/abb": {
            "post": {
                "description": "获得支付需要的来自支付宝的OrderInfo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "告知服务器想要支付",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "userID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "需要接受的信息",
                        "name": "wantPayR",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.wantPayAbbRecv"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回订单号和OrderInfo",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/v1.wantPayAbbSend"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.JSONResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.JSONResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "v1.payOKAbbRecv": {
            "type": "object",
            "required": [
                "order_info",
                "order_outside_id"
            ],
            "properties": {
                "order_info": {
                    "type": "string"
                },
                "order_outside_id": {
                    "type": "string"
                }
            }
        },
        "v1.refundAbbRecv": {
            "type": "object",
            "required": [
                "order_outside_id"
            ],
            "properties": {
                "order_outside_id": {
                    "type": "string"
                }
            }
        },
        "v1.wantPayAbbRecv": {
            "type": "object",
            "required": [
                "order_outside_id"
            ],
            "properties": {
                "order_outside_id": {
                    "type": "string"
                }
            }
        },
        "v1.wantPayAbbSend": {
            "type": "object",
            "properties": {
                "order_info": {
                    "type": "string"
                },
                "order_outside_id": {
                    "type": "string"
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
	Version:     "1.0",
	Host:        "localhost:8082",
	BasePath:    "/pay/api/v1",
	Schemes:     []string{},
	Title:       "支付服务",
	Description: "负责处理与支付和退款相关的业务",
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
