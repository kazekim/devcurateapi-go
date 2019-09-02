/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
	"github.com/kazekim/devcurateapi-go/api/app/usecase"
	"github.com/kazekim/devcurateapi-go/api/config"
	"github.com/kazekim/devcurateapi-go/api/framework/mongofw"
	"github.com/kazekim/devcurateapi-go/api/route"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

var db *mgo.Database

func main() {
	// Connect to mongo

	mongo, err := mongofw.Open()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()

	db = mongo.DB(config.MONGO_DB_NAME)

	// Set up routes
	r := mux.NewRouter()
	r.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		route.HealthCheck(w, r, db)
	}).
		Methods("GET")
	r.HandleFunc("/key/{id}", func(w http.ResponseWriter, r *http.Request) {
		route.GetKey(w, r, db)
	}).
		Methods("GET")
	r.HandleFunc("/keys", func(w http.ResponseWriter, r *http.Request) {
		route.PostKeys(w, r, db)
	}).
		Methods("POST")

	go func() {
		gocron.Every(30).Minutes().Do(clearKey)
		<- gocron.Start()
	}()


	_ = http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
	log.Println("Listening on port 8080...")
}

func clearKey() {

	u, err := usecase.BuildKeyUseCase(db)
	if err != nil {
		return
	}

	_ = u.RemoveKeyCreatedMoreThanOnHour()

}