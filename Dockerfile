FROM golang:1.14-rc
COPY handler /go/src/github.com/eduardoleal1981/go-server/handler
COPY main.go /go/src/github.com/eduardoleal1981/go-server/
COPY go.mod /go/src/github.com/eduardoleal1981/go-server/
RUN ls -la /go/src/github.com/eduardoleal1981/
RUN ls -la /go/src/github.com/eduardoleal1981/go-server
RUN ls -la /go/bin
RUN go install github.com/eduardoleal1981/go-server

FROM alpine:3.11
COPY --from=0 /go/bin/go-server .
COPY public .
EXPOSE 8080
CMD ["./go-server"]