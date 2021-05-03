package consul

import (
	"fmt"
	"net/http"
)

const registryUrl = "http://%s/v1/agent/service/register"
const deRegistryUrl = "http://%s/v1/agent/service/deregister/%s"
const pingUrl = "http://%s/v1/agent/self"

type AgentClient struct {
	client *http.Client
}

func (client *AgentClient) ping(context *RequestContext) *http.Response {
	request, err := http.NewRequest("GET", fmt.Sprintf(pingUrl, context.Registry),nil)
	if err != nil {
		return nil
	}
	response, _:= client.client.Do(request)
	return response
}

func (client *AgentClient) register(request *http.Request, context *RequestContext) *http.Response {
	request, err := http.NewRequest("PUT", fmt.Sprintf(registryUrl, context.Registry),request.Body)
	if err != nil {
		return nil
	}
	response, err := client.client.Do(request)
	if err != nil {
		return nil
	}
	return response
}

func (client *AgentClient) deregister(context *RequestContext) *http.Response {
	request, err := http.NewRequest("PUT", fmt.Sprintf(deRegistryUrl, context.Registry, context.Service),nil)
	if err != nil {
		return nil
	}
	response, _:= client.client.Do(request)
	return response
}

func NewAgentClient(client *http.Client) *AgentClient {
	return &AgentClient{client: client}
}
