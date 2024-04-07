package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", index)

	router.Run()
}

func index(c *gin.Context) {

	c.String(http.StatusOK, "index")
}
