.DEFAULT: all

all: download-go-to-protobuf generate

download-go-to-protobuf:
	go get k8s.io/code-generator/cmd/go-to-protobuf github.com/gogo/protobuf/protoc-gen-gogo golang.org/x/tools/cmd/cover golang.org/x/tools/cmd/goimports
# Generate code
generate:
	GO111MODULE=on go mod vendor; chmod +x hack/*.sh; hack/verify-codegen.sh; hack/generate-proto.sh

