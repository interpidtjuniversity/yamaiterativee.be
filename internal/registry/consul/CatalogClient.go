package consul

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const getServicesUrl = "http://%s/v1/catalog/services"

type CatalogClient struct {
	client *http.Client
}

func (client *CatalogClient) getServices(context RequestContext) ([]byte, error) {

	request, err := http.NewRequest("GET", fmt.Sprintf(getServicesUrl, context.Registry),nil)
	if err != nil {
		return nil, err
	}
	response, _:= client.client.Do(request)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	return data, err
}

func NewCatalogClient(client *http.Client) *CatalogClient {
	return &CatalogClient{client: client}
}

