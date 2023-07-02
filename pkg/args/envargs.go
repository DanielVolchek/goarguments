package args

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvArgs(args ...string) []string {

	godotenv.Load(".env")

	if len(args) == 0 {
		return nil
	}

	parsed := []string{}

	for _, arg := range args {
		parsed = append(parsed, os.Getenv(arg))
	}

	return parsed
}
