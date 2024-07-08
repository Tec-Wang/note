package file_storage

import "os"

type LocalFileStorageService struct {
	config LocalStorageServiceConfig
}

type LocalStorageServiceConfig struct {
	RootDirectory string `json:"RootDirectory"`
}

func NewLocalFileStorageService(c LocalStorageServiceConfig) *LocalFileStorageService {
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
