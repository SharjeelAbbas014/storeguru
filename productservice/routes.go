package main

func initializeRoutes() {
	router.GET("/", addProductHandle)
	router.GET("/search", getProductHandle)
}
