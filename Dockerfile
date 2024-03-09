ARG BASE_IMAGE=golang:1.22.0-alpine3.19

FROM ${BASE_IMAGE} AS builder

WORKDIR /go/src/app

ADD . /go/src/app/

RUN go mod download && go mod verify

RUN go build -v -o /stori ./cmd/transactions/*.go

FROM ${BASE_IMAGE} AS candidate

ARG NAME=stori

WORKDIR /go/src/app

COPY --from=builder /$NAME .

EXPOSE 8080

ENTRYPOINT ["./stori"]