FROM golang:1.9

ARG package

WORKDIR /go/src/github.com/mmcken3/cufixit/go
COPY ./go /go/src/github.com/mmcken3/cufixit/go
RUN go-wrapper install "github.com/mmcken3/cufixit/go/cufixit/cmd/$package"