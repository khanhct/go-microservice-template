#!/bin/bash
declare -a services=("protos")
DESTDIR='build'
for entry in "${services[@]}"/*
do
  echo ${entry}
  if [ -d "${entry}" ]; then
    dir_name=$(basename ${entry})
    PY_OUT=${DESTDIR}/${dir_name}
    mkdir ${PY_OUT}
    protoc \
    --proto_path=$entry/ \
    --go_out=$PY_OUT \
    --go-grpc_out=$PY_OUT \
    $entry/*.proto
  fi
done

for SERVICE in "${services[@]}"; do
    protoc \
        --proto_path=$SERVICE/ \
        --go_out=$DESTDIR \
        --grpc_out=$DESTDIR \
        $SERVICE/*.proto
done
