package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipFolder(folderName string) error {
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
