module github.com/micrease/meshop/micrease-gateway

go 1.18

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	github.com/micrease/meshop/micrease-protos v0.0.0-incompatible
	google.golang.org/grpc v1.47.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220608133413-ed9918b62aac // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/micrease/meshop/micrease-protos => ../micrease-protos
