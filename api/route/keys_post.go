/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package route

import (
	"github.com/kazekim/devcurateapi-go/api/app/usecase"
	"github.com/kazekim/devcurateapi-go/api/pkg/jhhttp"
	"gopkg.in/mgo.v2"
	"net/http"
)


type PostKeysResponse struct {
	ID string
	Key string
}

func PostKeys(w http.ResponseWriter, r *http.Request, db *mgo.Database) {

	keyValue := r.FormValue("key")

	u, err := usecase.BuildKeyUseCase(db)
	if err != nil {
		jhhttp.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	key, err := u.CreateKey(keyValue)
	if err != nil {
		jhhttp.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := PostKeysResponse{
		ID: key.ID,
		Key: key.Value,
	}
	jhhttp.ResponseJSON(w, response)
}
