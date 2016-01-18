#!/bin/bash
set -e

mkdir -p bin

go generate

TARGETS=(darwin/amd64 linux/amd64)

for target in ${TARGETS[@]}; do
  export GOOS=${target%/*}
  export GOARCH=${target##*/}
  go build
  mv dockerfile-validator bin/dockerfile-validator"$GOOS"_"$GOARCH"
done
