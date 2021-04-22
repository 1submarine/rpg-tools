OS=darwin
ARCH=arm64

all: compile

compile:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o rmanage *.go
	
install: compile
	go install

debug:
	dlv debug github.com/1submarine/rpg-tools
