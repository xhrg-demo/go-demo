## go-grpc demo

#### 第一步：搭建go项目，初始化go-mod，配置ide(一般用goland)，全部略略略

#### 第二步：安装protobuf
1. brew install protobuf,这一步我安装的时候遇到一个问题，然后执行"sudo chown -R 我的登陆名 /usr/local/share/man/man8"后好了。
2. protoc --version  输出libprotoc 3.13.0代表安装成功。

#### 第三步：安装proto和protoc-gen-go

注意go get后可能需要重启ide

* go get -u -v github.com/golang/protobuf/proto
* go get -u -v github.com/golang/protobuf/protoc-gen-go

#### 第四步：编辑protobuf文件

#### 编译

protoc --go_out=plugins=grpc:. ./order_service.proto

#### 编写server代码

```golang
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

```


#### 编写client代码
```golang
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

```
