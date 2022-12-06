.PHONY:run-admin
run-admin: fmt
	go run cmd/bcop-admin/main.go --port 8080 --mysql "bcop:BCoP-2022@tcp(localhost:3306)/bcop?parseTime=true&collation=utf8mb4_bin"

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
