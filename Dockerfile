FROM golang:1.14-alpine

#RUN apk update && apk upgrade && \
#    apk add --no-cache bash git openssh

ENV GOPROXY https://goproxy.cn

WORKDIR /data/go_docker/

ADD  ./main main

EXPOSE 8080

RUN apk add --no-cache bash git openssh

ENTRYPOINT ["./main"]