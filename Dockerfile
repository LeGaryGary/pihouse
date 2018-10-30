FROM golang:1.11.1-alpine3.7 as builder

WORKDIR /go/src/pihouse
COPY . .

RUN apk add --no-cache git mercurial \
    && go get -d -v \
    && apk del git mercurial
ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=6
RUN go build -v

FROM arm32v6/alpine:3.7

WORKDIR /go/src/pihouse

COPY --from=builder . .

CMD ["./pihouse"]