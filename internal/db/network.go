package db

import (
	"fmt"
	"math/rand"
	"time"
	"xorm.io/builder"
)

/**
Type
	A: 10.0.0.0->10.255.255.255                       10.x.y.0/24            0<=x<=255 0<=y<=255
	B: 172.16.0.0->172.31.255.255                     172.x.y.0/24
	C: 192.168.0.0->192.168.255.255                   192.168.x.0/24
**/
const (
	TYPE_A = "10.%d.%d.0/24"         // check 0<=x<=255 0<=y<=255
	TYPE_B = "172.%d.%d.0/24"        // check 16<=x<=31 0<=y<=255
	TYPE_C = "192.168.%d.0/24"       // check 0<=x<=255
)
var TYPE_MAP = map[int]string{0:"TYPE_A", 1:"TYPE_B", 2:"TYPE_C"}

type NetWork struct {
	Id              int64  `xorm:"id autoincr pk"`
	NetWorkName     string `xorm:"network_name"` //maybe an app need more NetWorkName expend it to []string
	IPRange         string `xorm:"ip_range"`     //maybe an app need more IPRange expend it to []string
	ApplicationName string `xorm:"app_name"`
	Type            string `xorm:"type"`
	OwnerName       string `xorm:"owner_name"`
}

func CreateApplicationNetWork(networkName, ownerName, appName string) (*NetWork, error){
	rand.Seed(time.Now().Unix())
	networkType := rand.Intn(3)
	network := AllocateNetWork(TYPE_MAP[networkType])
	if network == ""{
		return nil, fmt.Errorf("allocate network error, networkType:%s", TYPE_MAP[networkType])
	}
	newNetwork := &NetWork{
		NetWorkName: networkName,
		OwnerName: ownerName,
		ApplicationName: appName,
		IPRange: network,
		Type: TYPE_MAP[networkType],
	}
	_, _ = x.Table("network").Insert(newNetwork)

	return newNetwork, nil
}

func GetApplicationNetWorkByName(networkName string) (*NetWork, error) {
	network := new(NetWork)
	exist, err := x.Table("network").Where(builder.Eq{"network_name": networkName}).Get(network)
	if err!=nil || !exist {
		return nil, err
	}
	return network, nil
}

func AllocateNetWork(networkType string) string{
	// 1. get all type of network, for each group compute it
	if networkType == "TYPE_A" {
		typeA, err := GetNetWorkIPRangeByType("TYPE_A")
		if err == nil {
			i, x, y := 0, 0, 0
			for i < len(typeA) {
				if fmt.Sprintf(TYPE_A, x, y) == typeA[i].IPRange {
					if y < 255 {
						y++
					} else {
						if x == 255 {
							return ""
						}
						y = 0
						x++
					}
				} else {
					return fmt.Sprintf(TYPE_A, x, y)
				}
				i++
			}
			return fmt.Sprintf(TYPE_A, x, y)
		}
		return ""
	} else if networkType == "TYPE_B" {
		typeB, err := GetNetWorkIPRangeByType("TYPE_B")
		if err == nil {
			i, x, y := 0, 16, 0
			for i < len(typeB) {
				if fmt.Sprintf(TYPE_B, x, y) == typeB[i].IPRange {
					if y < 255 {
						y++
					} else {
						if x == 31 {
							return ""
						}
						y = 0
						x++
					}
				} else {
					return fmt.Sprintf(TYPE_B, x, y)
				}
				i++
			}
			return fmt.Sprintf(TYPE_B, x, y)
		}
		return ""
	} else if networkType == "TYPE_C" {
		typeC, err := GetNetWorkIPRangeByType("TYPE_C")
		if err == nil {
			i, x := 0, 0
			for i < len(typeC) {
				if fmt.Sprintf(TYPE_C, x) == typeC[i].IPRange {
					if x < 255 {
						x++
					} else {
						return ""
					}
				} else {
					return fmt.Sprintf(TYPE_C, x)
				}
				i++
			}
			return fmt.Sprintf(TYPE_C, x)
		}
		return ""
	}
	return ""
}

func GetNetWorkIPRangeByType(networkType string) ([]*NetWork, error){
	var networks []*NetWork
	rows, err := x.Table("network").Cols("owner_name", "app_name","network_name","ip_range").Where(builder.Eq{"type":networkType}).Rows(new(NetWork))
	if err != nil {
		return networks, err
	}
	defer rows.Close()
	for rows.Next() {
		var network NetWork
		rows.Scan(&network)
		networks = append(networks, &network)
	}

	return networks, nil
}