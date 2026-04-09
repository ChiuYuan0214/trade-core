module local.exchange-demo/replay-tool-go

go 1.26.1

require local.exchange-demo/exchange-core-go v0.0.0

require (
	github.com/ChiuYuan0214/depin v0.0.0-20260331073741-bbb9f9ba1a61 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/grpc v1.76.0 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

replace local.exchange-demo/exchange-core-go => ../../modules/exchange-core-go
