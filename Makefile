build:
	go build -o bin/fixie

clean:
	rm -r -f ./bin $(EXEC)

compile:
	GOOS=linux GOARCH=arm64 go build -o bin/fixie-linux-arm64
	GOOS=darwin GOARCH=arm64 go build -o bin/fixie-macos-arm64
	GOOS=linux GOARCH=amd64 go build -o bin/fixie-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/fixie-macos-amd64
	GOOS=windows GOARCH=amd64 go build -o bin/fixie-windows-amd64

release:
	make clean
	make compile
	cp ./fixie.sh ./bin