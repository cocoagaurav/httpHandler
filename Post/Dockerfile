From golang:onbuild

WORKDIR /go/src/cocoagaurav/httpHandler/Post
COPY . .

RUN go get -d -v ./...
RUN go install ./

EXPOSE 8081