all: spaceshooter.wasm spaceshooter-linux spaceshooter.exe spaceshooter-mac

spaceshooter.wasm: *.go
	mkdir -p dist
	GOOS=js GOARCH=wasm go build -o dist/ ./...
	mv dist/spaceshooter dist/spaceshooter.wasm

spaceshooter-linux: *.go
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/ ./...
	mv dist/spaceshooter dist/spaceshooter-linux

spaceshooter.exe: *.go
	mkdir -p dist
	GOOS=windows GOARCH=amd64 go build -o dist/ ./...

# spaceshooter-mac: *.go
# 	mkdir -p dist
# 	GOOS=darwin GOARCH=amd64 go build -o dist/ ./...
# 	mv dist/spaceshooter dist/spaceshooter-mac
