package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Example() {
	router := gin.Default()

	router.GET("/user", examples)
	router.GET("/query", queryString)
	router.POST("/params/:email/:pass", paramString)
	router.POST("/login", reqBody)

	router.PATCH("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "coba registrasi",
		})
	})

	router.Run()
}

func examples(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  200,
		"message": "pongping",
	})
}

type qString struct {
	Limit string `form:"limit"`
	Page  string `form:"page"`
}

// ! localhost:8080/query?page=1&limit=10
func queryString(ctx *gin.Context) {
	// page := ctx.Query("page")
	// limit := ctx.Query("limit")

	var data qString
	if err := ctx.ShouldBind(&data); err != nil {
		log.Println(err)

	}

	ctx.JSON(200, gin.H{
		"status": 200,
		"page":   data.Page,
		"limit":  data.Limit,
	})
}

type pString struct {
	Email string `from:"email"`
	Pass  string `from:"pass"`
}

// ! localhost:8080/params
func paramString(ctx *gin.Context) {
	// email := ctx.Param("email")
	// pass := ctx.Param("pass")

	var data pString
	if err := ctx.ShouldBind(&data); err != nil {
		log.Println(err)

	}

	ctx.JSON(200, gin.H{
		"status": 200,
		"email":  data.Email,
		"pass":   data.Pass,
	})
}

type body struct {
	Name  string `form:"name"`
	Image string `form:"image"`
}

func reqBody(ctx *gin.Context) {

	file, err := ctx.FormFile("image")

	if err != nil {
		log.Println(err)
	}

	log.Println(file.Filename)

	var data body
	if err := ctx.ShouldBind(&data); err != nil {
		log.Println(err)
	}

	ctx.JSON(200, data)
}
