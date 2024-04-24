package gcs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Download gets a file from GCS bucket, Takes file path as a path param from request
func (a *StorageConnection) DownloadFile(w http.ResponseWriter, r *http.Request, domain string) error {

	clientCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	department := mux.Vars(r)["department"]
	eid := mux.Vars(r)["eid"]
	documentType := mux.Vars(r)["dtype"]
	filename := mux.Vars(r)["filename"]
	nd := strings.Replace(domain, ".", "_", -1)
	filePath := fmt.Sprintf("unirms/%s/%s/%s/%s/%s", nd, department, eid, documentType, filename)

	b := os.Getenv("GOOGLE_STORAGE_BUCKET")
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	reader, err := a.Client.Bucket(b).UserProject(pid).Object(filePath).NewReader(clientCtx)
	if err != nil {
		fmt.Println("Error ", err.Error())
		return err
	}
	defer reader.Close()
	contentType := reader.Attrs.ContentType
	size := strconv.FormatInt(reader.Attrs.Size, 10)
	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error ", err.Error())
		return err
	}
	w.Header().Set("Content-Type", contentType)
	disposition := "attachment"
	if filename == "thumbnail.jpg" {
		disposition = "inline"
	}
	w.Header().Set("Content-Disposition", disposition+"; filename="+filename)
	w.Header().Set("Content-Length", size)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

	return nil
}
