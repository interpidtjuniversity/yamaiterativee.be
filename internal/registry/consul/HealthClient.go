package consul

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const getServiceInstancesUrl = "http://%s/v1/health/service/%s?filter=%s"

type HealthClient struct {
	client *http.Client
}

func (client *HealthClient) getServiceInstances(context RequestContext) ([]byte, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(getServiceInstancesUrl, context.Registry, context.Service, context.Filter),nil)
	if err != nil {
		return nil, err
	}
	response, _:= client.client.Do(request)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	return data, err
}

func NewHealthClient(client *http.Client) *HealthClient {
	return &HealthClient{client: client}
}