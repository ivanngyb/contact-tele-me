package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ErrorObject struct {
	Message string `json:"message"`
}

func WithError(next func(w http.ResponseWriter, r *http.Request) (int, error)) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		contents, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewReader(contents))
		code, err := next(w, r)
		if err != nil {
			errObj := &ErrorObject{Message: err.Error()}
			jsonErr, wErr := json.Marshal(errObj)
			if wErr != nil {
				http.Error(w, `{"message":"JSON failed, please contact IT.","error_code":"00001"}`, code)
				return
			}
			http.Error(w, string(jsonErr), code)
			return
		}
	}
	return fn
}
