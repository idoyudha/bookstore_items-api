package application

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/idoyudha/bookstore_items_api/clients/elasticsearch"
)

var router = mux.NewRouter()

func StartApplication() {
	elasticsearch.GetEsClient()

	mapUrls()

	srv := &http.Server{
		Addr:         "localhost:8081",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
