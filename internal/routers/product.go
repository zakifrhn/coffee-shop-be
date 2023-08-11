package routers

import (
	"inter/internal/handlers"
	"inter/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.NewProduct(db)
	handler := handlers.NewProduct(repo)

	route.POST("/", handler.PostData)
	route.PATCH("/:id", handler.UpdateData)
	route.DELETE("/:id", handler.DeleteData)
	route.GET("/", handler.GetAllData)
	route.GET("/product-category", handler.GetProductCategory)
	route.GET("/product-name", handler.GetProductName)
}
