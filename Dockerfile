FROM golang:1.14


WORKDIR /go/src/hoya-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go" , "run", "main.go"]
