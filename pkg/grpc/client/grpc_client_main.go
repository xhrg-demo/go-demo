package main

import (
	"context"
	"fmt"
	"go-demo/pkg/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

	//grpc的header
	md := metadata.Pairs(
		"header1", "headerValue1",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	response, err := c.Query( /*context.Background()换成带header的ctx*/ ctx, request)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("我是grpc_client，我收到了数据是：%v", response.Data)
}
