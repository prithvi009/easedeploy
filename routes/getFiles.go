package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
