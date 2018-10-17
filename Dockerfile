From golang:onbuild
WORKDIR /go/src/test_proc
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000



