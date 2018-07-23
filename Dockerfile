FROM golang:1.10-alpine3.7

LABEL maintainer 'Lucas Marques <lucasmarques73@hotmail.com>'

RUN apk add git

ADD ./src /go/src

ADD .env /go/src/api/.env

WORKDIR /go/src/api

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep status && dep ensure

# RUN go run Infra/creatingDatabase.go

EXPOSE 80

# CMD [ "go run main.go" ]