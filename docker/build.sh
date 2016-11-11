#!/bin/sh
parent=`dirname $PWD`

build_torsten() {
    docker run --rm -v "$parent":/go/src/github.com/kildevaeld/apprun \
    -v gobuilder:/go \
    -w /go/src/github.com/kildevaeld/apprun/apprun \
    -ti \
    --name apprun-builder \
    kildevaeld/go-builder sh -c make

    mv -f ../apprun/apprun apprun
    
    
}



build() {
    build_torsten
    docker build --tag kildevaeld/apprun .
}


build

