package router

import (
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

	err = utils.AddDockerfiles(folderName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Zip the folder
	err = zipFolder(folderName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Upload the zip file to Amazon S3 bucket
	err = uploadZipToS3(filepath.Join("zippedfolders", folderName+".zip"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder zipped and uploaded to S3 successfully",
		"repoUrl": repoUrl,
	})
}

func zipFolder(folderName string) error {
	zipFileName := filepath.Join("zippedfolders", folderName+".zip")
	err := os.MkdirAll("zippedfolders", os.ModePerm)
	if err != nil {
		return err
	}

	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	baseFolder := filepath.Join("repos", folderName)

	err = filepath.Walk(baseFolder, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relPath, err := filepath.Rel(baseFolder, filePath)
			if err != nil {
				return err
			}

			zipFile, err := zipWriter.Create(relPath)
			if err != nil {
				return err
			}

			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(zipFile, file)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func uploadZipToS3(filePath string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("us-west-2"),
		},
	}))

	client := s3.New(sess)
	bucketName := "deployease"

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucketName),
		Key:                  aws.String(filepath.Base(filePath)),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}

// SetupRouter sets up the router
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/files", addFiles)

	return router
}
