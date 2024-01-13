package utils

import (
	// import built-in packages
	"os"
)

func GetPort() string {
	// 设置端口
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	return port
}
