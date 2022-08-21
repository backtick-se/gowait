
.PHONY: grpc
grpc:
	protoc --go_out=. --go_opt=module=cowait --go-grpc_out=. --go-grpc_opt=module=cowait protobuf/cowait.proto && \
    cd python && poetry run python3 -m grpc_tools.protoc -I../protobuf --python_out=./cowait/pb --grpc_python_out=./cowait/pb ../protobuf/cowait.proto

.PHONY: tidy
tidy:
	go mod tidy && \
	rm -rf vendor && \
	go mod vendor
