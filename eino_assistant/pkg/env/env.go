package env

import (
	"log"
	"os"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("❌ [ERROR] Error loading .env file: %v", err)
// 	}

// }

func MustHasEnvs(envs ...string) {
	for _, env := range envs {
		if os.Getenv(env) == "" {
			log.Fatalf("❌ [ERROR] env [%s] is required, but is not set now, please check your .env file", env)
		}
	}
}
