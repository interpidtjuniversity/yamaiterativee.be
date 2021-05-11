package consul

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"net/http"
	"yama.io/yamaIterativeE/internal/resource"
)

var consul = &Consul{}
var req = &RequestContext{}

type Consul struct {
	client        *http.Client
	agentClient   *AgentClient
	healthClient  *HealthClient
	catalogClient *CatalogClient
	statusClient  *StatusClient
}

type RequestContext struct {
	Registry string
	Service  string
}

func InitConsul() {
	req.Registry = fmt.Sprintf("%s:8500", resource.GLOBAL_CONSUL_IP)
	consul.client = &http.Client{Transport: http.DefaultTransport}
	consul.agentClient = NewAgentClient(consul.client)
	consul.catalogClient = NewCatalogClient(consul.client)
	consul.healthClient = NewHealthClient(consul.client)
	consul.statusClient = NewStatusClient(consul.client)
}

func Ping(ctx *macaron.Context) *http.Response{
	// mock
	return consul.agentClient.ping(req)
}

func Register(ctx *macaron.Context) *http.Response{
	return consul.agentClient.register(ctx.Req.Request, req)
}

func DeRegister(ctx *macaron.Context) *http.Response{
	req.Service = ctx.Params(":service")
	return consul.agentClient.deregister(req)
}

func GetServices(ctx *macaron.Context) ([]byte, error) {
	return consul.catalogClient.getServices(req)
}

func GetServiceInstances(ctx *macaron.Context) ([]byte, error) {
	req.Service = ctx.Params(":service")
	return consul.healthClient.getServiceInstances(req)
}

func Leader(ctx *macaron.Context) ([]byte, error) {
	return consul.statusClient.leader(req)
}