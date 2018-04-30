FROM golang:1.10.1-alpine3.7
RUN apk add --no-cache ca-certificates git protobuf
RUN go get github.com/twitchtv/twirp/protoc-gen-twirp \
           github.com/golang/protobuf/protoc-gen-go \
           github.com/cuelabs/sptfy/rpc/sptfyapi
#RUN mv /go/src/github.com/cuelabs/sptfy/rpc/sptfyapi/sptfyapi.proto .
COPY sptfyapi.proto .
RUN protoc --go_out=. --twirp_out=. sptfyapi.proto
