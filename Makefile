include .env

.EXPORT_ALL_VARIABLES:

dev:
	go run cmd/cmd.go

prod:
	CGO_ENABLED=0
	go build -o ./bin/go-api-boilerplate ./cmd/cmd.go

upx: prod
	upx ./bin/go-api-boilerplate

docker:
	docker build -t douglaszuqueto/go-api-boilerplate .
	docker push douglaszuqueto/go-api-boilerplate

.PHONY: dev prod upx docker