package mq

import (
	"github.com/cloudwego/biz-demo/gomall/app/email/conf"
	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(conf.GetConf().Nats.Address)
	if err != nil {
		panic(err)
	}
}
