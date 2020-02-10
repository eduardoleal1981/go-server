FROM golang:1.14-rc
COPY server/go/src/main /go/src/main
RUN go install main

FROM alpine:3.11
COPY --from=0 /go/bin/main .
COPY api /public/api
COPY png /public/png
COPY svg /public/svg
COPY js /public/js
COPY css /public/css
COPY index.html /public/index.html
EXPOSE 8080
CMD ["./main"]
