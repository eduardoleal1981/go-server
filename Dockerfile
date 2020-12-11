FROM golang:1.14-rc
WORKDIR /libs/go
RUN export GOPATH=/libs/go
RUN go install github.com/eduardoleal1981/go-server

FROM alpine:3.11
COPY --from=0 /libs/go/bin/server-app .
COPY public /public
EXPOSE 8080
CMD ["./go-server"]