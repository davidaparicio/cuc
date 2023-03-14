#!/usr/bin/env bash

VERSION=$(git name-rev --tags --name-only "$(git rev-parse HEAD)")
GIT_COMMIT=$(git rev-list -1 HEAD)
BUILT_DATE=$(date)
PACKAGE="github.com/davidaparicio/cuc/internal"
#GO111MODULE="" "on"
#TZ="Europe/Paris"

LDFLAGS="-s -w -X '${PACKAGE}.Version=${VERSION}' -X '${PACKAGE}.GitCommit=${GIT_COMMIT}' -X '${PACKAGE}.BuiltDate=${BUILT_DATE}'"
CGO_ENABLED=1
TARGETOS="darwin" #GOOS=linux or GOOS=windows
TARGETARCH="amd64" #go env 32bit: GOARCH=386
#echo ${LDFLAGS}

#https://blog.alexellis.io/inject-build-time-vars-golang/
CGO_ENABLED="${CGO_ENABLED}" GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" \
go build -ldflags \
"${LDFLAGS}" \
-a -installsuffix cgo -o ./cmd/cuc .