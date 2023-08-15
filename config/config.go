package config

import "github.com/gin-contrib/cors"

type Result struct {
	Data    interface{}
	Meta    interface{}
	Message interface{}
}

type Metas struct {
	Next  interface{} `json:"next"`
	Prev  interface{} `json:"prev"`
	Total int         `json:"total"`
}

var CorsConfig = cors.Config{
	AllowOrigins:     []string{"https://foo.com", "*"},
	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
}
