FROM golang:1.11

WORKDIR /go/src/api
COPY . .

RUN go build -o main . 

ENTRYPOINT [ "/go/src/api/main" ]