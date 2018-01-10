FROM golang:1.9.2-stretch

ENV GOPATH $GOPATH:/go
ENV GOBIN /go/bin

RUN apt-get update && \
    apt-get upgrade -y

ADD . /go/src/omt-project
WORKDIR /go/src/omt-project

RUN go get github.com/codegangsta/gin
RUN curl https://glide.sh/get | sh
RUN glide install

EXPOSE 9000

# For dev hotloading
CMD sh ./script/start-server.sh
