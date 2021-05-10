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
	// "|LTAI5|tJ8|otA|jZGng|83ksbpXX|
	// "|OfcgD|Qjx8|RyNQSKZ9v|D3jt8H|zTkc9X|"
	client, _ := alidns.NewClientWithAccessKey("cn-qingdao", "xxx", "xxx")
	AliDNSClient = client
}

func ApplyApplicationDomain(envType, domainName string) error{

	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = Schema
	request.DomainName = MainSite
	request.Value = MainSiteHost
	request.Type = DomainRecordType
	if envType!= "" {
		request.RR = fmt.Sprintf("%s.%s", envType, domainName)
	}else {
		request.RR = fmt.Sprintf(domainName)
	}

	_, err := AliDNSClient.AddDomainRecord(request)
	return err
}