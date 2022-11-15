package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stephennwachukwu/go-RestAPI/helpers"
	"github.com/stephennwachukwu/go-RestAPI/models"
)

type UserController struct {
	UserService helpers.UserService
}

func New(userservice helpers.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

// creating the user controller
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// GetUser controller (a single user)
func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error})
	}
	ctx.JSON(http.StatusOK, user)
}

// GetAllUsers controller
func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error})
	}
	ctx.JSON(http.StatusOK, users)
}

// updateUser controller
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}

// creating the routes for the user controller

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")

	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.GET("/getall", uc.GetAllUsers)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:name", uc.DeleteUser)
}
