package controller

import (
	"example/Go-Api-Tutorial/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func addProduct(c *gin.Context) {
	var newProduct model.Product
	err := c.ShouldBindJSON(&newProduct)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
    return
	}

}

func editProduct(c *gin.Context) {

}

func getAllProduct(c *gin.Context){

}
func getProductbyId(c *gin.Context){

}
func deleteProduct(c *gin.Context){

}