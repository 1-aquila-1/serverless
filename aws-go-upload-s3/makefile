clean:
	rm -rf ./bin

build: clean
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/upload upload/*
	tar.exe -a -c -f bin/function.zip bin/*

deploy: clean build
	sls deploy --verbose

remove:
	sls remove --verbose