FROM docker.io/library/golang:1.19-alpine

WORKDIR /go/src/app

COPY *.go ./

RUN go mod init && \
    go build -o /go/src/app/log-spitter

CMD [ "/go/src/app/log-spitter" ]
