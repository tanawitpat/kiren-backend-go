FROM golang:1.13-alpine

WORKDIR /go/src/kiren-backend-go
COPY . .
RUN go build -o app .

CMD ["/go/src/kiren-backend-go/app"]