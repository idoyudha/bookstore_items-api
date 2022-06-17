package items

import (
	"errors"

	"github.com/idoyudha/bookstore_items_api/clients/elasticsearch"
	"github.com/idoyudha/bookstore_utils-go/rest_errors"
)

const indexItems = "items"

func (i *Item) Save() rest_errors.RestErr {
	_, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	return nil
}
