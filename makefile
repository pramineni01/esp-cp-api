build:
	cd cmd/cp-api && go get github.com/gobuffalo/packr/v2/packr2 && packr2 && go build -o esp-cp-api
fmt:
	cd cmd/cp-api && go fmt
generate:
	cd cmd/cp-api && go run github.com/99designs/gqlgen generate
run:
	cd cmd/cp-api && go run server.go
