package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/prithvi009/easedeploy/config"
)

func CloneRepository(repoURL, folderName string) error {

	destinationDir := filepath.Join("./repos", folderName)
	_, err := git.PlainClone(destinationDir, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "prithvi009",
			Password: config.GithubAccessToken(),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	fmt.Println("Successfully cloned repository to")

	return nil
}
