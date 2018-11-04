FROM golang:1.11.1-alpine3.7 as builder

WORKDIR /go/src/pihouse
COPY . .

WORKDIR /go/src/pihouse/pihouseclient

RUN export GOOS=linux
RUN export GOARCH=arm
RUN export GOARM=6
RUN apk add --no-cache git mercurial \
    && go get -d -v \
    && apk del git mercurial

RUN go build -v

FROM arm32v6/alpine:3.7

WORKDIR /pihouse

COPY --from=builder /go/src/pihouse/pihouseclient .

CMD ["/bin/bash"]