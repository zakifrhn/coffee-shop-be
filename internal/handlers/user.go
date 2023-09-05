package handlers

import (
	"inter/config"
	"inter/internal/models"
	"inter/internal/repositories"
	"inter/pkg"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	repositories.RepoUserIF
}

func NewUser(r repositories.RepoUserIF) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostData(ctx *gin.Context) {

	var ers error
	dataUser := models.User{
		Role: "user",
	}

	if err := ctx.ShouldBind(&dataUser); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//TODO payload validation and Hash Password here
	_, ers = govalidator.ValidateStruct(&dataUser)
	if ers != nil {
		pkg.NewRes(401, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	dataUser.Pass_user, ers = pkg.HashPasword(dataUser.Pass_user)
	if ers != nil {
		ctx.AbortWithError(401, gin.Error{
			Err: ers,
		})
		return
	}

	respone, err := h.CreateUser(&dataUser)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, respone).Send(ctx)
}

func (h *HandlerUser) UpdateData(ctx *gin.Context) {

	var user models.User
	var ers error
	user.Id_user = ctx.Param("id_user")

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.Pass_user, ers = pkg.HashPasword(user.Pass_user)
	if ers != nil {
		ctx.AbortWithError(401, gin.Error{
			Err: ers,
		})
		return
	}

	respone, err := h.UpdateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"message":     respone,
	})
}

func (h *HandlerUser) GetDataUser(ctx *gin.Context) {

	var user models.User
	user.Id_user = ctx.Param("id_user")

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.GetUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerUser) GetAllData(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindUri(&user); err != nil {
		pkg.NewRes(400, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	respone, err := h.GetAllUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)

}

func (h *HandlerUser) DeleteData(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.DeleteUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"message":     respone,
	})
}
