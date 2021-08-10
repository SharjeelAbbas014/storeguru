package main

func initializeRoutes() {
	router.POST("/", addProductHandle)
	router.GET("/search", getProductHandle)
}
