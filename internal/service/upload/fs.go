package upload_service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveToFS(file multipart.File, handler *multipart.FileHeader) error {
	uploadDir := filepath.Join(".", "uploads")
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
			return err
		}
	}

	dst, err := os.Create(filepath.Join(uploadDir, handler.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	return nil
}