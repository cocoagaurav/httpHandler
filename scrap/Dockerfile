
WORKDIR /go/src/cocoagaurav/httpHandler/scrap
COPY . .

RUN go get -d -v ./...
RUN go install ./

EXPOSE 8082