package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	TOKEN string
	// BOTNAME string
	// PREFIX     string
	// DATABASE   string
	// DBUSERNAME string
	// DBPASSWORD string
	// DBHOST     string
	// DBNAME     string
	// DBPORT     string
	// DBSSL      string
)

func Env() {

	//LOAD ENV DATA
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	uToken, check := os.LookupEnv("TOKEN")
	if !check {
		fmt.Println("No Token defined in env vars.")
	}
	TOKEN = uToken

	// uBotname, check := os.LookupEnv("BOTNAME")
	// if !check {
	// 	fmt.Println("No Botname defined in env vars.")
	// }
	// BOTNAME = uBotname

	// uPrefix, check := os.LookupEnv("PREFIX")
	// if !check {
	// 	fmt.Println("No Prefix defined in env vars.")
	// }
	// PREFIX = uPrefix

	// db, check := os.LookupEnv("DATABASE")
	// if !check {
	// 	fmt.Println("No Database defined in env vars.")
	// }
	// DATABASE = db

	// dbusername, check := os.LookupEnv("DBUSERNAME")
	// if !check {
	// 	fmt.Println("No Database username defined in env vars.")
	// }
	// DBUSERNAME = dbusername

	// dbpassword, check := os.LookupEnv("DBPASSWORD")
	// if !check {
	// 	fmt.Println("No Database password defined in env vars.")
	// }
	// DBPASSWORD = dbpassword

	// dbhost, check := os.LookupEnv("DBHOST")
	// if !check {
	// 	fmt.Println("No Database host defined in env vars.")
	// }
	// DBHOST = dbhost

	// dbname, check := os.LookupEnv("DBNAME")
	// if !check {
	// 	fmt.Println("No Database name defined in env vars.")
	// }
	// DBNAME = dbname

	// dbport, check := os.LookupEnv("DBPORT")
	// if !check {
	// 	fmt.Println("No Database port defined in env vars.")
	// }
	// DBPORT = dbport

	// dbssl, check := os.LookupEnv("DBSSL")
	// if !check {
	// 	fmt.Println("No Database ssl defined in env vars.")
	// }
	// DBSSL = dbssl
}
