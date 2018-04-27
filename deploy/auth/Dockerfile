FROM golang:1.10.1-alpine3.7
RUN apk add --no-cache git ca-certificates
RUN go get -u github.com/bitly/oauth2_proxy
#ENV OAUTH2_PROXY_CLIENT_ID
#ENV OAUTH2_PROXY_CLIENT_SECRET
ENTRYPOINT oauth2_proxy  --upstream="http://cuesite:10001/" \
                         --provider=github \
                         --login-url=https://accounts.spotify.com/authorize \
                         --http-address="http://0.0.0.0:4180" \
                         --email-domain=* \
                         --cookie-secure=false \
                         --cookie-secret=this-is-a-secret
