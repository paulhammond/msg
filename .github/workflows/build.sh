#!/bin/sh

set -ex

DIR=msg-${OS}-${REF}
rm -rf "$DIR"
mkdir "$DIR"

cp README.md "$DIR"

cp LICENSE.txt "$DIR"

GOOS=${OS}
[ "$GOOS" = "mac" ] && GOOS=darwin
GOARCH=amd64 GOOS="$GOOS" go build -o "${DIR}/msg" -ldflags "-X github.com/paulhammond/msg/internal/msg.version=${VERSION}" ./cmd/msg

tar -czf "${DIR}.tgz" "$DIR"