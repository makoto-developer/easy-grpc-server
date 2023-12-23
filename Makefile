generate:
	protoc --go_out=pb/ --go_opt=paths=source_relative \
           --go-grpc_out=pb/ --go-grpc_opt=paths=source_relative \
           --go-grpc_opt require_unimplemented_servers=false src/proto/*.proto
server:
	go run src/app/server/server.go
client:
	go run src/app/client/client.go
