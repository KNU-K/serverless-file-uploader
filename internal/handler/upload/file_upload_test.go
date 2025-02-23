package upload

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHandler(t *testing.T) {
    body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("file", "test.txt")
    if err != nil {
        t.Fatalf("Error creating form file: %v", err)
    }
    part.Write([]byte("This is a test file"))
    writer.Close()

    req := httptest.NewRequest(http.MethodPost, "/upload", body)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    rr := httptest.NewRecorder()

    Handler(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := "File uploaded successfully"
    if rr.Body.String() != expected {
        t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }

    uploadedFilePath := filepath.Join(".", "uploads", "test.txt")
    if _, err := os.Stat(uploadedFilePath); os.IsNotExist(err) {
        t.Errorf("Expected file to be uploaded, but it does not exist")
    }

    os.Remove(uploadedFilePath)
}