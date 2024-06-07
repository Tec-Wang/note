package config

// DBConfig contains the database configuration
type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func MySQLConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "password",
		Database: "mydatabase",
	}
}

func (c *DBConfig) ConnectionString() string {
	return c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}
