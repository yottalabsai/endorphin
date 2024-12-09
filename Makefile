
.PHONY: generate-proto
generate-proto:
	cd proto && buf generate
	make generate-synapse-grpc-mock

.PHONY: generate-synapse-grpc-mock
generate-synapse-grpc-mock:
	mockgen -source=pkg/services/synapse/synapse_grpc.pb.go > pkg/services/synapse/mock_synapse/mock_synapse.go


.PHONY: install-tools
install-tools:
	rm -f ${GOPATH}/bin/protoc-gen-go
	rm -f ${GOPATH}/bin/protoc-gen-go-grpc
	rm -f ${GOPATH}/bin/buf
	rm -f ${GOPATH}/bin/mockgen
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/bufbuild/buf/cmd/buf@v1.6.0
	go install github.com/golang/mock/mockgen@v1.6.0