FROM golang:1.10.1-alpine3.7 AS b0
ADD main.go .
RUN apk add --no-cache git ca-certificates
RUN go get github.com/gorilla/mux \
           github.com/GeertJohan/go.rice
RUN rice embed-go
RUN go build -v -o cuesite
FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY --from=b0 /go/cuesite .
ENTRYPOINT ./cuesite
