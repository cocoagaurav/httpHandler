From golang:1.11

WORKDIR /go/src/github.com/cocoagaurav/httpHandler/scrap
COPY . .

RUN go get -d -v ./...
RUN go install ./

EXPOSE 8082