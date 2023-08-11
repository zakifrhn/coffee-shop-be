package routers

import (
	"inter/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	//! cara manual utk cors
	//router.Use(middleware.CORSMiddleware)

	//? cors using package
	//? code dibawah ini pindahkan ke file config.go
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://foo.com", "*"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
	// 	AllowHeaders:     []string{"Origin", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// AllowOriginFunc: func(origin string) bool {
	// 	return origin == "https://github.com"
	// },
	//MaxAge: 12 * time.Hour,
	// }))
	// router.Run()

	//? jika sudah dipindahkan tinggal panggil
	router.Use(cors.New(config.CorsConfig))

	user(router, db)
	product(router, db)

	return router
}
