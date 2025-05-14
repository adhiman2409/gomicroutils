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

func isValidSalarySlipFileName(s string) bool {
	// Pattern: employeeId (alphanumeric or numeric) + _ + 4-digit year + _ + 2-digit month (01–12)
	pattern := `^[a-zA-Z0-9]+_\d{4}_(0[1-9]|1[0-2])\.pdf$`
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func (a *StorageConnection) UploadSalarySlip(r *http.Request, domain string) (FileUploadResponse, error) {

	pathPrefix := "SalarySlips/"
	namePrefix := "Salary_Slip_"

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

	if !isValidSalarySlipFileName(fname) {
		return FileUploadResponse{}, errors.New("invalid file name format; it should be in the format: employeeId_YYYY_MM")
	}

	eid := strings.Split(fname, "_")[0]
	yearStr := strings.Split(fname, "_")[1]
	monthStr := strings.Split(fname, "_")[2]

	nd := GetUpdatedFinanceDomain(domain)

	storagePath := pathPrefix + eid + "/" + yearStr + "/" + namePrefix + eid + "_" + yearStr + "_" + monthStr + ".pdf"

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
