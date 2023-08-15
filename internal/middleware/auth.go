package middleware

import (
	"inter/config"
	"inter/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwt(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			pkg.NewRes(401, &config.Result{
				Message: "Please Login",
			}).Send(ctx)
			return
		}

		if !strings.Contains(header, "Bearer") {
			pkg.NewRes(401, &config.Result{
				Message: "Invalid Header Value",
			}).Send(ctx)
			return
		}

		tokens := strings.Replace(header, "Bearer ", "", -1)

		//? mengecek tokennya, apakah masih aman atau sudah expired
		check, err := pkg.VerifyToken(tokens)
		if err != nil {
			pkg.NewRes(401, &config.Result{
				Message: err.Error(),
			}).Send(ctx)
			return
		}

		for _, r := range role {
			if r == check.Role {
				valid = true
			}
		}

		if !valid {
			pkg.NewRes(401, &config.Result{
				Data: "you not have permission to access",
			}).Send(ctx)
			return
		}

		//! mengambil id user
		ctx.Set("userId", check.Id)
		ctx.Next()
	}

}
