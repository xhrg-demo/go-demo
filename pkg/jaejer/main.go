package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"os"
	"time"
)

//main方法执行代码
func main() {

	tracer, closer := initJaeger()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("span_demo")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	foo3("Hello foo3", ctx)
	span.Finish()
}

//配置jaeger实例
func initJaeger() (opentracing.Tracer, io.Closer) {
	//cfg := &config.Configuration{
	//	Sampler: &config.SamplerConfig{
	//		Type:  "const",
	//		Param: 1,
	//	},
	//	Reporter: &config.ReporterConfig{
	//		LogSpans:           true,
	//		LocalAgentHostPort: "127.0.0.2:6831",
	//	},
	//}
	//tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	//if err != nil {
	//	panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	//}
	//return tracer, closer

	os.Setenv("JAEGER_SERVICE_NAME", "demo-jr-service")
	os.Setenv("JAEGER_AGENT_HOST", "127.0.0.1")
	os.Setenv("JAEGER_AGENT_PORT", "6831")
	os.Setenv("JAEGER_SAMPLER_TYPE", "const")
	os.Setenv("JAEGER_SAMPLER_PARAM", "1")
	cfg, _ := config.FromEnv()
	tracer, i, err := cfg.NewTracer()
	if err != nil {
		logrus.Error(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, i
}

//调用jaeger过程
func foo3(req string, ctx context.Context) (reply string) {
	//1.创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "span_demo_2")
	defer func() {
		//4.接口调用完，在tag中设置request和reply
		span.SetTag("request", req)
		span.SetTag("reply", reply)
		span.Finish()
	}()
	println(req)
	//2.模拟处理耗时
	time.Sleep(time.Second / 2)
	//3.返回reply
	reply = "foo3Reply"
	return
}
