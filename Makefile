build:
	go build -o bin/fixie-wrench

clean:
	rm -r -f ./bin $(EXEC)

compile:
	GOOS=linux GOARCH=arm64 go build -o bin/fixie-wrench-linux-arm64
	GOOS=darwin GOARCH=arm64 go build -o bin/fixie-wrench-macos-arm64
	GOOS=linux GOARCH=amd64 go build -o bin/fixie-wrench-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/fixie-wrench-macos-amd64
	GOOS=windows GOARCH=amd64 go build -o bin/fixie-wrench-windows-amd64

release:
	make clean
	make compile
	cp ./fixie-wrench.sh ./bin/fixie-wrench