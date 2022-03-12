#!/usr/bin/env bash
export REPO=edgi-io
export TAG=foobar
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
docker run --rm -it $(docker build -q $SCRIPT_DIR/$1)
