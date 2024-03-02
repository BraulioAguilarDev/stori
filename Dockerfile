ARG BASE_IMAGE=golang:1.22.0-alpine3.19

FROM ${BASE_IMAGE} AS builder

WORKDIR /go/src/app

ADD . /go/src/app/

RUN go mod download && go mod verify

RUN go build -v -o /stori ./cmd/transactions/main.go

FROM ${BASE_IMAGE} AS candidate

ARG NAME=stori

# Install nice to haves
RUN apk add --no-cache openssl ncurses-libs libstdc++ libgcc curl libressl htop nano

WORKDIR /go/src/app

COPY --from=builder /$NAME .
# COPY --from=builder /go/src/app/db/migrations /go/src/app/db/migrations
# COPY --from=builder /go/src/app/template /go/src/app/template
# COPY --from=builder /go/src/app/firebase-admin.json /go/src/app/firebase-admin.json

EXPOSE 8080

ENTRYPOINT ["./stori"]