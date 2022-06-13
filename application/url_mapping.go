package application

import (
	"net/http"

	"github.com/idoyudha/bookstore_items_api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods((http.MethodGet))
}
