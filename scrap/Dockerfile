From golang:1.11

WORKDIR /go/src/github.com/cocoagaurav/httpHandler/scrap
COPY . .

RUN go get -d -v ./scrap/...
RUN go install ./scrap/

EXPOSE 8082