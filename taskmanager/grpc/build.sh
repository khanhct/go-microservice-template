#!/bin/bash
declare -a services=("protos")
DESTDIR='build'
for entry in "${services[@]}"/*
do
  echo ${entry}
  if [ -d "${entry}" ]; then
    dir_name=$(basename ${entry})
    GO_OUT=${DESTDIR}/${dir_name}
    echo $GO_OUT
    mkdir ${GO_OUT}
    protoc \
    --proto_path=$entry/ \
    --go-grpc_out=$GO_OUT \
    $entry/*.proto
  fi
done

for SERVICE in "${services[@]}"; do
    protoc \
        --proto_path=$SERVICE/ \
        --go_out=$DESTDIR \
        --go-grpc_out=$DESTDIR \
        $SERVICE/*.proto
done
