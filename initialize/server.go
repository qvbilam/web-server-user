package initialize

import (
	"fmt"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
	publicProto "user/api/qvbilam/public/v1"
	proto "user/api/qvbilam/user/v1"
	"user/global"
)

type dialConfig struct {
	host string
	port int64
}

type serverClientConfig struct {
	userDialConfig   *dialConfig
	publicDialConfig *dialConfig
}

func InitServer() {

	s := serverClientConfig{
		userDialConfig: &dialConfig{
			host: global.ServerConfig.UserServerConfig.Host,
			port: global.ServerConfig.UserServerConfig.Port,
		},
		publicDialConfig: &dialConfig{
			host: global.ServerConfig.PublicServerConfig.Host,
			port: global.ServerConfig.PublicServerConfig.Port,
		},
	}

	s.initUserServer()
	s.initPublicServer()
}

func clientOption() []retry.CallOption {
	opts := []retry.CallOption{
		retry.WithBackoff(retry.BackoffLinear(100 * time.Millisecond)), // 重试间隔
		retry.WithMax(3), // 最大重试次数
		retry.WithPerRetryTimeout(1 * time.Second),                                 // 请求超时时间
		retry.WithCodes(codes.NotFound, codes.DeadlineExceeded, codes.Unavailable), // 指定返回码重试
	}
	return opts
}

func (s *serverClientConfig) initUserServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.userDialConfig.host, s.userDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)),
		// 链路追踪
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", global.ServerConfig.UserServerConfig.Name, err)
	}

	userClient := proto.NewUserClient(conn)
	accountClient := proto.NewAccountClient(conn)

	global.UserServerClient = userClient
	global.AccountServerClient = accountClient
	fmt.Printf("grpc server: %s[%s:%d]\n", global.ServerConfig.UserServerConfig.Name, global.ServerConfig.UserServerConfig.Host, global.ServerConfig.UserServerConfig.Port)
}

func (s *serverClientConfig) initPublicServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.publicDialConfig.host, s.publicDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)),
		// 链路追踪
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", global.ServerConfig.PublicServerConfig.Name, err)
	}

	smsClient := publicProto.NewSmsClient(conn)
	global.PublicSmsServerClient = smsClient
	fmt.Printf("grpc server: %s[%s:%d]\n", global.ServerConfig.PublicServerConfig.Name, global.ServerConfig.PublicServerConfig.Host, global.ServerConfig.PublicServerConfig.Port)
}
