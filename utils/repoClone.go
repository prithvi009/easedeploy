package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func CloneRepository(repoURL, folderName string) error {

	// Clone the repository
	destinationDir := filepath.Join("./repos", folderName)
	_, err := git.PlainClone(destinationDir, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "prithvi009",
			Password: "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", // Replace with your GitHub Personal Access Token
		},
	})
	if err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	fmt.Println("Successfully cloned repository to")

	return nil
}
