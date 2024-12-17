package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// what is interface and how it works??
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}