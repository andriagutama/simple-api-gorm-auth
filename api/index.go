package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv" // go get -u github.com/joho/godotenv
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal("Error env")
	}
	fmt.Println("POSTGRES_URL :", os.Getenv("POSTGRES_URL"))
}
