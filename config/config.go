package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int64
}

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to Load env file", err)
		return
	}
	Version := os.Getenv("VERSION")
	if Version == "" {
		fmt.Println("Version is Required")
		os.Exit(1)
	}
	Service := os.Getenv("SERVICE_NAME")
	if Service == "" {
		fmt.Println("ServiceName is Required")
		os.Exit(1)
	}
	Httpport := os.Getenv("HTTP_PORT")
	if Httpport == "" {
		fmt.Println("Port is Required")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(Httpport, 10, 64)
	if err != nil {
		os.Exit(1)
	}
	configuration = Config{
		Version:     Version,
		ServiceName: Service,
		HttpPort:    port,
	}
}
func GetConfig() Config {
	Load()
	return configuration
}
