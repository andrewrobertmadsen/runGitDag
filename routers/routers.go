package routers

import (
	"github.com/gin-gonic/gin"
	"runGitDag/models"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/connections", models.Connection.CreateConnection())
	r.GET("/connections", models.Connection.GetConnections())
	r.POST("/jobs", models.Job.CreateJob())
	r.GET("/jobs", models.Job.GetJobs())

	return r
}