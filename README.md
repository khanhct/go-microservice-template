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

## How to run
```sh
# Example
go get google.golang.org/grpc/examples/helloworld/helloworld
go get -u -d github.com/golang-migrate/migrate/cmd/migrate

export GOBIN=./bin
go build cmd/api
go build cmd/manage
go build cmd/taskmanaget

# Example
./api
./manager --version
./taskmanager
```

# Buy Me A Coffe
Email: trongkhanh.chu@gmail.com
## BTC
[<img src="assets/BTC.jpg">](BTC)

## USDT
[<img src="assets/USDT.jpg">](USDT)
