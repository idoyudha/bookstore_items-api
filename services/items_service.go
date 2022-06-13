package services

import (
	"net/http"

	"github.com/idoyudha/bookstore_items_api/domain/items"
	"github.com/idoyudha/bookstore_utils-go/rest_errors"
)

var ItemsServices itemsServiceInterface = &itemsService{}

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(items.Item) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, &rest_errors.RestErr{
		Status:  http.StatusNotImplemented,
		Message: "implement me!",
		Error:   "not_implemented",
	}
}

func (s *itemsService) Get(items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, &rest_errors.RestErr{
		Status:  http.StatusNotImplemented,
		Message: "implement me!",
		Error:   "not_implemented",
	}
}
