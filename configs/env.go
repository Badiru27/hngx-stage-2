package configs

import (
	//"log"
	"os"
	//"github.com/joho/godotenv"
)

func EnvMongoURI() string{
	// error := godotenv.Load();

	// if error != nil {
	// 	log.Fatal("Error loading ENV file")
	// }

	return os.Getenv("MONGOURI")
}