package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func cloneRepository(repoURL, folderName string) error {

	password := os.Getenv("GITHUB_ACCESS_KEY")

	github_access_key := password
	// Clone the repository
	repo, err := git.PlainClone(folderName, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "prithvi009",
			Password: github_access_key, // Replace with your GitHub Personal Access Token
		},
	})
	if err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	// List files in the repository
	files, err := worktree.Filesystem.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	// Create the sourcefiles directory if it doesn't exist
	sourcefilesDir := filepath.Join(folderName, "sourcefiles")
	err = os.MkdirAll(sourcefilesDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create sourcefiles directory: %w", err)
	}

	// Copy files to the sourcefiles directory
	for _, file := range files {
		fileName := generateUniqueFileName(filepath.Base(file.Name()))
		filePath := filepath.Join(sourcefilesDir, fileName)

		// Copy file to the sourcefiles directory
		err := os.Rename(file.Name(), filePath)
		if err != nil {
			return fmt.Errorf("failed to copy file: %w", err)
		}
	}

	return nil
}

func generateUniqueFileName(fileName string) string {
	// Generate a unique file name by appending a timestamp
	extension := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, extension)
	uniqueName := fmt.Sprintf("%s_%d%s", baseName, time.Now().UnixNano(), extension)
	return uniqueName
}

func main() {
	repoURL := "https://github.com/prithvi009/easedeploy"
	folderName := "repo_folder"

	err := cloneRepository(repoURL, folderName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Repository cloned successfully.")
}
