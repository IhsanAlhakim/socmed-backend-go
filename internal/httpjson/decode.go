package httpjson

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var ErrEmptyBody = errors.New("Request body cannot be empty")

func Decode(r *http.Request, dataStorage any) error {
	if err := json.NewDecoder(r.Body).Decode(dataStorage); err != nil {
		if err == io.EOF {
			return ErrEmptyBody
		}
		return err
	}
	return nil
}
