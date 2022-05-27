package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
)

func showVideo(uuid string, w http.ResponseWriter) {
	// Check if user exists
	exists, errorMsg := extensions.UserExists(adminExtensions, uuid)
	if !exists {
		w.WriteHeader(403)
		log.Println(errorMsg)
		fmt.Fprintln(w, errorMsg)
		return
	}

	video, err := os.ReadFile(videoPath)
	if err != nil {
		w.WriteHeader(404)
		log.Println("No video stored:", err)
		return
	}

	fmt.Fprintf(w, "https://www.youtube.com/embed/"+string(video))
}
