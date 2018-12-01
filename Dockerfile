From golang

WORKDIR /go/src/cocoagaurav/httpHandler
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080