package application

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"yama.io/yamaIterativeE/internal/resource"
)

var AliDNSClient *alidns.Client = nil
const MainSite = "manchestercity.ren"
const MainSiteHost = "47.114.153.37"
const DomainRecordType = "A"
const Schema = "https"

func InitAliYunDNS() {
	client, _ := alidns.NewClientWithAccessKey("cn-qingdao", resource.GLOBAL_ALIYUN_ACCESSKEY, resource.GLOBAL_ALIYUN_ACCESSKEY_SECRET)
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