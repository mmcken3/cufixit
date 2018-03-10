FROM golang:1.9

ARG package

WORKDIR /go/src/github.com/cufeedback/go
COPY ./go /go/src/github.com/cufeedback/go
RUN go-wrapper install "/go/src/github.com/cufeedback/go/cufeedback/cmd/$package"