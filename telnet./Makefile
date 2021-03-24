name = telnet-cli

build:
	GOOS=windows GOARCH=amd64 go build -o $(name).win64.exe main.go
	GOOS=windows GOARCH=386 go build -o $(name).win32.exe main.go
	go build -o $(name).linux.bin main.go


