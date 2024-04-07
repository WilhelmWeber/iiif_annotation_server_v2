package controllers

import (
	"fmt"
	"net/http"

	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/libs"
	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/model"
	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/repository"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *repository.UserService
}

func NewUserController(s *repository.UserService) *UserController {
	return &UserController{s}
}

/*POST: /user */
func (u *UserController) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}

	if isExist := u.Service.GetByName(user.Username); isExist.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Request username has already existed",
		})
		return
	}

	encryptPw, err := libs.Encrypt(user.Password)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}

	obj := &model.User{Username: user.Username, Password: encryptPw}
	if err := u.Service.Create(obj); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created with success",
		"user": gin.H{
			"name": obj.Username,
		},
	})
}

/*POST: /auth */
func (u *UserController) Login(c *gin.Context) {
	var user *model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error. Please Inform it of Administrator",
		})
		return
	}

	found_user := u.Service.GetByName(user.Username)
	if found_user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to Login",
		})
		return
	}
	if err := libs.Compare(user.Password, found_user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to Login",
		})
		return
	}

	token, err := libs.GenerateToken(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot Create Token",
		})
		return
	}
	//TODO: Domain名に変える
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message":  "login with success",
		"username": user.Username,
	})
}
