From golang:onbuild
WORKDIR /go/src/cocoagaurav/httpHandler
COPY . .

RUN go get -d -v ./...
RUN go install ./Post

EXPOSE 8081