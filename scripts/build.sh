#!/usr/bin/env bash

if [ -z "$GOOS" ]; then
	echo "Error: GOOS environment variable is not set."
	exit 1
fi

if [ -z "$GOARCH" ]; then
	echo "Error: GOOS environment variable is not set."
	exit 1
fi

set +e
rm -rf output

set -e
mkdir -p output/atsa-notifier
if [ "$GOOS" = "windows" ]; then
	go build -o ./output/atsa-notifier.exe ./cmd/atsa-notifier
else
	go build -o ./output/atsa-notifier ./cmd/atsa-notifier
fi
cp -r ./web ./output/atsa-notifier
cp -r ./conf ./output/atsa-notifier
cp -r ./data ./output/atsa-notifier
cd ./output
if [ "$GOOS" = "linux" ]; then
	tar zcvf "atsa-notifier-${GOOS}-${GOARCH}.tar.gz" ./atsa-notifier
else
	zip -r "atsa-notifier-${GOOS}-${GOARCH}.zip" ./atsa-notifier
fi
rm -rf atsa-notifier
cd ..
