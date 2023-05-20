package controllers

import (
	"net/http"
	"strconv"

	"example.com/raaho-api/models"
	"example.com/raaho-api/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{UserService: userservice}
}

func (uc *UserController) SaveNewProduct(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.SaveNewProduct(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func (uc *UserController) EditProductDetails(ctx *gin.Context) {
	var details models.User
	if err := ctx.ShouldBindJSON(&details); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.EditProductDetails(&details)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetProductDetailsByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	idres, err := uc.UserService.GetProductDetailsByID(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, idres)
}

func (uc *UserController) GetProductsByRating(ctx *gin.Context) {
	ratingStr := ctx.Param("rating")
	rating, err := strconv.Atoi(ratingStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Rating"})
		return
	}
	res, err := uc.UserService.GetProductsByRating(&rating)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, res)
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/products")
	userroute.PUT("/addNewProduct", uc.SaveNewProduct)
	userroute.POST("/editProductDetails", uc.EditProductDetails)
	userroute.GET("/getProductById/:id", uc.GetProductDetailsByID)
	userroute.GET("/getProductByRating/:rating", uc.GetProductsByRating)
}
