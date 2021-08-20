package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addProductHandle(c *gin.Context) {
	svc := engageDynamo()
	var prod Product
	c.BindJSON(&prod)
	fmt.Println(prod)
	addProduct(svc, prod.ProductId, prod.StoreId, prod.Title, prod.Price, prod.Tag, prod.Quantity, prod.LocationN)
	c.IndentedJSON(http.StatusOK, "result: Success")
}

func getProductHandle(c *gin.Context) {
	svc := engageDynamo()
	prod := getProduct(svc, "3")
	fmt.Println(prod)
	c.IndentedJSON(http.StatusOK, prod)
}
