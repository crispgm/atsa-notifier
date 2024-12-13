#!/usr/bin/env bash

set +e
rm -rf output

set -e
mkdir -p output/atsa-notifier
go build -o ./output/atsa-notifier ./cmd/atsa-notifier
cp -r ./web ./output/atsa-notifier
cp -r ./conf ./output/atsa-notifier
cp -r ./data ./output/atsa-notifier
cd ./output
zip -r "atsa-notifier-${GOOS}-${GOARCH}.zip" ./atsa-notifier
rm -rf atsa-notifier
cd ..
