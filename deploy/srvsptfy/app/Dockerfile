FROM golang:1.10.1-alpine3.7 AS b0
ADD main.go /tmp/main.go
RUN apk add --no-cache git ca-certificates
RUN go get github.com/gorilla/mux
RUN cd /tmp && go build -v -o srvsptfy
FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY --from=b0 /tmp/srvsptfy .
ENTRYPOINT ./srvsptfy
