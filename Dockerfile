FROM golang:1.20.4-alpine as builder

WORKDIR /usr/app/imaginator

COPY go.* ./

RUN go mod tidy

COPY . .

RUN go build -o .bin/main cmd/app/main.go

FROM alpine

WORKDIR /usr/app/imaginator

COPY --from=builder /usr/app/imaginator/.bin/main ./.bin/main

CMD [".bin/main"]