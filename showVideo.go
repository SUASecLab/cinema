package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
)

func showVideo(userToken string, w http.ResponseWriter) {
	// Check if user exists
	allowed, err := extensions.AuthRequestAndDecision("http://" + sidecarUrl +
		"/auth?token=" + userToken + "&service=showVideo")
	if !allowed {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err != nil {
		errorMsg := "Could not get authentication decision"
		log.Println(errorMsg, err)
		fmt.Fprintln(w, errorMsg)
		return
	}

	video, err := os.ReadFile(videoPath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("No video stored:", err)
		return
	}

	fmt.Fprintf(w, "https://www.youtube-nocookie.com/embed/"+string(video))
}
