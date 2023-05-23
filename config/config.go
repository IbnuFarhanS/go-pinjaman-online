package config

import utility "github.com/IbnuFarhanS/go-pinjaman-online/utils"

type config struct {
	ConnectionString string
	ServerAddress    string
}

func NewConfig() *config {
	return &config{}
}

func (c *config) Load() {
	c.ConnectionString = utility.GetEnv("CONNECTION_STRING", "host=localhost user=postgres password=aingmaung26 dbname=user_management_db port=5432 sslmode=disable")
	c.ServerAddress = utility.GetEnv("SERVER_ADDRESS", "localhost:8080")
}