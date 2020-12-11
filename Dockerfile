FROM golang:1.14-rc
RUN go install github.com/eduardoleal1981/go-server/server-app

FROM alpine:3.11
COPY --from=0 /go/bin/server-app .
COPY img /img
COPY docs /docs/pdf
COPY js /js
COPY css /css
COPY index.html /index.html
EXPOSE 80
CMD ["./server-app"]