OS=darwin
ARCH=

all: compile

compile:
	go build -o rmanage *.go
