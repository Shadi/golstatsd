FROM golang:alpine

ADD ./main.go /go/src/app/main.go
WORKDIR /go/src/app

ENV PORT=8125

CMD ["go", "run", "main.go"]