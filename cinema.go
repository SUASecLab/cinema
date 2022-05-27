package main

import (
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const videoPath = "/var/cinema/video.txt"

var adminExtensions string

func handleCinemaRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch new URL if handed over
	url := r.URL.Query().Get("url")
	uuid := r.URL.Query().Get("uuid")

	if len(url) > 0 {
		updateVideo(url, uuid, w)
		return
	}
	showVideo(uuid, w)
}

func main() {
	log.SetFlags(0)

	adminExtensions = os.Getenv("ADMIN_EXTENSIONS")

	r := mux.NewRouter()
	r.HandleFunc("/", handleCinemaRequest)

	log.Println("Cinema is listening on port 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("Cinema failed:", err)
	}
}
