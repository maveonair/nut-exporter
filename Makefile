.PHONY:  build clean dev test release

default: build

build: clean
	CGO_ENABLED=0 go build -o ./dist/nut-exporter -a -ldflags '-s' -installsuffix cgo cmd/nut-exporter/main.go

clean:
	rm -rf ./dist/*

dev:
	gow run cmd/nut-exporter/main.go

test:
	go test -v ./...

release: clean
	goreleaser release
