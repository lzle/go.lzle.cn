package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)


func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.188.105:2379","192.168.188.110:2379","192.168.188.37:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}

	defer cli.Close()

	resp,err := cli.Grant(context.TODO(),10)
	if err != nil {
		log.Println(err)
	}

	_,err = cli.Put(context.TODO(),"lzl","dsg",clientv3.WithLease(resp.ID))
	if err != nil {
		log.Println(err)
	}

	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		fmt.Println(time.Now())
		ka,ok := <-ch
		fmt.Println(ok)
		fmt.Println("ttl:", ka.TTL)
	}

}
