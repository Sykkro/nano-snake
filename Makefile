GO=go
BIN=bin
APP=snake
OPTIONS=-ldflags \"-s -w\" -buildvcs=false -trimpath

.PHONY: build
build:
	mkdir -p ${BIN}
	${GO} build cmd/${APP} -o ${BIN}/${APP} ${OPTIONS} 
