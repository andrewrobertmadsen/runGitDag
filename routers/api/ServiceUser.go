package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateServiceUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Good Job! you created a ServiceUser (not really)"})
}

func GetServiceUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Here is a list of ServiceUsers (not really)"})
}

func GetServiceUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Here is the ServiceUser you asked for (not really)"})
}

func DeleteServiceUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ServiceUser deleted (not really)"})
}
