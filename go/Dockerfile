FROM golang:1.9

ARG package

RUN apt-get update && apt-get install -y postgresql

WORKDIR /go/src/github.com/mmcken3/cufixit/go
COPY ./go /go/src/github.com/mmcken3/cufixit/go
RUN go-wrapper install "github.com/mmcken3/cufixit/go/cufixit/cmd/$package"