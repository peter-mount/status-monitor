#!/bin/sh

docker build -t test . &&\
docker run \
  -it --rm \
  -v /tmp/statusdb:/database \
  test
