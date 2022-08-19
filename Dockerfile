FROM golang:latest
MAINTAINER reddyprabhakar528@gmail.com

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/BoggalaPrabhakar007/golang-assignment/cmd/main
RUN cd /build && git clone https://github.com/BoggalaPrabhakar007/golang-assignment.git

RUN cd /build/golang-assignment/cmd && go build

EXPOSE 8080

ENTRYPOINT ["/build/golang-assignment/cmd/cmd"]