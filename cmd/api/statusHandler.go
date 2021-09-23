package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}
	app.logger.Printf("Response sent with: ", js)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	wrote, err := w.Write(js)
	if err != nil {
		app.logger.Println(err)
	}
	app.logger.Printf("JSON wrote with: ", wrote)
}
