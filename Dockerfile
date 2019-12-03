FROM golang:1.12.5 as builder
RUN mkdir /go/src/tcs
WORKDIR /go/src/tcs
COPY . .
#RUN go get strategy/service
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /tcs tcs/service

FROM alpine:latest
RUN apk add ca-certificates
WORKDIR /usr/local/bin
COPY --from=builder /tcs .
COPY ./config.prod.ini ./config.ini
COPY ./migrations ./migrations
EXPOSE 9000
CMD ["/usr/local/bin/tcs"]
