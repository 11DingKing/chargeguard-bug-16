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
		if stored != 0 || active != 0 {
			http.Error(w, "batch rollback incomplete", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]int{"stored": stored, "active": active})
		return
	}
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
