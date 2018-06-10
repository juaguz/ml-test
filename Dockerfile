FROM golang:1.10
WORKDIR /go/src/github.com/juaguz/ml-test
ADD ./ .
RUN go build -o main main.go

EXPOSE 8080
ENTRYPOINT ["./main", "-api=true"]