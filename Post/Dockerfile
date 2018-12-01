From golang

WORKDIR /go/src/github.com/cocoagaurav/httpHandler/Post
COPY . .

RUN go get -d -v ./...
RUN go install ./

EXPOSE 8081