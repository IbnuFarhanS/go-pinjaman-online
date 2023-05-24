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
	c.ConnectionString = utility.GetEnv("CONNECTION_STRING", "host=localhost user=postgres password=sql1234 dbname=pinjaman_online port=5432 sslmode=disable")
	c.ServerAddress = utility.GetEnv("SERVER_ADDRESS", "localhost:8080")
}
