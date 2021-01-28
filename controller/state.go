package controller

import (
	"encoding/json"
	"log"

	"github.com/lucas-engen/WarTelemetry/model/state"
	network "github.com/lucas-engen/WarTelemetry/network/http"
	"github.com/lucas-engen/WarTelemetry/utils"
)

// GetStateData function retrieves data about running missions
func GetStateData() (st *state.AircraftState) {
	// Sends a HTTP request
	data := network.GetDataFromURL(state.GetURL())

	// Process JSON into a readable format
	data = utils.ProcessJSON(data)

	// Decode JSON into a struct
	marshalErr := json.Unmarshal(data, &st)
	if marshalErr != nil {
		log.Fatalf("[FATAL] Failed to decode JSON data: %v", marshalErr)
	}

	return
}