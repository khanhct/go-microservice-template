#!/bin/bash
# LINIUX 
# curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
# echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
# apt-get update
# apt-get install -y migrate

go get -u -d github.com/golang-migrate/migrate/cmd/migrate
cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
export TAG=v4.6.2
git checkout $TAG
# Go 1.16+
go install -tags 'postgres' -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate@$TAG