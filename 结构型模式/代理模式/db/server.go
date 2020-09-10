// 数据库是一个Key-Value数据库，使用map存储数据，下面为数据库的服务端实现，db.Server实现了db.KvDb接口
package db

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"time"
)

// 数据库服务端实现
type Server struct {
	// 采用map存储key-value数据
	data map[string]string
}

func (s *Server) Save(record Record, reply *bool) error {
	if s.data == nil {
		s.data = make(map[string]string)
	}

	s.data[record.Key] = record.Value
	*reply = true
	return nil
}

func (s *Server) Get(key string, reply *string) error {
	val, ok := s.data[key]
	if !ok {
		*reply = ""
		return errors.New("db has no key " + key)
	}
	*reply = val
	return nil
}

// 消息处理系统和数据库并不在同一台机器上，因此消息处理系统不能直接调用db.Server的方法进行数据存储，像这种服务提供者和服务使用者不在同一机器上的场景，使用远程代理再适合不过了。

// 远程代理中，最常见的一种实现是远程过程调用（Remote Procedure Call，简称 RPC），它允许客户端应用可以像调用本地对象一样直接调用另一台不同的机器上服务端应用的方法。在Go语言领域，除了大名鼎鼎的gRPC，Go标准库net/rpc包里也提供了RPC的实现。下面，我们通过net/rpc对外提供数据库服务端的能力：

// 启动数据库，对外提供RPC接口进行数据库的访问
func Start() {
	rpcServer := rpc.NewServer()
	server := &Server{data: make(map[string]string)}
	// 将数据库接口注册到RPC服务器上
	if err := rpcServer.Register(server); err != nil {
		fmt.Printf("Register Server to rpc failed,error: %v", err)
		return
	}

	l, err := net.Listen("tcp", "127.0.0.1:5678")
	if err != nil {
		fmt.Printf("Listen tcp failed,error: %v", err)
		return
	}
	go rpcServer.Accept(l)
	time.Sleep(1 * time.Second)
	fmt.Println("Rpc server start success")
}
