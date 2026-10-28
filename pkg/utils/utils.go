package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
