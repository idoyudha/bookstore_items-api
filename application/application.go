package application

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func StartApplication() {
	mapUrls()
	srv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
