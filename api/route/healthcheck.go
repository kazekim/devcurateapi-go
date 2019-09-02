/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package route

import (
	"github.com/kazekim/devcurateapi-go/api/pkg/jhhttp"
	"gopkg.in/mgo.v2"
	"net/http"
)


type GetHealthCheckResponse struct {
	Message string
}

func HealthCheck(w http.ResponseWriter, r *http.Request, db *mgo.Database) {

	response := GetHealthCheckResponse{
		Message: "Server is still Alive!!",
	}

	jhhttp.ResponseJSON(w, response)
}
