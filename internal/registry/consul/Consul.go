package consul

import (
	"encoding/json"
	"fmt"
	"gopkg.in/macaron.v1"
	"net"
	"net/http"
	"strings"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/registry/consul/meta"
	"yama.io/yamaIterativeE/internal/resource"
)

const (
	DEFAULT_GROUP = "NONE_GROUP"
	DEFAULT_TYPE = "NONE_TYPE"

)

var consul = &Consul{}
var req = RequestContext{}
var EnvInvokeMap = map[string]string{
	DEFAULT_TYPE: "dev",
	"dev": "stable",
	"stable": "stable",
	"test": "test",
	"pre": "pre",
	"prod": "prod",
}

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
	Filter   string
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
	serverServiceName := strings.ReplaceAll(req.Service, "-", ".")
	requestServerIP, _ := getIPAddress(ctx.Req)

	requestServerType, requestServerGroupId, _ := db.GetServerTypeAndGroupByIP(requestServerIP)
	// requestServerType, requestServerGroupId := "dev", int64(17)
	if requestServerType == "" {
		requestServerGroupId = DEFAULT_GROUP
		requestServerType = DEFAULT_TYPE
	}
	// 1. if group exist
	groupServers, _ := db.GetSameGroupServerByGroupIdAndServiceName(requestServerGroupId, req.Service)
	// 2. traverse groupServer and find out target service
	var groupServiceServerIPs []string
	for _, groupServer := range groupServers {
		if strings.Contains(groupServer.Name, serverServiceName){
			groupServiceServerIPs = append(groupServiceServerIPs, groupServer.IP)
		}
	}
	var reqFilters []string
	if len(groupServiceServerIPs) == 0 {
		//3. default env map servers
		reqFilters = append(reqFilters, fmt.Sprintf("%s%sin%sService.Tags", EnvInvokeMap[requestServerType], "%20", "%20"))
	} else {
		for _, targetIP := range groupServiceServerIPs {
			reqFilters = append(reqFilters, fmt.Sprintf("Service.Address==\"%s\"", targetIP))
		}
	}
	filterHealthyServers := make([]meta.HealthyServiceMeta,0)
	for _, filter := range reqFilters {
		req.Filter = filter
		data, _ := consul.healthClient.getServiceInstances(req)
		healthyServers := make([]meta.HealthyServiceMeta, 1)
		err := json.Unmarshal(data, &healthyServers)
		if err != nil {
			return nil, err
		}
		filterHealthyServers = append(filterHealthyServers, healthyServers...)
	}

	fiterData, _ := json.Marshal(filterHealthyServers)

	return fiterData, nil
}

func Leader(ctx *macaron.Context) ([]byte, error) {
	return consul.statusClient.leader(req)
}

func getIPAddress(req macaron.Request) (string, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return "", err
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return "", err
	}
	return userIP.String(), nil
}