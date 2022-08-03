# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as dev

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 50051
CMD ["go", "run", "cmd/app/main.go"]


FROM golang:1.18-alpine as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary ${ROOT}/cmd/server/


FROM scratch as prod

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 50051
CMD ["/go/src/app/binary"]
