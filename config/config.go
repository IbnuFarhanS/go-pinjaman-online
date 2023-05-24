package config

import (
	utility "github.com/IbnuFarhanS/go-pinjaman-online/utils"
)

// var db *gorm.DB

type config struct {
	ConnectionString string
	ServerAddress    string
}

func NewConfig() *config {
	return &config{}
}

func (c *config) Load() {
	// var err error
	c.ConnectionString = utility.GetEnv("CONNECTION_STRING", "host=localhost user=postgres password=sql1234 dbname=pinjaman_online port=5432 sslmode=disable")

	// dsn := utility.GetEnv("CONNECTION_STRING")
	// c.ConnectionString = dsn
	// db, err = gorm.Open(postgres.Open(c.ConnectionString), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect to database")
	// }
	c.ServerAddress = utility.GetEnv("SERVER_ADDRESS", "localhost:8080")
}
