build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build .
run:
	./goweb
start:
	make build
	make run

