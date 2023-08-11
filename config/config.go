package config

import "github.com/gin-contrib/cors"

var CorsConfig = cors.Config{
	AllowOrigins:     []string{"https://foo.com", "*"},
	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
	AllowHeaders:     []string{"Origin", "Authorization"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
}
