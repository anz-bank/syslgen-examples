#!/bin/sh

set -e -x

TAG=$1
COMMAND=$2

usage() {
    echo "Usage: sh build-image.sh <tag> (build|deploy)"
    echo "eg: sh build-image.sh 0.0.1 build"  
}

if [ -z "$TAG" ]; then
    echo "Tag is not specified."
    usage
    exit 1
fi

if [ "$COMMAND" != "build" ] && [ "$COMMAND" != "deploy" ]; then
    echo "Invalid command or command not specified."
    usage
    exit 1
fi

SYSL2_VERSION=0.2.5-ubuntu18.04
DOCKER_IMAGE="anzbank/syslgen-examples"
DOCKER_IMAGE_VERSIONED="$DOCKER_IMAGE:${TAG}"

build() {
    docker build --build-arg SYSL2_VERSION=$SYSL2_VERSION -t "$DOCKER_IMAGE" -f Dockerfile .
    docker tag "$DOCKER_IMAGE" "$DOCKER_IMAGE_VERSIONED"
}

deploy() {
    docker push "$DOCKER_IMAGE_VERSIONED" && docker push "$DOCKER_IMAGE"
}

if [ "$COMMAND" = "build" ]; then
    build
elif [ "$COMMAND" = "deploy" ]; then
    deploy
fi
