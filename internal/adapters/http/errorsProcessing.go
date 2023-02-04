package http

import (
	"encoding/json"
)

func makeJsonFromError(err error) []byte {

	answer := Answer{}
	answer.Success = false
	answer.ErrorText = err.Error()
	jsonResp, _ := json.Marshal(answer)

	return jsonResp
}
