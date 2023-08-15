package routers

import (
	"inter/internal/handlers"
	"inter/internal/middleware"
	"inter/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ! / user
func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	//dependncy injection
	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", handler.PostData)
	route.PATCH("/:id_user", middleware.AuthJwt("admin", "user"), handler.UpdateData)
	route.GET("/:id_user", handler.GetDataUser)
	route.GET("/", handler.GetAllData)
	route.DELETE("/:id_user", handler.DeleteData)
}
