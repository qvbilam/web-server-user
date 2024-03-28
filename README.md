## 依赖包
```shell
# gin
go get -u github.com/gin-gonic/gin

# 日志
go get -u go.uber.org/zap

# 配置
go get -u github.com/spf13/viper

# grpc
go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc
go get -u google.golang.org/genproto

# grpc 超时重试
go get -u github.com/grpc-ecosystem/go-grpc-middleware

# validate 表单验证
go get -u github.com/gin-gonic/gin/binding

# jwt
go get -u github.com/dgrijalva/jwt-go

# client ip
go get -u github.com/thinkeridea/go-extend/exnet

# user-agent
go get -u github.com/mssola/useragent

# jaeger 链路追踪
go get github.com/uber/jaeger-client-go
go get -u github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc
```