package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Connection struct {
	Host string
	Port int
	Type ConnectionType
}

type ConnectionType string

const (
	Hive     ConnectionType = "Hive"
	Impala                  = "Impala"
	Presto                  = "Presto"
	SparkSql                = "SparkSql"
)

func CreateConnection (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Good Job! you created a Connection (not really)"})
}
func GetConnections (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Here is a list of Connections (not really)"})
}
