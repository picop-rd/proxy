FROM golang:1.19-bullseye AS builder

WORKDIR /go/src/github.com/picop-rd/proxy/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /proxy ./cmd/proxy/main.go


FROM scratch

COPY --from=builder /proxy /bin/proxy
ENTRYPOINT [ "/bin/proxy" ]
CMD [ "--proxy-port", "9000", "--admin-port", "9001", "--default-addr", "localhost:9002", "--propagate=true", "--controller-url", "http://localhost:8080", "--id", ""]

