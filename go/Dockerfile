FROM golang:1.9

ARG package

WORKDIR /go/src/github.com/mmcken3/cufeedback/go
COPY ./go /go/src/github.com/mmcken3/cufeedback/go
RUN go-wrapper install "github.com/mmcken3/cufeedback/go/cufeedback/cmd/$package"