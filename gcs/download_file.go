package gcs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Download gets a file from GCS bucket, Takes file path as a path param from request
func (a *StorageConnection) DownloadFile(w http.ResponseWriter, r *http.Request, domain string) error {
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	clientCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	department := mux.Vars(r)["department"]
	eid := mux.Vars(r)["eid"]
	category := mux.Vars(r)["category"]
	documentType := mux.Vars(r)["dtype"]
	filename := mux.Vars(r)["filename"]
	nd := GetUpdatedDomain(domain)
	filePath := fmt.Sprintf("%s/%s/%s/%s/%s", department, eid, category, documentType, filename)

	reader, err := a.Client.Bucket(nd).UserProject(pid).Object(filePath).NewReader(clientCtx)
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

// Download gets a file from GCS bucket, Takes file path as a path param from request
func (a *StorageConnection) DownloadImage(w http.ResponseWriter, r *http.Request) error {
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	clientCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var req ImageDownloadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	nd := GetUpdatedDomain(req.Domain)
	filePath := fmt.Sprintf("%s/%s/%s/%s/%s", req.Department, req.EID, req.Category, req.IType, req.Name)

	reader, err := a.Client.Bucket(nd).UserProject(pid).Object(filePath).NewReader(clientCtx)
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
	disposition := "inline"
	w.Header().Set("Content-Disposition", disposition+"; filename="+req.Name)
	w.Header().Set("Content-Length", size)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

	return nil
}

func (a *StorageConnection) DownloadStaticFile(w http.ResponseWriter, r *http.Request, domain string) error {
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	clientCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filename := mux.Vars(r)["filename"]
	nd := GetUpdatedDomain(domain)
	filepathwithname := "static/" + filename

	reader, err := a.Client.Bucket(nd).UserProject(pid).Object(filepathwithname).NewReader(clientCtx)
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

func (a *StorageConnection) DownloadPolicy(w http.ResponseWriter, r *http.Request, domain string) error {
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	clientCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filename := mux.Vars(r)["filename"]
	nd := GetUpdatedDomain(domain)
	filepathwithname := "policies/" + filename

	reader, err := a.Client.Bucket(nd).UserProject(pid).Object(filepathwithname).NewReader(clientCtx)
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
