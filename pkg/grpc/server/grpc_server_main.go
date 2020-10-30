package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"go-demo/pkg/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	/*
			这一个方法reflection.Register会启动grpc反射服务，如果不需要刻意不开启，比如线上环境。
			grpcurl -plaintext localhost:1535 list就可以看到服务列表了
		>安装grpcurl
		>go get github.com/fullstorydev/grpcurl
		>go install github.com/fullstorydev/grpcurl/cmd/grpcurl
	*/
	reflection.Register(s)

	s.Serve(listen)

}
