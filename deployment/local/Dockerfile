FROM golang:1.13-alpine
WORKDIR /go/src/kiren-backend-go

COPY vendor vendor
COPY main.go main.go
COPY internal internal
COPY configs configs
COPY data data

RUN go build -o app .
CMD ["/go/src/kiren-backend-go/app"]
