FROM golang:alpine

WORKDIR $GOPATH/src/github.com/ziyitony/simpleItem/userMS
COPY . $GOPATH/src/github.com/ziyitony/simpleItem/userMS
RUN go build -o user

EXPOSE 44444
ENTRYPOINT ["./user"]