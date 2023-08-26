clean:
	rm -rf ./bin

build: clean
	GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bin/aluno/bootstrap src/aluno/*
	cd ./bin/aluno && zip aluno.zip bootstrap

	GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bin/escola/bootstrap src/escola/*
	cd ./bin/escola && zip escola.zip bootstrap

deploy: build
	sls deploy --verbose

remove:
	sls remove --verbose