FROM golang:1.9.5-alpine3.7 AS b0
ADD main.go /tmp/main.go
RUN cd /tmp && go build -v -o srvsptfy
FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY --from=b0 /tmp/srvsptfy .
ENTRYPOINT ./srvsptfy
