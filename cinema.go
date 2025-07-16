package main

import (
	"html"
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	sidecarUrl string
	video      string
)

func handleCinemaRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch new URL if handed over
	url := r.URL.Query().Get("url")
	userToken := r.URL.Query().Get("token")

	// escape parameters
	url = html.EscapeString(url)
	userToken = html.EscapeString(userToken)

	if len(url) > 0 {
		updateVideo(url, userToken, w)
		return
	}
	showVideo(userToken, w)
}

func main() {
	log.SetFlags(0)
	var exists bool

	// Set SNZR video
	video = "aWKusTpZ3FI"

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
