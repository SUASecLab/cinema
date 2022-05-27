package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
)

func updateVideo(videoUrl, uuid string, w http.ResponseWriter) {
	// Check if user exists
	exists, errorMsg := extensions.UserExists(adminExtensions, uuid)
	if !exists {
		log.Println(errorMsg)
		fmt.Fprintln(w, errorMsg)
		return
	}

	// Change video
	realUrl, err := url.Parse(videoUrl)
	if err != nil {
		msg := "Could not parse URL:"
		log.Println(msg, err)
		fmt.Fprintln(w, msg, err)
		return
	}

	if realUrl.Host != "www.youtube.com" ||
		realUrl.Path != "/watch" {
		fmt.Fprintln(w, "Not a YouTube video")
		return
	}

	videoId := realUrl.Query().Get("v")

	if len(videoId) != 11 {
		fmt.Fprintln(w, "Invalid video ID")
		return
	}

	videoFile, err := os.Create(videoPath)
	defer videoFile.Close()

	msg := "Could store video:"
	if err != nil {
		log.Println(msg, err)
		fmt.Fprintln(w, msg, err)
		return
	}

	_, err = videoFile.WriteString(videoId)
	if err != nil {
		log.Println(msg, err)
		fmt.Fprintln(w, msg, err)
		return
	}

	fmt.Fprintf(w, "Stored video successfully")
}
