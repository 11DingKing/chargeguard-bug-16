package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"errors"
	"net/http"
)

var stationBatch charging.StationBatch

func ResetTaskHTTPState() { stationBatch = charging.StationBatch{} }
func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	err := stationBatch.Register([]string{"station-a", "station-b", "invalid", "station-c"})
	stored, active := stationBatch.Stats()
	if errors.Is(err, charging.ErrInvalidStation) {
		w.WriteHeader(http.StatusBadRequest)
	}
	_ = json.NewEncoder(w).Encode(map[string]int{"stored": stored, "active": active})
}
