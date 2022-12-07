.PHONY:run-proxy
run-proxy: fmt
	go run cmd/proxy/main.go --proxy-port 9000 --admin-port 9001 --default-addr localhost:9002 --propagate true

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
