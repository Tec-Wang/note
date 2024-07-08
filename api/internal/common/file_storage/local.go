package file_storage

import "os"

type LocalFileStorageService struct {
	config FileStorageServiceConfig
}

type FileStorageServiceConfig struct {
	RootDirectory string
}

func NewLocalFileStorageService(c FileStorageServiceConfig) *LocalFileStorageService {
	return &LocalFileStorageService{config: c}
}

func (s *LocalFileStorageService) UploadFile(fileName string, fileContent []byte) error {
	filePath := s.config.RootDirectory + "/" + fileName
	return os.WriteFile(filePath, fileContent, 0644)
}

func (s *LocalFileStorageService) GetFileContent(fileName string) ([]byte, error) {
	filePath := s.config.RootDirectory + "/" + fileName
	return os.ReadFile(filePath)
}
