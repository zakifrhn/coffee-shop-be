package handlers

import (
	"fmt"
	"inter/internal/models"
	"inter/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	*repositories.RepoProduct
}

func NewProduct(r *repositories.RepoProduct) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) PostData(ctx *gin.Context) {

	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"message":     respone,
	})
}

func (h *HandlerProduct) UpdateData(ctx *gin.Context) {

	var product models.Product
	product.Id_product = ctx.Param("id")

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	respone, err := h.UpdateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerProduct) DeleteData(ctx *gin.Context) {

	var product models.Product
	product.Id_product = ctx.Param("id")

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	respone, err := h.DeleteProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerProduct) GetAllData(ctx *gin.Context) {

	var product models.Product

	page := ctx.DefaultQuery("page", "1")

	limit := ctx.DefaultQuery("limit", "5")

	pageInt, _ := strconv.Atoi(page)
	limInt, _ := strconv.Atoi(limit)

	if err := ctx.ShouldBindUri(&product); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	respone, err := h.GetAllProduct(&product, pageInt, limInt)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	products := map[string]interface{}{
		"data":     respone,
		"page":     pageInt,
		"perPages": len(respone),
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *HandlerProduct) GetProductCategory(ctx *gin.Context) {

	var product models.Product

	page := ctx.DefaultQuery("page", "1")

	limit := ctx.DefaultQuery("limit", "5")
	category := ctx.DefaultQuery("category", "")
	search := ctx.Query("search")
	fmt.Println(search)

	pageInt, _ := strconv.Atoi(page)
	limInt, _ := strconv.Atoi(limit)

	if err := ctx.ShouldBindQuery(&product); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	respone, err := h.GetCategory(&product, pageInt, limInt, category, search)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	products := map[string]interface{}{
		"data":     respone,
		"page":     pageInt,
		"perPages": len(respone),
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *HandlerProduct) GetProductName(ctx *gin.Context) {

	var product models.Product

	page := ctx.DefaultQuery("page", "1")

	limit := ctx.DefaultQuery("limit", "5")
	name := ctx.DefaultQuery("name", "")
	//nameProduct := ctx.DefaultQuery("name", "")
	// fmt.Println(name)

	pageInt, _ := strconv.Atoi(page)
	limInt, _ := strconv.Atoi(limit)

	if err := ctx.ShouldBindQuery(&product); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	respone, err := h.GetNameProduct(&product, pageInt, limInt, name)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	products := map[string]interface{}{
		"data":     respone,
		"page":     pageInt,
		"perPages": len(respone),
	}

	ctx.JSON(http.StatusOK, products)
}
