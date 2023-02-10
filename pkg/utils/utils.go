package utils

import (
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, model interface{}) {
	body, err := ioutil.ReadAll(request.Body)

	if err == nil {

	}
}
