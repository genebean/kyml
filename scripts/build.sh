#!/bin/bash

set -eo pipefail

if [ "$TRAVIS" = "true" ]; then
	if [ "$TRAVIS_TAG" != "" ]; then
		if [[ $TRAVIS_TAG != v* ]]; then
			echo >&2 "TRAVIS_TAG is set but does not start with \"v\""
			exit 1
		fi
		version="${TRAVIS_TAG:1}"
	else
		version="untagged"
	fi
else
	version=${1:?"Version (arg 1) missing"}
fi

echo "Version: $version"

ldflags="-X github.com/frigus02/kyml/pkg/commands.version=$version"

echo "Build darwin"
GOARCH=amd64 GOOS=darwin go build -o "bin/kyml_${version}_darwin_amd64" -ldflags "$ldflags"

echo "Build linux"
GOARCH=amd64 GOOS=linux go build -o "bin/kyml_${version}_linux_amd64" -ldflags "$ldflags"

echo "Build windows"
GOARCH=amd64 GOOS=windows go build -o "bin/kyml_${version}_windows_amd64.exe" -ldflags "$ldflags"

echo "Create checksums"
(
	cd bin/
	sha256sum "kyml_${version}"* >checksums.txt
)
