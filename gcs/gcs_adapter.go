package gcs

import (
	"context"
	"io"
	"mime/multipart"
	"os"

	"github.com/google/martian/v3/log"
)

// UploadFileToGCSBucket will create a date wise directory bucket
func (c *StorageConnection) UploadFileToGCSBucket(ctx context.Context, fh *multipart.FileHeader, filepath string) error {

	file, err := fh.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	b := os.Getenv("GOOGLE_STORAGE_BUCKET")
	pid := os.Getenv("GOOGLE_PROJECT_ID")
	wc := c.Client.Bucket(b).UserProject(pid).Object(filepath).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		log.Errorf("file upload error: %v", err)
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}
