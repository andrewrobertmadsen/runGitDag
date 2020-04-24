package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateJob(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Good Job! you created a Job (not really)"})
}

func GetJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Here is a list of Jobs (not really)"})
}

func GetJob(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Here is the Job you asked for (not really)"})
}

func DeleteJob(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted (not really)"})
}
