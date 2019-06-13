#!/bin/sh

ARCHS="darwin linux freebsd windows"
BINDIR="bin"
BINARY="super-hacker"

if [ -z $1 ]; then 
    echo "error: requires argument of [release|freebsd|darwin|linux|windows]"
    exit 1
fi

if [ -z $2 ]; then 
    echo "error: requires version argument"
    exit 1
fi
VERSION=$2

LDFLAGS="-X main.gitSHA=$(git rev-parse HEAD) -X main.version=${VERSION} -X main.name=${BINARY}"

if [ $1 == "release" ]; then
    echo "Generating ${BINARY} release binaries..."
    for arch in ${ARCHS}; do
        GOOS=${arch} GOARCH=amd64 go build -v -ldflags "${LDFLAGS}" -o ${BINDIR}/${BINARY}-${arch}
    done
fi

case "$1" in
    "release") 
        echo "Building release..."
        for arch in ${ARCHS}; do
            GOOS=${arch} GOARCH=amd64 go build -v -ldflags "${LDFLAGS}" -o ${BINDIR}/${BINARY}-${arch}
            tar -czvf ${BINDIR}/${BINARY}-${arch}.tar.gz ${BINDIR}/${BINARY}-${arch}
        done
        ;;
    "freebsd") 
        echo "Building binary for FreeBSD..."
        GOOS=freebsd GOARCH=amd64 go build -v -ldflags "${LDFLAGS}" -o ${BINDIR}/${BINARY}-freebsd
        ;;
    "darwin") 
        echo "Building binary for Darwin..."
        GOOS=darwin GOARCH=amd64 go build -v -ldflags "${LDFLAGS}" -o ${BINDIR}/${BINARY}-darwin
        ;;
    "linux") 
        echo "Building binary for Linux..."
        GOOS=linux GOARCH=amd64 go build -v -ldflags "${LDFLAGS}" -o ${BINDIR}/${BINARY}-linux
        ;;
    "windows") 
        echo "Building binary for Windows..."
        GOOS=windows GOARCH=amd64 go build -v -ldflags "${LDFLAGS}" -o ${BINDIR}/${BINARY}-windows.exe
        ;;
esac

exit 0
