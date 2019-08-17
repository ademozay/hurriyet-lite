FROM golang:1.11.1 as builder
WORKDIR /go/src/app
COPY . .
RUN go get github.com/gorilla/mux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api . 

FROM alpine as certs
RUN apk update && apk add ca-certificates

FROM alpine
COPY --from=builder /go/src/app/api /
COPY --from=certs /etc/ssl/certs /etc/ssl/certs
ENTRYPOINT [ "/api" ]