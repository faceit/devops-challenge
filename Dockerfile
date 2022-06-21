FROM golang:1.17-alpine

WORKDIR /app

COPY ./test-app/go.mod ./test-app/go.sum ./

RUN go mod download

COPY ./test-app/cmd/ops-test-app/main.go ./

RUN go build -o /test-app

CMD [ "/test-app" ]
