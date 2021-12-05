package operate

import (
	"fmt"
	"github/ssezhangpeng/go-zk/connect"
)

func delete(path string) error {
	conn, err := connect.NewZkConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	isExist, stat, err := conn.Exists(path)
	if err != nil {
		fmt.Printf("isExist failed, err: %+v\n", err)
		return err
	}

	fmt.Printf("path(%s) is exist(%t).\n", path, isExist)

	if err = conn.Delete(path, stat.Version); err != nil {
		fmt.Printf("delete path(%s) failed, err: %+v\n", path, err)
		return err
	}

	fmt.Printf("delete path(%s) success.\n", path)

	return nil
}
