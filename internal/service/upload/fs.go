package upload_service

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveToFS(file multipart.File, handler *multipart.FileHeader) error {
    uploadDir := filepath.Join(".", "uploads")
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
            return err
        }
    }

    // Generate a random filename
    randomFilename := generateRandomFilename(handler.Filename)

    dst, err := os.Create(filepath.Join(uploadDir, randomFilename))
    if err != nil {
        return err
    }
    defer dst.Close()

    if _, err := io.Copy(dst, file); err != nil {
        return err
    }

    return nil
}

func generateRandomFilename(originalFilename string) string {
    timestamp := time.Now().Format("20060102150405")
    randomBytes := make([]byte, 8)
    rand.Read(randomBytes)
    randomString := fmt.Sprintf("%x", randomBytes)
    extension := filepath.Ext(originalFilename)
    return fmt.Sprintf("%s_%s%s", timestamp, randomString, extension)
}