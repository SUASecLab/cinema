package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
)

func showVideo(userToken string, w http.ResponseWriter) {
	// Check if user is allowed to see the video
	decision, err := extensions.GetAuthDecision("http://" + sidecarUrl +
		"/auth?token=" + userToken + "&service=showVideo")
	if !decision.Allowed {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err != nil {
		errorMsg := "Could not get authentication decision"
		log.Println(errorMsg, err)
		fmt.Fprintln(w, errorMsg)
		return
	}

	fmt.Fprintf(w, "https://www.youtube-nocookie.com/embed/"+video)
}
