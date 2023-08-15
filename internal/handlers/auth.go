package handlers

import (
	"fmt"
	"inter/config"
	"inter/internal/repositories"
	"inter/pkg"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `db:"email" json:"email" form:"email"`
	Password string `db:"pass" json:"password" form:"password"`
}

type HandlerAuth struct {
	*repositories.RepoUser
}

func NewAuth(r *repositories.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var dataUser User

	if ers := ctx.ShouldBind(&dataUser); ers != nil {
		pkg.NewRes(500, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	users, err := h.GetAuthData(dataUser.Email)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(users.Pass_user, dataUser.Password); err != nil {
		pkg.NewRes(401, &config.Result{
			Data: "Password Salah",
		}).Send(ctx)
		return
	}

	jwtt := pkg.NewToken(users.Id_user, users.Role)
	//fmt.Println(jwtt)
	tokens, err := jwtt.Generate()
	if err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	fmt.Println(tokens)
	pkg.NewRes(200, &config.Result{Data: tokens}).Send(ctx)
}
