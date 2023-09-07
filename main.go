package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

const FILE = "./data.json"

func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPort(port string) string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		return ":" + envPort
	}

	return ":" + port
}

type StringList []string

func main() {
	router := httprouter.New()
	data_file_content, err := os.ReadFile(FILE)
	checkNilError(err)

	var data_json StringList
	json.Unmarshal(data_file_content, &data_json)
	fmt.Printf("%T: %v \n", data_json[1], data_json[1])

	router.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.Write([]byte(data_json[1]))
	})

	port := getPort("8080")
	log.Fatal(http.ListenAndServe("0.0.0.0"+port, router))
}
