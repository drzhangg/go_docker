FROM golang:1.14-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

ENV GOPROXY https://goproxy.cn

COPY ../go.mod .

RUN go mod download

COPY docker .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]