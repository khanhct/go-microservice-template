# Best Go Micro Service template

## Architecture
### 1. Gin Restful API
### 2. Grpc server
### 3. Database mysql(migrate, downgrate)
### 4. Loging logrus
### 5. Configuration
### ....

## How to update modules
```sh
go mod tiny
```

## How to build protoc 

```sh
cd taskmanager/grpc
./clear.sh
./build.sh
```

## How to run
```sh
docker-compose up 

go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u -d github.com/golang-migrate/migrate/cmd/migrate

export GOBIN=./bin
go build -o ${GOBIN} cmd/api/
go build -o ${GOBIN} cmd/manage/
go build -o ${GOBIN} cmd/taskmanage/

./bin/api
./bin/manage --help
./bin/taskmanager
```

## Test HTTP Request
```sh
curl -X GET http://127.0.0.1:8090/casorder/api/v1/health/check
```

## Test Grpc
```sh
./bin/manage test-grpc
```


# Buy Me A Coffee
## BTC
[<img src="assets/BTC.jpg">](BTC)

## USDT
[<img src="assets/USDT.jpg">](USDT)