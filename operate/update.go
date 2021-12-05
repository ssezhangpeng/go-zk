package operate

import (
	"encoding/json"
	"fmt"
	"github.com/go-zookeeper/zk"
	"github.com/ssezhangpeng/go-zk/connect"
)

func update(path string, data interface{}) error {
	conn, err := connect.NewZkConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	b, err := json.Marshal(data)
	if path, err = conn.Create(path, b, 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Printf("operate path failed, err: %+v\n", err)
		return err
	}

	_, stat, err := conn.Get(path)
	if err != nil {
		fmt.Printf("get path(%s) failed, err: %+v\n", path, err)
		return err
	}

	if stat, err = conn.Set(path, b, stat.Version); err != nil {
		fmt.Printf("set path(%s) value(%+v) failed, err: %+v\n", path, data, err)
		return err
	}

	fmt.Printf("set path(%s) new value(%+v) success.\n", path, data)

	return nil
}
