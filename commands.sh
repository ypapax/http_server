#!/usr/bin/env bash

set -ex

build(){
	go install
}

run() {
	build
	http_server -v 4 -alsologtostderr
}

test(){
	curl localhost:8080/random_path
}

$@