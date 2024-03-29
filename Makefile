SWAGGER_VERSION=v0.17.0

DEV?=true

app: compile
	bin/greetme -dev=$(DEV)

compile:
	go build -o bin/greetme cmd/*.go

generate: generate-greetme generate-user-service
generate-greetme:
	docker run --rm -it -v $(PWD):/go/src/greetme -w /go/src/greetme -e GO111MODULE=off quay.io/goswagger/swagger:$(SWAGGER_VERSION) generate server --target=pkg --name=greetme --spec=api/greetme/swagger.yml --exclude-main
generate-user-service:
	docker run --rm -it -v $(PWD):/go/src/greetme -w /go/src/greetme -e GO111MODULE=off quay.io/goswagger/swagger:$(SWAGGER_VERSION) generate server --target=pkg/mocks/services/user --name=mock_user_service --spec=api/mocks/user.yml --exclude-main

collection:
	newman run postman/collection.json -e postman/local.environment.json
