FROM golang:1.18.3-alpine3.16

WORKDIR /usr/src/app

RUN adduser -s /sbin/nologin testapp --no-create-home --disabled-password

COPY go.mod ./

COPY go.sum ./

RUN go mod download

RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/src/app/ops-test-app ./cmd/ops-test-app

ENTRYPOINT ["/usr/src/app/ops-test-app"]
