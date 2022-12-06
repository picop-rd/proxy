FROM golang:1.19-bullseye AS builder

WORKDIR /go/src/github.com/hiroyaonoe/bcop-proxy/
COPY . .

RUN CGO_ENABLED=0 go build -o /bcop-admin ./cmd/bcop-admin/main.go

FROM scratch

COPY --from=builder /bcop-admin /bin/bcop-admin
ENTRYPOINT [ "/bin/bcop-admin" ]
CMD [ "--port", "8080", "--mysql", "user:password@tcp(localhost:3306)/db?parseTime=true&collation=utf8mb4_bin" ]
