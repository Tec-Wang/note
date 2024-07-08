package svc

import (
	"wangzheng/brain/api/internal/common/file_storage"
	"wangzheng/brain/api/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	FileStorage file_storage.FileStorageService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		FileStorage: file_storage.NewLocalFileStorageService(c.LocalStorageServiceConfig),
	}
}
