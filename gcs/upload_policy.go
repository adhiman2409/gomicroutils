package gcs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/martian/v3/log"
)

func (a *StorageConnection) UploadPolicy(r *http.Request, domain string) (FileUploadResponse, error) {

	pid := os.Getenv("GOOGLE_PROJECT_ID")

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

	folderName := "policies"
	fname := fh.Filename

	nd := strings.Replace(domain, ".", "_", -1)

	filepath := fmt.Sprintf("%s/%s", folderName, fname)

	wc := a.Client.Bucket(nd).UserProject(pid).Object(filepath).NewWriter(context.Background())
	if _, err = io.Copy(wc, f); err != nil {
		log.Errorf("file upload error: %v", err)
		return FileUploadResponse{}, err
	}
	if err := wc.Close(); err != nil {
		return FileUploadResponse{}, err
	}

	info := FileUploadResponse{
		Domain:  domain,
		DocName: fname,
		DocPath: filepath,
	}
	return info, nil
}
