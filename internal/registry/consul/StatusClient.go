package consul

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const leaderUrl = "http://%s/v1/status/leader"

type StatusClient struct {
	client *http.Client
}

func (client *StatusClient) leader(context RequestContext) ([]byte, error){
	request, err := http.NewRequest("GET", fmt.Sprintf(leaderUrl, context.Registry),nil)
	if err != nil {
		return nil, err
	}
	response, _:= client.client.Do(request)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	return data, err
}

func NewStatusClient(client *http.Client) *StatusClient {
	return &StatusClient{client: client}
}
