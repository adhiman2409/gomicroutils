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

func (a *StorageConnection) UploadTicketData(r *http.Request, domain string) (TicketFileUploadResponse, error) {

	pid := os.Getenv("GOOGLE_PROJECT_ID")
	department := "Ticket"
	ticket_id := r.FormValue("ticket_id")
	documentType := r.FormValue("dtype")

	var ticketPathArr []string
	var fNameArr [] string 
	
	err := r.ParseMultipartForm(50 << 20) // Max Size Limit is 50 MB
	if err != nil {
		fmt.Println("Error ", err.Error())
		return TicketFileUploadResponse{}, errors.New("size is > 50 mb")
	}

	for _, fh := range r.MultipartForm.File["ticket_images"] {
		f,  err := fh.Open()
		if err != nil {
			fmt.Println("Error ", err.Error())
			return TicketFileUploadResponse{}, errors.New("file read error")
		}
		defer f.Close()

		fname := ticket_id + "_" + fh.Filename
		fNameArr = append(fNameArr, fname)

		nd := strings.Replace(domain, ".", "_", -1)

		filepath := fmt.Sprintf("%s/%s/%s/%s", department, ticket_id, documentType, fname)
		ticketPathArr = append(ticketPathArr, filepath)

		wc := a.Client.Bucket(nd).UserProject(pid).Object(filepath).NewWriter(context.Background())
		if _, err = io.Copy(wc, f); err != nil {
			log.Errorf("file upload error: %v", err)
			return TicketFileUploadResponse{}, err
		}
		if err := wc.Close(); err != nil {
			return TicketFileUploadResponse{}, err
		}
	}

	info := TicketFileUploadResponse{
		Domain:      domain,
		TicketId: 	 ticket_id,
		Department:  department,
		DocName:     fNameArr,
		DocType:     documentType,
		TicketPathArr: ticketPathArr,
	}
	return info, nil
}
