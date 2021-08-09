package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addProductHandle(c *gin.Context) {
	svc := engageDynamo()
	addProduct(svc, "3", "4", "NEW PORD", 2.2, "shaapm")
	c.IndentedJSON(http.StatusOK, "result: Success")
}

func getProductHandle(c *gin.Context) {
	svc := engageDynamo()
	prod := getProduct(svc, "3")
	fmt.Println(prod)
	c.IndentedJSON(http.StatusOK, prod)
}
