
.PHONY: install
install:
	go mod download

.PHONY: grpc
grpc:
	protoc --go_out=. --go_opt=module=cowait --go-grpc_out=. --go-grpc_opt=module=cowait protobuf/cowait.proto && \
    cd sdk/cowait-python && poetry run python3 -m grpc_tools.protoc -I../../protobuf --python_out=./cowait/pb --grpc_python_out=./cowait/pb ../../protobuf/*.proto

.PHONY: grpc-py
grpc-py:
    cd sdk/cowait-python && poetry run python3 -m grpc_tools.protoc -I../../protobuf --python_out=./cowait/pb --grpc_python_out=./cowait/pb ../../protobuf/*.proto

.PHONY: grpc-js
grpc-js:
	cd sdk/cowait-js && npx grpc_tools_node_protoc \
		--js_out=import_style=commonjs,binary:./src/pb \
		--grpc_out=generate_package_definition:./src/pb \
		-I ../../protobuf \
		../../protobuf/*.proto && \
	npx grpc_tools_node_protoc \
		--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
		--ts_out=generate_package_definition:./src/pb \
		-I ../../protobuf \
		../../protobuf/*.proto

.PHONY: tidy
tidy:
	go mod tidy && \
	go mod download && \
	rm -rf vendor && \
	go mod vendor
