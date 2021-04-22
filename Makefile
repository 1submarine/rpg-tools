OS=darwin
ARCH=arm64

all: compile

compile:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o rpg-tools *.go
	
install: compile
	go install github.com/1submarine/rpg-tools

debug:
	dlv debug github.com/1submarine/rpg-tools
