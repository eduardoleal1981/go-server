FROM golang:1.14-rc
WORKDIR /libs/go
RUN export GOPATH=/libs/go
RUN go get github.com/eduardoleal1981/go-server

FROM alpine:3.11
COPY --from=0 /libs/go/bin/server-app .
COPY img /img
COPY docs /docs/pdf
COPY js /js
COPY css /css
COPY index.html /index.html
EXPOSE 80
CMD ["./server-app"]