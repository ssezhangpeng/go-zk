package operate

import (
	"encoding/json"
	"fmt"
	"github.com/go-zookeeper/zk"
	"github/ssezhangpeng/go-zk/connect"
)

// CreatePersistNode 创建持久节点
func CreatePersistNode(path string, data interface{}) error {
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

	fmt.Printf("operate persist node(%s, %+v) success\n", path, data)

	return nil
}

// CreateEphemeralNode 创建临时节点
func CreateEphemeralNode(path string, data interface{}) error {
	conn, err := connect.NewZkConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	b, err := json.Marshal(data)
	if path, err = conn.Create(path, b, zk.FlagEphemeral, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Printf("operate path failed, err: %+v\n", err)
		return err
	}

	fmt.Printf("operate ephemeral node(%s, %+v) success\n", path, data)

	return nil
}

// CreatePersistSequenceNode 创建持久时序节点
func CreatePersistSequenceNode(path string, data interface{}) error {
	conn, err := connect.NewZkConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	b, err := json.Marshal(data)
	if path, err = conn.Create(path, b, zk.FlagSequence, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Printf("operate path failed, err: %+v\n", err)
		return err
	}

	fmt.Printf("operate persist sequence node(%s, %+v) success\n", path, data)

	return nil
}

// CreateEphemeralSequenceNode 创建临时时序节点
func CreateEphemeralSequenceNode(path string, data interface{}) error {
	conn, err := connect.NewZkConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	b, err := json.Marshal(data)
	if path, err = conn.Create(path, b, zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Printf("operate path failed, err: %+v\n", err)
		return err
	}

	fmt.Printf("operate ephemeral sequence node(%s, %+v) success\n", path, data)

	return nil
}


