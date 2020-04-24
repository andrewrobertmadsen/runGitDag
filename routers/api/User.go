package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runGitDag/common"
	"runGitDag/models"
)

func CreateUser(c *gin.Context) {
	db := common.InitDb()
	defer db.Close()

	var user models.User
	c.Bind(&user)
	db.Create(&user)
	c.JSON(200, gin.H{"success": user})
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Here is a list of Users (not really)"})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Here is the User you asked for (not really)"})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User deleted (not really)"})
}
