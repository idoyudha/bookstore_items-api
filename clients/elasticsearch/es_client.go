package elasticsearch

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/idoyudha/bookstore_utils-go/logger"
)

var Client esClientInterface = &esClient{}

type esClientInterface interface {
	setClient(*elasticsearch.Client)
	Index(string, interface{}) (*esapi.Response, error)
	Get(string, string) (*esapi.Response, error)
}

type esClient struct {
	client *elasticsearch.Client
}

func GetEsClient() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9201",
		},
		Username: "foo",
		Password: "bar",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) setClient(client *elasticsearch.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*esapi.Response, error) {
	// Build the request body.
	data, errMarshall := json.Marshal(doc)
	if errMarshall != nil {
		log.Fatalf("Error marshaling document: %s", errMarshall)
	}

	// Set up the request object.
	req := esapi.IndexRequest{
		Index: index,
		// DocumentID: strconv.Itoa(id + 1),
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), c.client)
	if err != nil {
		logger.Error(
			fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	return res, nil
}

func (c *esClient) Get(index string, id string) (*esapi.Response, error) {
	result, err := c.client.Get(index, id)
	if err != nil {
		log.Println("Error get id")
		return nil, err
	}
	return (*esapi.Response)(result), err
}
