package env

import (
	"fmt"
	"os"

	"automatic-doodle/types"

	"github.com/joho/godotenv"
)

var GO_ENV types.GoEnv = types.GoEnvDev

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	GO_ENV = types.ParseGoEnv(os.Getenv("GO_ENV"))
}
