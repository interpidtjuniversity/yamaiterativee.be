package application

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

var AliDNSClient *alidns.Client = nil
const MainSite = "manchestercity.ren"
const MainSiteHost = "47.114.153.37"
const DomainRecordType = "A"
const Schema = "https"

func init() {
	client, _ := alidns.NewClientWithAccessKey("cn-qingdao", "LTAI5tJ8otAjZGng83ksbpXX", "OfcgDQjx8RyNQSKZ9vD3jt8HzTkc9X")
	AliDNSClient = client
}

func ApplyApplicationDomain(envType, application string) error{

	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = Schema
	request.DomainName = MainSite
	request.Value = MainSiteHost
	request.Type = DomainRecordType
	request.RR = fmt.Sprintf("%s.%s", envType, application)

	_, err := AliDNSClient.AddDomainRecord(request)
	return err
}