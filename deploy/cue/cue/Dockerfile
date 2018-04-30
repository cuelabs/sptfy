FROM golang:1.10.1-alpine3.7 AS b0
RUN echo http://mirror.yandex.ru/mirrors/alpine/v3.7/main >> /etc/apk/repositories
RUN echo http://mirror.yandex.ru/mirrors/alpine/v3.7/community >> /etc/apk/repositories
ADD . . 
RUN apk add --no-cache git ca-certificates
RUN go get github.com/kelseyhightower/envconfig \
           github.com/cuelabs/sptfy/pkg/user \
           github.com/gorilla/mux \
           github.com/lib/pq \
           github.com/jmoiron/sqlx \
           github.com/GeertJohan/go.rice \
           github.com/GeertJohan/go.rice/rice \
           github.com/llater/cue/models \
           github.com/llater/cue/app \
           github.com/llater/cue/environment
RUN rice embed-go
RUN go build -v -o cue
FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY --from=b0 /go/cue .
ENTRYPOINT ./cue
