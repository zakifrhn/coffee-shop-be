package handlers

import (
	"fmt"
	"inter/internal/models"
	"inter/internal/repositories"
	"inter/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostData(ctx *gin.Context) {

	var ers error
	dataUser := models.User{
		Role: "user",
	}

	//	if dataUser.Role == "" {
	//dataUser.Role = "user"
	//return
	//}
	//else {
	// 	dataUser.Role = dataUser.Role
	// }

	fmt.Printf("ini adalah role nya: %s", dataUser.Role)

	if err := ctx.ShouldBind(&dataUser); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//TODO payload validation and Hash Password here

	dataUser.Pass_user, ers = pkg.HashPasword(dataUser.Pass_user)
	if ers != nil {
		ctx.AbortWithError(401, gin.Error{
			Err: ers,
		})
	}

	respone, err := h.CreateUser(&dataUser)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerUser) UpdateData(ctx *gin.Context) {

	var user models.User
	user.Id_user = ctx.Param("id_user")

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Print(user)

	respone, err := h.UpdateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
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
		ctx.AbortWithError(http.StatusBadRequest, err)
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

	ctx.JSON(200, respone)
}
