package routers

import (
	"github.com/gin-gonic/gin"
	"runGitDag/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Connections
	r.POST("/connections", api.CreateConnection)
	r.GET("/connections", api.GetConnections)
	r.GET("/connections/:id", api.GetConnection)
	r.DELETE("/connections/:id", api.DeleteConnection)

	// Jobs
	r.POST("/jobs", api.CreateJob)
	r.GET("/jobs", api.GetJobs)
	r.GET("/jobs/:id", api.GetJob)
	r.DELETE("/jobs/:id", api.DeleteJob)

	// ServiceUsers
	r.POST("/serviceusers", api.CreateServiceUser)
	r.GET("/serviceusers", api.GetServiceUsers)
	r.GET("/serviceusers/:id", api.GetServiceUser)
	r.DELETE("/serviceusers/:id", api.DeleteServiceUser)

	// Users
	r.POST("/users", api.CreateUser)
	r.GET("/users", api.GetUsers)
	r.GET("/users/:id", api.GetUser)
	r.DELETE("/users/:id", api.DeleteUser)

	return r
}