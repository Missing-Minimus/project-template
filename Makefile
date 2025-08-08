.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bootstrap ./main.go

clean:
	rm -f bootstrap

deploy: clean build
	sh proddeploy.sh

deploy-dev: clean build
	sh devdeploy.sh
