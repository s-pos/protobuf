generate:
		protoc --go_out=./go --go-grpc_out=./go ./pb/*/*.proto