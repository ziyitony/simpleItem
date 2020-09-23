FROM golang:alpine

WORKDIR $GOPATH/src/github.com/ziyitony/simpleItem
COPY . $GOPATH/src/github.com/ziyitony/simpleItem
RUN go build -o main

EXPOSE 12345
ENTRYPOINT ["./main"]