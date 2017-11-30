FROM golang:1.9.2-stretch

ENV GOPATH $GOPATH:/go

RUN apt-get update && \
    apt-get upgrade -y

RUN go get github.com/codegangsta/gin

ADD . /go/src/omt-project
WORKDIR /go/src/omt-project

EXPOSE 9000

# For dev hotloading
CMD sh start-server.sh
