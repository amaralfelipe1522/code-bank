FROM golang:1.15

WORKDIR /go/src
COPY . .
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

CMD [ "go", "run", "main.go" ]