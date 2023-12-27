# File monitoring gRPC service
This gRPC service is used and called by the [RESTful Gateway project](https://github.com/Rashad-j/gateway-grpc).

A gRPC microservice that monitors a given folder with files (e.g. json files), uses in memory cache to update any file changed. It has one method `ParseJsonFiles` which when called will return all files in the cache.

## Design pattern
The decorator design pattern is used for the file service, to extend the package without changing existing code. This pattern was used achieve the second principle of SOLID: open/close principle. Other principles were also used from SOLID overall: single responsibility and interfaces. 

## In memory cache
Thread safe with auto changes detection mechanism. Incoming requests will be served the cache, which is faster and more efficient than reading files on every request.

## Unit tests
The are some unit test examples for the cache and file services.

## Docker
This service is created as docker image and pushed to docker hub in order to use it later in the gateway network and container.

## How to test?
The default config has a port of `8081`. If you want to change simply create a `.env` file and then run `make dockerBuildRun`. Note that in either way, the `.env` file is required for docker to run. 

### Test using grpcurl
To test the service after running you can call `make grpcurl` or:
```
$ grpcurl -plaintext -d '{}' localhost:8080 JsonParsingService.ParseJsonFiles
```
Have a look on `Makefile` for more examples. 

### Test using the client
Alternatively, you can run the client in `cmd/client`:
```go
$ go run .
```

## Tools and frameworks used
* Zerlog for logging