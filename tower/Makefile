all: bundle run

run:
	python -m http.server 8080

gen:
	go generate

bundle:
	GOOS=js GOARCH=wasm go build -o towerapp.wasm
