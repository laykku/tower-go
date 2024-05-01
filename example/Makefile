all: bundle run

run:
	cd build && python -m http.server 8080

gen:
	go generate

bundle:
	GOOS=js GOARCH=wasm go build -o build/exampleapp.wasm
