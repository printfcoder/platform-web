module github.com/micro-in-cn/platform-web

go 1.12

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

require (
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/hashicorp/consul v1.5.1
	github.com/lib/pq v1.1.1
	github.com/micro/cli v0.1.0
	github.com/micro/go-config v1.1.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.2.0
	github.com/micro/go-web v1.0.0
	go.uber.org/zap v1.10.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
