# Build
FROM golang:1.13.5-alpine3.10 as builder

RUN apk update && apk upgrade && \
    apk --update add git gcc make

WORKDIR /go/src/github.com/jayvib/gophr

COPY . .

ENV GO111MODULE on

RUN make build

# Distribute
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    apk add --no-cache bash && \
    mkdir /gophr

WORKDIR /gophr

EXPOSE 8080

COPY --from=builder /go/src/github.com/jayvib/gophr/gophr.linux /gophr

CMD ["/gophr/gophr.linux"]
