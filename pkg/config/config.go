package config

import (
	// import built-in packages
	"log"
	"os"

	// import third-party packages
	"github.com/dotenv-org/godotenvvault"

	// import local packages
	"brimstone/global"
)

var Conf *Config

func Init() {
	// 加载环境变量
	err := godotenvvault.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database_client := os.Getenv("DATABASE_CLIENT")

	if host == "" {
		host = global.HOST
	}

	if port == "" {
		port = global.PORT
	}

	if database_client == "" {
		database_client = global.DATABASE_CLIENT
	}

	Conf = &Config{
		Port: port,
	}
}

type Config struct {
	Host string
	Port string
}
