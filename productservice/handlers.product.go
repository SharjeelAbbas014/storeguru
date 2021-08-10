package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addProductHandle(c *gin.Context) {
	// svc := engageDynamo()
	bs, _ := json.Marshal(c.Request.PostForm)
	fmt.Println(string(bs))
	c.IndentedJSON(http.StatusOK, "result: Success")
}

func getProductHandle(c *gin.Context) {
	svc := engageDynamo()
	prod := getProduct(svc, "3")
	fmt.Println(prod)
	c.IndentedJSON(http.StatusOK, prod)
}
