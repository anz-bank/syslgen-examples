#!/bin/sh

set -e -x

TAG=$1

SYSL2_VERSION=0.2.5-ubuntu18.04
DOCKER_IMAGE_LATEST="anzbank/syslgen-examples"
DOCKER_IMAGE_VERSIONED="$DOCKER_IMAGE_LATEST:${TAG}"

docker build --build-arg SYSL2_VERSION=$SYSL2_VERSION -t "$DOCKER_IMAGE_LATEST" -f Dockerfile .
docker tag "$DOCKER_IMAGE_LATEST" "$DOCKER_IMAGE_VERSIONED"
