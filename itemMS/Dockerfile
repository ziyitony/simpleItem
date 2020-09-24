FROM golang:alpine

WORKDIR $GOPATH/src/github.com/ziyitony/simpleItem/itemMS
COPY . $GOPATH/src/github.com/ziyitony/simpleItem/itemMS
RUN go build -o item

EXPOSE 55555
ENTRYPOINT ["./item"]
