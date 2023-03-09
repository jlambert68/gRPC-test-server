
# cat -e -t -v Makefile

RunGrpcGui:
	cd ~/egen_kod/go/go_workspace/src/jlambert/grpcui/standalone && grpcui -plaintext localhost:6670


compileProto_go:
	@echo "Compile ExecutionServer proto file..."

 # generate the messages
	protoc --go_out=simpleGrpcServer simpleGrpcServer/*.proto

# generate the services
	protoc --go-grpc_out=simpleGrpcServer simpleGrpcServer/*.proto
