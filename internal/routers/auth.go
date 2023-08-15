package routers

import (
	"inter/internal/handlers"
	"inter/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/auth")

	repo := repositories.NewUser(db)
	handler := handlers.NewAuth(repo)

	route.POST("/login", handler.Login)
}
