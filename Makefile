GO=go
BIN=out
OPTIONS=-ldflags \"-s -w\" -buildvcs=false -trimpath
OPTIONS=

all: vendor build

.PHONY: vendor
vendor:
	${GO} mod vendor
	${GO} mod tidy

.PHONY: build
build:
	mkdir -p ${BIN}
	${GO} build -o ${BIN}/snake cmd/snake/main.go ${OPTIONS} 

.PHONY: clean
clean:
	rm -rf ${BIN}
