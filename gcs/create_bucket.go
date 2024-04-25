package gcs

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func (a *StorageConnection) CreateBucket(domain string) error {
	ctx := context.Background()

	pid := os.Getenv("GOOGLE_PROJECT_ID")

	bucketName := strings.Replace(domain, ".", "_", -1)

	// Setup client bucket to work from
	bucket := a.Client.Bucket(bucketName)

	buckets := a.Client.Buckets(ctx, pid)
	for {
		if bucketName == "" {
			return fmt.Errorf("bucket name entered is empty %v", bucketName)
		}
		attrs, err := buckets.Next()
		// Assume bucket not found if at Iterator end and create
		if err == iterator.Done {
			// Create bucket
			if err := bucket.Create(ctx, pid, &storage.BucketAttrs{
				Location: "US",
			}); err != nil {
				return fmt.Errorf("failed to create bucket: %v", err)
			}
			log.Printf("Bucket %v created.\n", bucketName)
			return nil
		}
		if err != nil {
			return fmt.Errorf("issues setting up bucket(%q).objects(): %v. double check project id", attrs.Name, err)
		}
		if attrs.Name == bucketName {
			log.Printf("Bucket %v exists.\n", bucketName)
			return nil
		}
	}
}
