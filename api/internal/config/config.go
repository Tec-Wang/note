package config

import (
	"github.com/zeromicro/go-zero/rest"
	"wangzheng/brain/api/internal/common/file_storage"
)

type Config struct {
	rest.RestConf
	file_storage.LocalStorageServiceConfig `yaml:"LocalStorageServiceConfig"`
}
