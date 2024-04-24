package gcs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (a *StorageConnection) UploadFile(r *http.Request, domain string) (FileUploadResponse, error) {
	department := r.FormValue("department")
	eid := r.FormValue("eid")
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

	nd := strings.Replace(domain, ".", "_", -1)

	filepath := fmt.Sprintf("unirms/%s/%s/%s/%s/%s", nd, department, eid, documentType, fname)

	if err = a.UploadFileToGCSBucket(context.Background(), fh, filepath); err != nil {
		fmt.Println("Error ", err.Error())
		return FileUploadResponse{}, errors.New("unable to process input file")
	}

	info := FileUploadResponse{
		Name:         fname,
		Path:         filepath,
		DocumentType: documentType,
	}
	return info, nil
}
