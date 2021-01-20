package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/reflection"
	pb "go.lzle.cn/etcd/pb"
	"net"
	"time"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("request 8002")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}


func main() {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.188.105:2379","192.168.188.110:2379","192.168.188.37:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd err:", err)
		return
	}

	// 创建命名解析
	r := etcdnaming.GRPCResolver{Client: cli}

	// 将本服务注册添加etcd中，服务名称为myService，服务地址为本机8001端口
	r.Update(context.TODO(),"myService",naming.Update{naming.Add,":8002",""})

	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	fmt.Println("server start")
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}