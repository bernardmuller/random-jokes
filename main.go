package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func getPort(port string) string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		return ":" + envPort
	}

	return ":" + port
}

func main() {
	router := httprouter.New()

	port := getPort("8080")
	log.Fatal(http.ListenAndServe("0.0.0.0"+port, router))
}
