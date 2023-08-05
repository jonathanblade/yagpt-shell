VERSION=0.1.0-dev

BUILD_LDFLAGS="-s -w -X 'main.Version=${VERSION}'"

build-mac-arm64:
	$(eval EXE := yagpt)
	GOOS=darwin GOARCH=arm64 go build -ldflags=${BUILD_LDFLAGS} -o ${EXE} main.go
