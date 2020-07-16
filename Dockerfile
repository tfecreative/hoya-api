FROM golang:1.14


WORKDIR /go/src/hoya-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go get github.com/go-delve/delve/cmd/dlv github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/hoya-api/" -command="./hoya-api"