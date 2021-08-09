package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addProductHandle(c *gin.Context) {
	svc := engageDynamo()
	addProduct(svc, "3", "4", "NEW PORD", 2.2, "shaapm")
	c.IndentedJSON(http.StatusOK, "result: Success")
}
