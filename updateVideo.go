package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
)

func updateVideo(videoUrl, userToken string, w http.ResponseWriter) {
	// Check if user exists
	allowed, err := extensions.AuthRequestAndDecision("http://" + sidecarUrl +
		"/auth?token=" + userToken + "&service=updateVideo")
	if !allowed {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err != nil {
		log.Println("Could not update video: ", err)
		fmt.Fprintf(w, "Could not update video")
		return
	}

	// Change video
	realUrl, err := url.Parse(videoUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Could not parse URL:"
		log.Println(msg, err)
		fmt.Fprintln(w, msg, err)
		return
	}

	if realUrl.Host != "www.youtube.com" ||
		realUrl.Path != "/watch" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Not a YouTube video")
		return
	}

	videoId := realUrl.Query().Get("v")

	if len(videoId) != 11 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid video ID")
		return
	}

	videoFile, err := os.Create(videoPath)
	defer videoFile.Close()

	msg := "Could store video:"
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(msg, err)
		fmt.Fprintln(w, msg, err)
		return
	}

	_, err = videoFile.WriteString(videoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(msg, err)
		fmt.Fprintln(w, msg, err)
		return
	}

	fmt.Fprintf(w, "Stored video successfully")
}
