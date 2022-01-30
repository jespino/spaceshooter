all: spaceshooter.wasm spaceshooter-linux spaceshooter.exe spaceshooter-mac

spaceshooter.wasm: *.go
	GOOS=js GOARCH=wasm go build -o spaceshooter.wasm *.go

spaceshooter-linux: *.go
	GOOS=linux GOARCH=amd64 go build -o spaceshooter-linux *.go

spaceshooter.exe: *.go
	GOOS=windows GOARCH=amd64 go build -o spaceshooter.exe *.go

spaceshooter-mac: *.go
	GOOS=darwin GOARCH=amd64 go build -o spaceshooter-mac *.go
