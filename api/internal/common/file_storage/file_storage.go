package file_storage

type FileStorage struct {
}

type FileStorageService interface {
	UploadFile(fileName string, fileContent []byte) error
	GetFileContent(fileName string) ([]byte, error)
}
