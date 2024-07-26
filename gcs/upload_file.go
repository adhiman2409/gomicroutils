package gcs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/martian/v3/log"
)

func (a *StorageConnection) UploadFile(r *http.Request, domain string) (FileUploadResponse, error) {

	pid := os.Getenv("GOOGLE_PROJECT_ID")
	department := r.FormValue("department")
	eid := r.FormValue("eid")
	category := r.FormValue("category")
	documentType := r.FormValue("dtype")

	err := r.ParseMultipartForm(50 << 20) // Max Size Limit is 50 MB
	if err != nil {
		fmt.Println("Error ", err.Error())
		return FileUploadResponse{}, errors.New("size is > 50 mb")
	}

	f, fh, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error ", err.Error())
		return FileUploadResponse{}, errors.New("file read error")
	}
	defer f.Close()

	fname := eid + "_" + fh.Filename

	nd := GetUpdatedDomain(domain)

	filepath := fmt.Sprintf("%s/%s/%s/%s/%s", department, eid, category, documentType, fname)

	wc := a.Client.Bucket(nd).UserProject(pid).Object(filepath).NewWriter(context.Background())
	if _, err = io.Copy(wc, f); err != nil {
		log.Errorf("file upload error: %v", err)
		return FileUploadResponse{}, err
	}
	if err := wc.Close(); err != nil {
		return FileUploadResponse{}, err
	}

	info := FileUploadResponse{
		EID:         eid,
		Domain:      domain,
		Department:  department,
		DocName:     fname,
		DocCategory: category,
		DocType:     documentType,
		DocPath:     filepath,
	}
	return info, nil
}
