build: gomod proto compile

gomod:
	go mod tidy

compile:
	go build -o bin/grpc-helloworld main.go 

gp:
	protoc -I ./proto --go_out ./proto --go_opt paths=source_relative --go-grpc_out ./proto --go-grpc_opt paths=source_relative ./proto/helloworld/hello_world.proto

image:
	docker build . -t grpc-helloworld:latest