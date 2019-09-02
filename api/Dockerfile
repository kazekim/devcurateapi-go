FROM golang:1.12.7

WORKDIR /go/src/github.com/kazekim/devcurateapi-go/api
COPY . .

RUN go get github.com/pilu/fresh
RUN go get ./...

CMD [ "fresh" ]