package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func logIn(w http.ResponseWriter, r *http.Request) {
	resp := authData{}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write(makeJsonFromError(err))
	} else {
		err = json.Unmarshal(b, &resp)
		if err != nil {
			w.Write(makeJsonFromError(err))
		}
	}

}
