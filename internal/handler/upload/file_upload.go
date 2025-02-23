package upload

import (
	"net/http"
	upload_service "serverless-file-uploader/internal/service/upload"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB max memory
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    if err := upload_service.SaveToFS(file, handler); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("File uploaded successfully"))
}