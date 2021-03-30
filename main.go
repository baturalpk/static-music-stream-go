package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			audioFile, err := ioutil.ReadFile("./audio.mp3")

			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "audio/mpeg")
			audioChunks := bytes.NewBuffer(audioFile)

			if _, err := audioChunks.WriteTo(w); err != nil {
				log.Println(err)
			}
			break
		}
	})

	var ServerPort = os.Getenv("PORT")

	if ServerPort == "" {
		ServerPort = "4444"
	}

	log.Printf("The server is streaming on http://localhost:%s/stream", ServerPort)
	log.Fatal(http.ListenAndServe(":" + ServerPort, router), nil)
}