package connect

import (
	"github.com/go-zookeeper/zk"
	"time"
)

func NewZkConn() (conn *zk.Conn, err error) {
	hosts := []string{"127.0.0.1:2181", "127.0.0.1:2182", "127.0.0.1:2183"}
	conn, _, err = zk.Connect(hosts, time.Second)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
