From golang:1.11

WORKDIR /go/src/github.com/cocoagaurav/httpHandler/Post
COPY . .

RUN go get -d -v ./Post/...
RUN go install ./post/

EXPOSE 8081