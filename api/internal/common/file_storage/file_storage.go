package file_storage

import "io"

type FileStorageService interface {
	UploadFile(fileName string, fileContent io.Reader) error
	GetFileContent(fileName string) ([]byte, error)
}
