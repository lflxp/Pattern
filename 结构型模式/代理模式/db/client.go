package db

import (
	"fmt"
	"net/rpc"
)

// 到目前为止，我们已经为数据库提供了对外访问的方式。现在，我们需要一个远程代理来连接数据库服务端，并进行相关的数据库操作。对消息处理系统而言，它不需要，也不应该知道远程代理与数据库服务端交互的底层细节，这样可以减轻系统之间的耦合。因此，远程代理需要实现db.KvDb：

// 数据库服务端远程代理，实现db.KvDb接口
type Client struct {
	// RPC客户端
	cli *rpc.Client
}

func (c *Client) Save(record Record, reply *bool) error {
	var ret bool
	// 通过RPC调用服务端的接口
	err := c.cli.Call("Server.Save", record, &ret)
	if err != nil {
		fmt.Printf("Call db Server.Save rpc failed,error: %v", err)
		*reply = false
		return err
	}
	*reply = ret
	return nil
}

func (c *Client) Get(key string, reply *string) error {
	var ret string
	// 通过RPC调用服务端的接口
	err := c.cli.Call("Server.Get", key, &ret)
	if err != nil {
		fmt.Printf("Call db Server.Get rpc failed, error: %v", err)
		*reply = ""
		return err
	}
	*reply = ret
	return nil
}

// 工厂方法，返回远程代理实例
func CreateClient() *Client {
	rpcCli, err := rpc.Dial("tcp", "127.0.0.1:5678")
	if err != nil {
		fmt.Printf("Create rpc client failed, error: %v", err)
		return nil
	}
	return &Client{cli: rpcCli}
}
