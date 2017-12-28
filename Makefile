NAME=netlink
IMPORT=github.com/mickep76/$(NAME)

all: format build readme

format:
	gofmt -w=true .

test:
	golint -set_exit_status
	go vet .
#       go test

readme:
	godoc2md $(IMPORT) >README.md
	cat README.footer >>README.md

build: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
