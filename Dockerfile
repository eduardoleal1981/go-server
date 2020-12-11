FROM golang:1.14-rc
RUN export GOPATH=/libs/go
COPY main.go /libs/go/src/github.com/eduardoleal1981/go-server
COPY handler /libs/go/src/github.com/eduardoleal1981/go-server/
RUN ls -la /libs/go/src/github.com/eduardoleal1981/go-server

FROM alpine:3.11
COPY --from=0 /libs/go/bin/server-app .
COPY public /public
EXPOSE 8080
CMD ["./go-server"]