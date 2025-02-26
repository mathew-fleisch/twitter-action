# build stage
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR $GOPATH/src/github.com/mathew-fleisch/twitter-action/
COPY . .

RUN go get . && CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /go/bin/twitter-action *.go

# final stage
FROM scratch
LABEL maintainer="Mathew Fleisch <mathew.fleisch@gmail.com>"

LABEL "com.github.actions.name"="Twitter Action"
LABEL "com.github.actions.description"="Update Status (tweet) on behalf of a user"
LABEL "com.github.actions.icon"="twitter"
LABEL "com.github.actions.color"="blue"

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /go/bin/twitter-action /usr/bin/twitter-action
ENTRYPOINT ["twitter-action"]
