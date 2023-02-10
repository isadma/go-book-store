package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, model interface{}) {
	body, err := ioutil.ReadAll(request.Body)

	if err == nil {
		err := json.Unmarshal([]byte(body), model)
		if err != nil {
			return
		}
	}
}
