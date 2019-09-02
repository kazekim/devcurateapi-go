/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package jhhttp

import (
	"encoding/json"
	"net/http"
)

func ResponseError(w http.ResponseWriter, message string, code int) {

	response := Response{
		Code: code,
		Message: message,
		Data: nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(response)
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {

	response := Response{
		Code: 200,
		Message: " Success",
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}