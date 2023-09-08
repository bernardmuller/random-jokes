package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

type Joke struct {
	Joke string `json:"joke"`
}

func main() {
	router := httprouter.New()
	data_file_content, err := os.ReadFile(FILE)
	checkNilError(err)

	var data_json StringList
	json.Unmarshal(data_file_content, &data_json)
	fmt.Printf("%T: %v \n", data_json[1], data_json[1])

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		randomJoke := data_json[rand.Intn(len(data_json))]
		query := r.URL.Query()
		format := query["format"]
		if format != nil && format[0] == "json" {
			joke := Joke{Joke: randomJoke}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(joke)
			return
		}
		w.Write([]byte(randomJoke))
	})

	port := getPort("8080")
	log.Fatal(http.ListenAndServe("0.0.0.0"+port, router))
}
