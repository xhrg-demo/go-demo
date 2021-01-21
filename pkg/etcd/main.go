package main
//go get go.etcd.io/etcd v3.3.25+incompatible
import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	//初始化客户端
	client, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 1 * time.Second,
	})

	//插入一个数据
	kv := clientv3.NewKV(client)
	kv.Put(context.Background(), "name", "zangsan")

	//读取数据
	res, _ := client.Get(context.Background(), "name")
	for _, item := range res.Kvs {
		log.Printf("%v:::%v", string(item.Key), string(item.Value))
	}
}
