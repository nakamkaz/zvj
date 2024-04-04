PROJ_NAME := "zvj"
VERSION := "1.0.3-20240404"
SRCS := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -extldflags \"-static\""

default: $(SRCS)
	go fmt 
	go build $(LDFLAGS) 
	GOOS=windows GOARCH=amd64 go build -ldflags="main.Version=$(VERSION)"

initmod: 
	go mod init $(PROJ_NAME)
