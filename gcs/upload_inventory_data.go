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

func (a *StorageConnection) UploadInventoryData(r *http.Request, domain string) (InventoryFileUploadResponse, error) {

	pid := os.Getenv("GOOGLE_PROJECT_ID")
	department := "Inventory"
	bucket_id := r.FormValue("bucket_id")
	store_id := r.FormValue("store_id")
	documentType := r.FormValue("dtype")

	var productPathArr []string
	var invoicePathArr []string
	var fNameArr [] string 
	
	err := r.ParseMultipartForm(50 << 20) // Max Size Limit is 50 MB
	if err != nil {
		fmt.Println("Error ", err.Error())
		return InventoryFileUploadResponse{}, errors.New("size is > 50 mb")
	}

	for _, fh := range r.MultipartForm.File["product_images"] {
		f,  err := fh.Open()
		if err != nil {
			fmt.Println("Error ", err.Error())
			return InventoryFileUploadResponse{}, errors.New("file read error")
		}
		defer f.Close()

		fname := store_id + "_"+bucket_id + "_" + fh.Filename
		fNameArr = append(fNameArr, fname)

		nd := strings.Replace(domain, ".", "_", -1)

		filepath := fmt.Sprintf("%s/%s/%s/%s/%s", department, bucket_id, store_id, documentType, fname)
		productPathArr = append(productPathArr, filepath)

		wc := a.Client.Bucket(nd).UserProject(pid).Object(filepath).NewWriter(context.Background())
		if _, err = io.Copy(wc, f); err != nil {
			log.Errorf("file upload error: %v", err)
			return InventoryFileUploadResponse{}, err
		}
		if err := wc.Close(); err != nil {
			return InventoryFileUploadResponse{}, err
		}
	}

	for _, fh := range r.MultipartForm.File["invoices"] {
		f,  err := fh.Open()
		if err != nil {
			fmt.Println("Error ", err.Error())
			return InventoryFileUploadResponse{}, errors.New("file read error")
		}
		f.Close()

		fname := store_id + "_"+bucket_id + "_" + fh.Filename
		fNameArr = append(fNameArr, fname)
		
		nd := strings.Replace(domain, ".", "_", -1)

		filepath := fmt.Sprintf("%s/%s/%s/%s/%s", department, bucket_id, store_id, documentType, fname)
		invoicePathArr = append(invoicePathArr, filepath)

		wc := a.Client.Bucket(nd).UserProject(pid).Object(filepath).NewWriter(context.Background())
		if _, err = io.Copy(wc, f); err != nil {
			log.Errorf("file upload error: %v", err)
			return InventoryFileUploadResponse{}, err
		}
		if err := wc.Close(); err != nil {
			return InventoryFileUploadResponse{}, err
		}
	}

	info := InventoryFileUploadResponse{
		Domain:      domain,
		BucketId: 	 bucket_id,
		StoreId: 	 store_id,
		Department:  department,
		DocName:     fNameArr,
		DocType:     documentType,
		ProductImagePathArr:     productPathArr,
		InvoicePathArr:     invoicePathArr,
	}
	return info, nil
}
