MODULE := github.com/MarioScripts/vlr-api
TAG := latest

pb:
	@echo Generating protobufs: $(shell find proto -type f -name '*.proto')
	@protoc -Iproto \
		--go_out=. --go_opt=module=$(MODULE) \
		--go-grpc_out=. --go-grpc_opt=module=$(MODULE) \
		--proto_path=proto $(shell find proto -type f -name '*.proto')

build:
	@echo Building Server
	@go build -o bin/server ./internal/server

build-docker: get-ecr
	@echo Building Docker Image
	@docker build -t $(ECR_REPO):$(TAG) .

publish-docker: build-docker
	@echo Publishing Docker Image $(ECR_REPO):$(TAG)
	@docker push $(ECR_REPO):$(TAG)

publish-go: 
	@echo Creating Git tag $(TAG)
	@git tag $(TAG)
	@git push origin $(TAG)
	@git checkout $(TAG)
	@echo Publishing Go module $(MODULE)@$(TAG)
	@GOPRIVATE=$(MODULE) 
	@GOPROXY=proxy.golang.org 
	@go list -m $(MODULE)@$(TAG)

run: pb build
	@./bin/server

get-ecr:
	$(eval ECR_REPO := $(shell aws ecr describe-repositories --query "repositories[?repositoryName=='vlr-api'].repositoryUri" --output text))
	@echo Found repo $(ECR_REPO)
