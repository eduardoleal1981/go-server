FROM golang:1.14-rc
COPY server/go/src/server-app /go/src/server-app
RUN go install server-app

FROM alpine:3.11
COPY --from=0 /go/bin/server-app .
COPY img /public/img
COPY docs /public/docs/pdf
COPY js /public/js
COPY css /public/css/normalize/normalize.css
COPY css /public/css/main.css
COPY index.html /public/index.html
EXPOSE 8080
CMD ["./server-app"]