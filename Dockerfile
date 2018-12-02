From golang:1.11

WORKDIR /go/src/github.com/cocoagaurav/httpHandler
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["httpHandler"]

EXPOSE 8080