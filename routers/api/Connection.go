package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateConnection(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Good Job! you created a Connection (not really)"})
}

func GetConnections(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Here is a list of Connections (not really)"})
}

func GetConnection(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Here is the Connection you asked for (not really)"})
}

func DeleteConnection(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Connection deleted (not really)"})
}
