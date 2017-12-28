NAME=netlink
IMPORT=github.com/mickep76/$(NAME)

all: test build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

all: test readme

format:
	gofmt -w=true .

test: format
	golint -set_exit_status
#       go vet .
#       go test

readme:
	cat README.header >README.md
	godoc2md $(IMPORT) >>README.md

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
