FROM golang:alpine

WORKDIR /go/src/simpleItem
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["simpleItem"]
