BUILDDIR=$(shell pwd)/cmd/gost
VERSION=2.12.2

build() {
    os=$1
    arch=$2
    ext=""
	if [[ "$os" == "windows" ]]; then
		ext=".exe"
    fi

    GOOS=$os GOARCH=$arch go build -v -ldflags "-w -s" -trimpath -o $BUILDDIR/gost$VERSION.$os-$arch$ext $BUILDDIR
}

build windows amd64
build windows 386
build linux amd64
build linux 386
build linux arm64
build linux arm
build darwin amd64