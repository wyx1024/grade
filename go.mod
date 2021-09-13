module go-growth

go 1.15

require (
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.3.0 // indirect
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/text v0.3.7
	google.golang.org/grpc v1.40.0
)

replace google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.26.0
