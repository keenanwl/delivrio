package httputils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Map = map[string]interface{}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UnmarshalRequestBody(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	log.Println("Body: ", string(body))

	if err := json.Unmarshal(body, &v); err != nil {
		return err
	}

	return nil
}
