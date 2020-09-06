package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"go-demo/pkg/grpc/proto"
	"google.golang.org/grpc"
	"net"
)

type OrderAction struct{}

func (h OrderAction) Query(context context.Context, request *proto.OrderRequest) (*proto.OrderResponse, error) {
	response := &proto.OrderResponse{}
	response.Code = 200
	response.Data = fmt.Sprintf("%s:%d", request.OrderName, request.OrderId)
	return response, nil
}
func (h OrderAction) Update(context context.Context, request *proto.OrderRequest) (*empty.Empty, error) {
	return nil, nil
}
func main() {
	listen, err := net.Listen("tcp", ":1535")
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}
	//实现gRPC Server
	s := grpc.NewServer()
	//注册helloServer为客户端提供服务,注意这里的OrderAction{}，等价new(OrderAction),等价&OrderAction{}
	proto.RegisterOrderActionServer(s, OrderAction{})
	s.Serve(listen)
}
