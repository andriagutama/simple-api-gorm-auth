package api

import (
	"fmt"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	fmt.Println("POSTGRES_URL :", os.Getenv("POSTGRES_URL"))
}
