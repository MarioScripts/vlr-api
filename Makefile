MODULE := github.com/MarioScripts/vlr-api

pb:
	@echo Generating protobufs: $(shell find proto -type f -name '*.proto')
	@protoc -Iproto \
		--go_out=. --go_opt=module=$(MODULE) \
		--go-grpc_out=. --go-grpc_opt=module=$(MODULE) \
		--proto_path=proto $(shell find proto -type f -name '*.proto')

build:
	@echo Building Server
	@go build -o bin/server ./internal/server

run: pb build
	@./bin/server