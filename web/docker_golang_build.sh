#!/bin/bash

_GO_VERSION=1.9.2
_IMAGE='maiastra/hvd_server'
_TAG='0.1'
docker run --rm -v $(pwd):/usr/src/myapp -w /usr/src/myapp golang:$_GO_VERSION bash -c ./build.sh

echo "building new image"
docker build -t $_IMAGE:$_TAG .
echo "image build you can start with :"
echo "docker run -d -p 8888:8080 $_IMAGE:$_TAG"
