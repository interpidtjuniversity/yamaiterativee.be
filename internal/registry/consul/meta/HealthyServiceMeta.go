package meta

type HealthyServiceMeta struct {
	Node    Node    `json:"Node"`
	Service Service `json:"Service"`
	Checks  []Check `json:"Checks"`
}

type Node struct {
	ID              string        `json:"ID"`
	Node            string        `json:"Node"`
	Address         string        `json:"Address"`
	Datacenter      string        `json:"Datacenter"`
	TaggedAddresses LanWanAddress `json:"TaggedAddresses"`
	Meta            NodeMeta      `json:"Meta"`
	CreateIndex     int           `json:"CreateIndex"`
	ModifyIndex     int           `json:"ModifyIndex"`
}
type LanWanAddress struct {
	Lan string `json:"lan"`
	Wan string `json:"wan"`
}
type NodeMeta struct {
	ConsulNetworkSegment string `json:"consul-network-segment"`
}


type Service struct {
	ID                string         `json:"ID"`
	Service           string         `json:"Service"`
	Tags              []string       `json:"Tags"`
	Address           string         `json:"Address"`
	Meta              ServiceMeta    `json:"Meta"`
	Port              int            `json:"Port"`
	Weights           ServiceWeights `json:"Weights"`
	EnableTagOverride bool           `json:"EnableTagOverride"`
	ProxyDestination  string         `json:"ProxyDestination"`
	Proxy             struct{}       `json:"Proxy"`
	Connect           struct{}       `json:"Connect"`
	CreateIndex       int            `json:"CreateIndex"`
	ModifyIndex       int            `json:"ModifyIndex"`
}

type ServiceMeta struct {
	GRPCPort string `json:"gRPC_port"`
	Secure   string `json:"secure"`
}

type ServiceWeights struct {
	Passing int `json:"Passing"`
	Warning int `json:"Warning"`
}

type Check struct {
	Node        string   `json:"Node"`
	CheckID     string   `json:"CheckID"`
	Name        string   `json:"Name"`
	Status      string   `json:"Status"`
	Notes       string   `json:"Notes"`
	Output      string   `json:"Output"`
	ServiceID   string   `json:"ServiceID"`
	ServiceName string   `json:"ServiceName"`
	ServiceTags []string `json:"ServiceTags"`
	Definition  struct{} `json:"Definition"`
	CreateIndex int      `json:"CreateIndex"`
	ModifyIndex int      `json:"ModifyIndex"`
}


