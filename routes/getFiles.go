package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prithvi009/easedeploy/utils"
)

func addFiles(c *gin.Context) {

	var jsonData struct {
		RepoURL string `json:"repoUrl"`
	}

	// Bind the JSON payload to the struct
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract repoUrl from the bound struct
	repoUrl := jsonData.RepoURL

	repoName := strings.Split(repoUrl, "/")[4]
	folderName := utils.GenerateUniqueFileName(repoName)

	err := utils.CloneRepository(repoUrl, folderName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Files added successfully",
		"repoUrl": repoUrl,
	})

}

// SetupRouter sets up the router
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/files", addFiles)

	return router
}