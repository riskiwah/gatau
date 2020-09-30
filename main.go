package main

import (
	"gatau/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.ShowToIndex)
	// r.GET("/", handlers.ShowHostname)
	r.Run()
}
