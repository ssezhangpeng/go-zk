package operate

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"github.com/ssezhangpeng/go-zk/connect"
)

func find(path string) (result []byte, state *zk.Stat, err error) {
	conn, err := connect.NewZkConn()
	if err != nil {
		fmt.Printf("operate zk connect failed, err: %+v\n", err)
		return result, state, err
	}
	defer conn.Close()

	return conn.Get(path)
}
