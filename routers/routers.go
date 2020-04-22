package routers

import (
	"github.com/gin-gonic/gin"
	"runGitDag/models"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/connections", models.CreateConnection)
	r.GET("/connections", models.GetConnections)
	r.POST("/jobs", models.CreateJob)
	r.GET("/jobs", models.GetJobs)

	return r
}