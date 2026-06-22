package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	args := map[string]any{
		"steam_id": r.URL.Query().Get("steam_id"),
		"precheck": r.URL.Query().Has("precheck"),
	}
	v := Main(args)
	b, _ := json.MarshalIndent(v, "", "\t")
	os.Stdout.Write(b)
}
