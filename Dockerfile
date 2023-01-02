# syntax=docker/dockerfile:1
FROM golang:1.19-bullseye AS builder

WORKDIR /go/src/github.com/hiroyaonoe/bcop-proxy/

RUN mkdir -p -m 0600 ~/.ssh \
	&& ssh-keyscan github.com >> ~/.ssh/known_hosts \
	&& git config --global url."git@github.com:".insteadOf "https://github.com/"
COPY go.mod go.sum ./
RUN --mount=type=ssh go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bcop-proxy ./cmd/proxy/main.go


FROM scratch

COPY --from=builder /bcop-proxy /bin/bcop-proxy
ENTRYPOINT [ "/bin/bcop-proxy" ]
CMD [ "--proxy-port", "9000", "--admin-port", "9001", "--default-addr", "localhost:9002", "--propagate=true", "--controller-url", "http://localhost:8080", "--id", ""]

