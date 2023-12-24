gen:
    protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    rpc/parser/parser.proto

build:
	go build -o bin/server cmd/server/main.go

run: build
	./bin/server

grpcurl:
    grpcurl -plaintext -d '{}' localhost:8080 JsonParsingService.ParseJsonFiles

loadEnv:
    export $(cat .env | xargs)

checkIFEnvExists:
    ifeq (,$(wildcard .env))
        $(error .env file does not exist)
    endif

dockerBuildRun: checkIFEnvExists
	docker build -t json-parser . && \
	docker run --rm -it -p 8080:8080 --env-file .env json-parser

