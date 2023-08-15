package middleware

import (
	"inter/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("banner_product")
	if err != nil {
		//log.Println("error disini")
		if err.Error() == "http: no such file" {
			ctx.Set("image", "")
			ctx.Next()
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Missing File",
		})
		return
	}

	//Open the file
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to open the file",
		})
		return
	}
	defer src.Close()

	result, err := pkg.Cloudinary(src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to open the file",
		})
		return
	}

	ctx.Set("image", result)
	ctx.Next()
}
