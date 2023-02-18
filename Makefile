generate-pb:
	@echo Generating protobufs: $(shell find proto -type f -name '*.proto')
	@protoc -Iproto \
		--go_out=. --go_opt=module=github.com/MarioScripts/vlr-api \
		--go-grpc_out=. --go-grpc_opt=module=github.com/MarioScripts/vlr-api \
		--proto_path=proto $(shell find proto -type f -name '*.proto')

build:
	@echo Building Server and Client
	@go build -o bin/server ./internal/server