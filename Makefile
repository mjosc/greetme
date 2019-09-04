SWAGGER_VERSION=v0.17.0

compile-mockserver:
	go build -o bin/swagger-mockserver mockserver/cmd/mockserver-server/*.go

generate: generate-greetme generate-mockserver
generate-greetme:
	docker run --rm -it -v $(PWD):/go/src/greetme -w /go/src/greetme -e GO111MODULE=off quay.io/goswagger/swagger:$(SWAGGER_VERSION) generate server --target=pkg --name=greetme --spec=api/greetme/swagger.yml --exclude-main
generate-mockserver:
	docker run --rm -it -v $(PWD):/go/src/greetme -w /go/src/greetme -e GO111MODULE=off quay.io/goswagger/swagger:$(SWAGGER_VERSION) generate server --target=pkg/mocks --api-package=mockops --server-package=mockapi --name=mock --spec=api/mockserver/swagger.yml --exclude-main