GO := go
NAME := samplewasm
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)

build: wasm server

start: server
	./server &

server: cmd/server/main.go
	go build -o server cmd/server/main.go

wasm: site/static/main.wasm

site/static/main.wasm: cmd/wasm/main.go
	GOOS=js GOARCH=wasm go build -o site/main.wasm cmd/wasm/main.go

clean:
	rm site/static/$(NAME).wasm