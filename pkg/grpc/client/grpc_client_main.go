package main

import (
	"context"
	"fmt"
	"go-demo/pkg/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	//连接gRPC服务器
	conn, err := grpc.Dial("127.0.0.1:1535", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	//初始化客户端
	c := proto.NewOrderActionClient(conn)
	//调用方法
	request := new(proto.OrderRequest)
	request.OrderId = 100
	request.OrderName = "订单名称"
	response, err := c.Query(context.Background(), request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Data)
}
