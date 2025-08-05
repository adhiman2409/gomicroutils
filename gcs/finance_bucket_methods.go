package gcs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func (a *StorageConnection) DownloadSalarySlip(w http.ResponseWriter, employeeId, year, month, domain string) error {
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	clientCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd := GetUpdatedFinanceDomain(domain)
	fileName := "Salary_Slip_" + employeeId + "_" + year + "_" + month + ".pdf"

	filepathwithname := "SalarySlips/" + employeeId + "/" + year + "/" + fileName
	if month == "2025" {
		fileName = "Form_16_" + employeeId + "_2024_2025.pdf"
		filepathwithname = "SalarySlips/" + employeeId + "/2025/" + fileName
	}

	fmt.Println("filepathwithname ", filepathwithname)
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
	disposition := "inline"
	w.Header().Set("Content-Disposition", disposition+"; filename="+fileName)
	w.Header().Set("Content-Length", size)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

	return nil
}

func (a *StorageConnection) DownloadAppraisalLetter(w http.ResponseWriter, employeeId, year, domain string) error {
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	clientCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd := GetUpdatedFinanceDomain(domain)
	filepathwithname := "IncrementLetters/" + employeeId + "/" + year + "/" + "Increment_Letter_" + employeeId + "_" + year + ".pdf"
	fmt.Println("filepathwithname ", filepathwithname)
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
	disposition := "inline"
	w.Header().Set("Content-Disposition", disposition+"; filename="+"Appraisal_Letter_"+employeeId+"_"+year+".pdf")
	w.Header().Set("Content-Length", size)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

	return nil
}

func (a *StorageConnection) GetAppraisalLetterByEID(employeeId, domain string) ([]string, error) {
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	clientCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd := GetUpdatedFinanceDomain(domain)
	prefix := "IncrementLetters/" + employeeId + "/"
	delim := ""

	it := a.Client.Bucket(nd).UserProject(pid).Objects(clientCtx, &storage.Query{
		Prefix:    prefix,
		Delimiter: delim,
	})
	res := []string{}
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("GetAppraisalLetterByEID Error: %v\n", err)
			continue
		}
		if strings.Contains(attrs.Name, ".pdf") {
			res = append(res, attrs.Name)
		}
	}
	return res, nil

}

func (a *StorageConnection) GetSalarySlipsByEID(employeeId, domain string) ([]string, error) {
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	clientCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd := GetUpdatedFinanceDomain(domain)
	prefix := "SalarySlips/" + employeeId + "/"
	delim := ""

	it := a.Client.Bucket(nd).UserProject(pid).Objects(clientCtx, &storage.Query{
		Prefix:    prefix,
		Delimiter: delim,
	})
	res := []string{}
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("GetSalarySlipsByEID Error: %v\n", err)
			continue
		}
		if strings.Contains(attrs.Name, ".pdf") {
			res = append(res, attrs.Name)
		}
	}
	return res, nil

}
