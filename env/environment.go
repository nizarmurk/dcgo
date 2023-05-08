package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	TOKEN   string
	PREFIX  string
	BOTNAME string
)

func getEnvVariable(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println(defaultValue)
	}
	return value
}

func Env() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	TOKEN = getEnvVariable("TOKEN", "No Token defined in env vars.")
	BOTNAME = getEnvVariable("BOTNAME", "No Botname defined in env vars.")
	// PREFIX = getEnvVariable("PREFIX", "No Prefix defined in env vars.")
	// DATABASE = getEnvVariable("DATABASE", "No Database defined in env vars.")
	// DBUSERNAME = getEnvVariable("DBUSERNAME", "No Database username defined in env vars.")
	// DBPASSWORD = getEnvVariable("DBPASSWORD", "No Database password defined in env vars.")
	// DBHOST = getEnvVariable("DBHOST", "No Database host defined in env vars.")
	// DBNAME = getEnvVariable("DBNAME", "No Database name defined in env vars.")
	// DBPORT = getEnvVariable("DBPORT", "No Database port defined in env vars.")
	// DBSSL = getEnvVariable("DBSSL", "No Database ssl defined in env vars.")
}
