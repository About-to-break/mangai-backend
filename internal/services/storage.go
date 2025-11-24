package services

import "io"

type StorageService interface {
	UploadFile(bucket string, objectName string, reader io.Reader, size int64, contentType string) error
}
