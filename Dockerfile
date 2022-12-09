FROM golang:1.19-bullseye AS builder

WORKDIR /go/src/github.com/hiroyaonoe/bcop-proxy/
COPY . .

RUN CGO_ENABLED=0 go build -o /bcop-proxy ./cmd/proxy/main.go


FROM scratch

COPY --from=builder /bcop-proxy /bin/bcop-proxy
ENTRYPOINT [ "/bin/bcop-proxy" ]
CMD [ "--proxy-port", "9000", "--admin-port", "9001", "--default-addr", "localhost:9002", "--propagate", "true" ]

