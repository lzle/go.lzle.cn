package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	pb "go.lzle.cn/etcd/pb"
	"google.golang.org/grpc"
	"time"
)



func main()  {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.188.105:2379","192.168.188.110:2379","192.168.188.37:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd err:", err)
		return
	}
	r := &etcdnaming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)

	conn, err := grpc.Dial("myService", grpc.WithBalancer(b), grpc.WithBlock(),grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	// 调用服务端的SayHello
	for {
		resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "连长"})
		if err != nil {
			fmt.Printf("could not greet: %v", err)
		}
		fmt.Printf("Greeting: %s !\n", resp.Message)
		time.Sleep(time.Millisecond*100)
	}

}
