package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadEnv(key string) string {
	typeEnv := os.Args[1]
	if typeEnv != "dev" && typeEnv != "prod" && typeEnv != "test" {
		typeEnv = "dev"
	}
	env := fmt.Sprintf(".%s.env", typeEnv)
	viper.SetConfigFile((env))

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file%s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion %s", key)
	}

	return value
}
