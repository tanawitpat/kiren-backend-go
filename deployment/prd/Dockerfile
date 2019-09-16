FROM golang:1.13-alpine
WORKDIR /go/src/kiren-backend-go
COPY . .
RUN go build -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/src/kiren-backend-go/app .
COPY mocked mocked
COPY configs configs
CMD ["./app"]
