package gcs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/google/martian/v3/log"
)

func isValidAppraissalLetterFileName(s string) bool {
	// Pattern: employeeId (alphanumeric or numeric) + _ + 4-digit year
	pattern := `^[a-zA-Z0-9]+_\d{4}$`
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func (a *StorageConnection) UploadAppraisalLetter(r *http.Request, domain string) (FileUploadResponse, error) {

	pathPrefix := "AppraisalLetter/"
	namePrefix := "Appraisal_Letter_"

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

	fname := fh.Filename

	if !isValidAppraissalLetterFileName(fname) {
		return FileUploadResponse{}, errors.New("invalid file name format; it should be in the format: employeeId_YYYY")
	}

	eid := strings.Split(fname, "_")[0]
	yearStr := strings.Split(fname, "_")[1]

	nd := GetUpdatedFinanceDomain(domain)

	storagePath := pathPrefix + eid + "/" + yearStr + "/" + namePrefix + eid + "_" + yearStr + ".pdf"

	wc := a.Client.Bucket(nd).UserProject(pid).Object(storagePath).NewWriter(context.Background())
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
		DocPath: storagePath,
	}
	return info, nil
}
