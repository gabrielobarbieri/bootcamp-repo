package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, body interface{}) {
	if body == nil {
		w.WriteHeader(code)
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)

	w.Write(bytes)
}
