
.PHONY: grpc
grpc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative core/pb/cowait.proto && \
    cd python && poetry run python3 -m grpc_tools.protoc -I../core/pb --python_out=./cowait/pb --grpc_python_out=./cowait/pb ../core/pb/cowait.proto

.PHONY: tidy
tidy:
	go mod tidy && \
	rm -rf vendor && \
	go mod vendor
