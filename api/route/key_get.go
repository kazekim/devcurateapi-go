/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package route

import (
	"github.com/gorilla/mux"
	"github.com/kazekim/devcurateapi-go/api/app/usecase"
	"github.com/kazekim/devcurateapi-go/api/pkg/jhhttp"
	"gopkg.in/mgo.v2"
	"net/http"
)

type GetKeyResponse struct {
	Key string
}

func GetKey(w http.ResponseWriter, r *http.Request, db *mgo.Database) {

	vars := mux.Vars(r)
	id := vars["id"]

	u, err := usecase.BuildKeyUseCase(db)
	if err != nil {
		jhhttp.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	key, err := u.GetKeyByID(id)
	if err != nil {
		jhhttp.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := GetKeyResponse{
		Key: key.Value,
	}

	jhhttp.ResponseJSON(w, response)
}

