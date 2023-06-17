package api

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	//fmt.Fprintf(w, "Postgresql : %s", os.Getenv("POSTGRES_URL"))
}
