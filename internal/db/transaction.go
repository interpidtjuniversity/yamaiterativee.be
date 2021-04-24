package db

import "xorm.io/xorm"

var Proxy TransactionProxy

type ProxySession struct {
	Session *xorm.Session
}

type TransactionProxy struct {
}

func (proxy *TransactionProxy) TransactionExecute(executable func(*ProxySession) (map[string]interface{}, error)) (map[string]interface{}, error) {
	proxySession := &ProxySession{Session: x.NewSession()}
	defer proxySession.Session.Close()
	_ = proxySession.Session.Begin()
	data, err := executable(proxySession)
	if err != nil {
		proxySession.Session.Rollback()
		return nil, err
	}
	err = proxySession.Session.Commit()
	return data, err
}
