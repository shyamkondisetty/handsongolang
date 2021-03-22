package utils

import (
	"encoding/json"
	"net/http"
)

func ParseRequest(req *http.Request, data interface{}) (err error) {

	decoder := json.NewDecoder(req.Body)

	err = decoder.Decode(&data)
	if err != nil {
		return err
	}
	return
}
