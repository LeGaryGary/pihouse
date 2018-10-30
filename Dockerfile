FROM arm32v6/golang:1.11.1-alpine3.7 as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v 
RUN go build -v

FROM arm32v6/alpine:3.7

COPY --from=builder . .

CMD ["./pihouse"]