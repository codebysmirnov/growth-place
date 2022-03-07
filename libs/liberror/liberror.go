package liberror

import (
	"encoding/json"
	"net/http"
)

// Error presents error structure
type Error struct {
	Message  string `json:"message"`
	Code     string `json:"code"`
	HTTPCode int    `json:"-"`
}

// String returns message error string
func (e Error) Error() string {
	return e.Message
}

// JSONError make error on json format
func JSONError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	e, ok := err.(*Error)
	if ok {
		w.WriteHeader(e.HTTPCode)
	}
	json.NewEncoder(w).Encode(err)
}
