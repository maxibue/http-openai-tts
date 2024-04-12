package main

import (
	"log"
	"net/http"

	"github.com/maximierung/http-openai-tts/api"
	"github.com/maximierung/http-openai-tts/database"
	"github.com/maximierung/http-openai-tts/utils"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	client := database.NewClient(config.MongoURI)
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		api.Ping(w, r)
	})

	if config.AllowAdmin {
		mux.HandleFunc("/admin/{type}", func(w http.ResponseWriter, r *http.Request) {
			api.AdminRequest(client, w, r)
		})
	}

	if config.AllowHosting {
		mux.HandleFunc("/tts", func(w http.ResponseWriter, r *http.Request) {
			api.SendRequest(client, w, r)
		})

		mux.HandleFunc("/raw", func(w http.ResponseWriter, r *http.Request) {
			api.SendRawRequest(client, w, r)
		})

		mux.Handle("/", http.FileServer(http.Dir("./output/")))
	} else {
		mux.HandleFunc("/tts", func(w http.ResponseWriter, r *http.Request) {
			api.SendRawRequest(client, w, r)
		})
	}

	log.Println("Starting server on port: " + config.ServerPort)
	if err := http.ListenAndServe(":"+config.ServerPort, mux); err != nil {
		log.Fatal(err)
	}
}
