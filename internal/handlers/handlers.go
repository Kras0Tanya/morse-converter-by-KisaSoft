package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Kras0Tanya/morse-converter-by-KisaSoft/internal/service"
)

// я всё ещё плачу :`(

func MainHandler(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to determine working directory: %v", err), http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(wd, "../index.html")
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read index.html: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // на это своего мозга уже не хватило, пришлось импров... заимствовать
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "File not found in form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusInternalServerError)
		return
	}

	converted, err := service.DetectAndConvert(string(content))
	if err != nil {
		http.Error(w, "Failed to convert file content", http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	ext := filepath.Ext(header.Filename)
	newFileName := fmt.Sprintf("converted_%s%s", timestamp, ext)

	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "Failed to create result file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(converted)
	if err != nil {
		http.Error(w, "Failed to write to result file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
