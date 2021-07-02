FROM golang:1.16.5-stretch

WORKDIR /go/src
COPY . .

RUN CGO_ENABLED=0 go get -d -v ./...
RUN CGO_ENABLED=0 go install -v ./...
