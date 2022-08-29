package main

import (
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const videoPath = "/var/cinema/video.txt"

var (
	sidecarUrl string
)

func handleCinemaRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch new URL if handed over
	url := r.URL.Query().Get("url")
	userToken := r.URL.Query().Get("token")

	if len(url) > 0 {
		updateVideo(url, userToken, w)
		return
	}
	showVideo(userToken, w)
}

func main() {
	log.SetFlags(0)
	var exists bool

	sidecarUrl, exists = os.LookupEnv("SIDECAR_URL")
	if !exists {
		log.Fatalln("No sidecar URL set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handleCinemaRequest)

	log.Println("Cinema is listening on port 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("Cinema failed:", err)
	}
}
