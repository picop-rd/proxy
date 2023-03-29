.PHONY:docker-build
docker-build: test
	docker build -t proxy --ssh default .

.PHONY:docker-run
docker-run:
	docker run -p 9000:9000 -p 9001:9001 proxy

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
