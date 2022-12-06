.PHONY:run-admin
run-admin: fmt
	go run cmd/bcop-admin/main.go --port 8080

.PHONY:run-proxy
run-proxy: fmt
	go run cmd/bcop-proxy/main.go --port 8081 --default localhost:9000 --propagate true

.PHONY:build
build: test
	go build

.PHONY:test
test: vet fmt
	go test ./...

.PHONY:test-with-coverage
test-with-coverage: vet fmt
	go test -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

.PHONY:vet
vet:
	go vet ./...

.PHONY:fmt
fmt:
	go fmt ./...
