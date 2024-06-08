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
		Password: "123456",
		Database: "note",
	}
}

func (c *DBConfig) ConnectionString() string {
	return c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func MongoConfig() *MongoDBConfig {
	return &MongoDBConfig{
		Host:     "localhost",
		Port:     "27017",
		User:     "mongoadmin",
		Password: "123456",
		Database: "note",
	}
}

type MongoDBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (c *MongoDBConfig) ConnectionString() string {
	return "mongodb://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.Database
}
