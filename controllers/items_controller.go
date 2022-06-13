package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/idoyudha/bookstore_items_api/domain/items"
	"github.com/idoyudha/bookstore_items_api/services"
	"github.com/idoyudha/bookstore_items_api/utils/http_utils"
	"github.com/idoyudha/bookstore_oauth-go/oauth"
	"github.com/idoyudha/bookstore_utils-go/rest_errors"
)

var ItemsController itemsControllerInterface = &itemsController{}

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (cont *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	// authenticating our request
	if err := oauth.AuthenticateRequest(r); err != nil {
		// return error json to user
		// http_utils.ResponseError(w, err)
		return
	}

	// if auth is valid, take body from request
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseError(w, *respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil { //using requestBody to fill itemRequest being an item
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.ResponseError(w, *respErr)
		return
	}

	itemRequest.Seller = oauth.GetClientId(r)

	result, createErr := services.ItemsServices.Create(itemRequest)
	if createErr != nil {
		http_utils.ResponseError(w, *createErr)
		return
	}
	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (cont *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
