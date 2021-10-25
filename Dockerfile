
########
## Build
########

FROM golang:1.16-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-sampleapp

#########
## Deploy
#########

FROM alpine:3.14


WORKDIR /

COPY --from=builder /go-sampleapp .

EXPOSE 8080

CMD [ "/go-sampleapp" ]
