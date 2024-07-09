package config

type LocalStorageServiceConfig struct {
	RootDirectory string `json:"RootDirectory"`
}

func (c *LocalStorageServiceConfig) GetPath(path string) string {
	return c.RootDirectory + "/" + path
}
