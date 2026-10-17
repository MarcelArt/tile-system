#!/usr/bin/env bash
# Usage: ./build.sh <linux|win> [arch]
#
# Native builds use cgo (raylib compiled from source). Cross builds use
# CGO_ENABLED=0, where raylib-go loads its embedded prebuilt raylib library via
# purego, so no cross C compiler is needed (amd64/arm64 only). Set CGO_ENABLED
# (and CC) yourself to override, e.g. CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc.
set -euo pipefail

target="${1:-}"

case "$target" in
	linux)
		goos="linux"
		out_dir="builds/linux"
		binary="tile-system"
		;;
	win | windows)
		goos="windows"
		out_dir="builds/windows"
		binary="tile-system.exe"
		;;
	*)
		echo "usage: $0 <linux|win> [arch]" >&2
		exit 1
		;;
esac

cd "$(dirname "$0")"

host_os="$(go env GOHOSTOS)"
goarch="${2:-${GOARCH:-$(go env GOHOSTARCH)}}"

if [ -z "${CGO_ENABLED:-}" ]; then
	if [ "$goos" = "$host_os" ]; then
		CGO_ENABLED=1
	else
		CGO_ENABLED=0
	fi
fi

if [ -d "$out_dir" ]; then
	rm -rf "$out_dir"
fi

# -p creates any missing parent directories and is a no-op for existing ones
mkdir -p "$out_dir"

if [ -d assets ]; then
	cp -r assets "$out_dir/assets"
else
	echo "warning: assets directory not found, skipping copy" >&2
fi

echo "building $goos/$goarch (host $host_os, CGO_ENABLED=$CGO_ENABLED)"
GOOS="$goos" GOARCH="$goarch" CGO_ENABLED="$CGO_ENABLED" \
	go build -o "$out_dir/$binary" main.go

echo "built $out_dir/$binary"
