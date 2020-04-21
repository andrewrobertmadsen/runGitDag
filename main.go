package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hey there, Casey! I think we should try gin for our v1 of runGitDag. What do you think?")
	})

	router.Run(":3333")
}
