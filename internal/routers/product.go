package routers

import (
	"inter/internal/handlers"
	"inter/internal/middleware"
	"inter/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.NewProduct(db)
	handler := handlers.NewProduct(repo)

	route.POST("/", middleware.AuthJwt("admin"), middleware.UploadFile, handler.PostData)
	route.PATCH("/:id", middleware.AuthJwt("admin"), middleware.UploadFile, handler.UpdateData)
	route.DELETE("/:id", middleware.AuthJwt("admin"), handler.DeleteData)
	route.GET("/", middleware.AuthJwt("admin", "user"), handler.GetAllData)
	route.GET("/product-category", handler.GetProductCategory)
	route.GET("/product-name", handler.GetProductName)
}
