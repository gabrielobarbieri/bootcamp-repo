package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrRequestJSONInvalid = errors.New("request json invalid")
)

// Decode request JSON into a target
func JSON(r *http.Request, ptr interface{}) error {
	err := json.NewDecoder(r.Body).Decode(ptr)
	if err != nil {
		return fmt.Errorf("%w. %v", ErrRequestJSONInvalid, err)
	}

	return nil
}
