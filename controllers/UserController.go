package controllers

import (
	"fmt"
	"net/http"
	"noteapp/database"
	"noteapp/models"
	structure "noteapp/struct"
	"noteapp/utils"

	"github.com/gin-gonic/gin"
)

func RegitserUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body = structure.RegisterInputStruct{}
		if err := ctx.Bind(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errror": err.Error(),
			})
			return
		}
		if body.Password != body.Confirm_Password {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Password and Confirm Password do not match.",
			})
			return
		}
		hash, err := utils.HashPassword(body.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		user := models.User{Email: body.Email, Password: string(hash), Full_name: body.Full_Name, Phone: body.Phone}
		result := database.DB.Create(&user)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": result.Error,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   structure.RegisterOutputStruct{Email: user.Email, Full_Name: user.Full_name, Phone: user.Phone, Created_at: user.CreatedAt},
		})
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body = structure.LoginInputStruct{}
		if err := ctx.Bind(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var user models.User
		if result := database.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
			ctx.JSON(http.StatusNoContent, gin.H{
				"error": "user not found.",
			})
			return
		}
		if err := utils.VerifyPassword(user.Password, body.Password); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "wrong credentials.",
			})
			fmt.Println("wrong password bro wrong")
			return
		}
		fmt.Println("yeppp correct password")

	}
}
