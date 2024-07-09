package file_storage

import (
	"io"
	"os"
	"wangzheng/brain/api/internal/config"
)

type LocalFileStorageService struct {
	config config.LocalStorageServiceConfig
}

func NewLocalFileStorageService(c config.LocalStorageServiceConfig) *LocalFileStorageService {
	if c.RootDirectory == "" {
		c.RootDirectory = "./"
	}

	// create root directory if not exists
	if _, err := os.Stat(c.RootDirectory); os.IsNotExist(err) {
		_ = os.Mkdir(c.RootDirectory, os.ModePerm)
	}
	return &LocalFileStorageService{config: c}
}

func (s *LocalFileStorageService) UploadFile(fileName string, fileContent io.Reader) error {
	file, err := os.Create(s.config.GetPath(fileName))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, fileContent)

	if err != nil {
		return err
	}
	return nil
}

func (s *LocalFileStorageService) GetFileContent(fileName string) ([]byte, error) {
	filePath := s.config.RootDirectory + "/" + fileName
	return os.ReadFile(filePath)
}
