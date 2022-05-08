FROM golang:1.18.1-alpine3.14 AS builder

WORKDIR /usr/src/app

RUN apk add --no-cache git

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
COPY . .

RUN CGO_ENABLED=0 && GOOS=linux && GOARCH=amd64
RUN go mod download && go mod verify

RUN go build -v -tags netgo -o /usr/local/bin/ ./...
#CMD ["log-server"]

# target stage
FROM scratch

WORKDIR /root/
COPY --from=builder /usr/local/bin/log-server /usr/local/bin/
CMD ["log-server"]
