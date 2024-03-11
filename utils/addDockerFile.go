package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func AddDockerfiles(repoPath string) error {

	currDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	backendPath := filepath.Join(currDir, "repos", repoPath, "backend")
	frontendPath := filepath.Join(currDir, "repos", repoPath, "frontend")

	if _, err := os.Stat(backendPath); err == nil {
		if err := addDockerfile(backendPath, "backend"); err != nil {
			return err
		}
	}

	if _, err := os.Stat(frontendPath); err == nil {
		if err := addDockerfile(frontendPath, "frontend"); err != nil {
			return err
		}
	}

	return nil
}

func addDockerfile(folderPath, folderName string) error {
	dockerfilePath := filepath.Join(folderPath, "Dockerfile")
	if _, err := os.Stat(dockerfilePath); err == nil {
		fmt.Printf("Dockerfile already exists in %s\n", folderPath)
		return nil
	}

	var dockerfileContent string
	switch folderName {
	case "backend":
		dockerfileContent = `
FROM node:alpine

WORKDIR /app

COPY package.json .
COPY package-lock.json .
RUN npm install --no-cache

COPY . .

EXPOSE 5000

CMD ["npm", "start"]
`
	case "frontend":
		dockerfileContent = `
FROM node:alpine

WORKDIR /app

COPY package.json .
COPY package-lock.json .
RUN npm install --no-cache

COPY . .

EXPOSE 3000

CMD ["npm", "start"]
`
	default:
		return fmt.Errorf("unsupported folder: %s", folderName)
	}

	if err := os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644); err != nil {
		return fmt.Errorf("failed to write Dockerfile: %w", err)
	}

	fmt.Printf("Dockerfile created in %s\n", folderPath)
	return nil
}
