FROM golang:1.18.3-alpine3.16

WORKDIR /app

RUN apk add build-base

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./cmd/protohedge.api/protohedge-api ./cmd/protohedge.api
EXPOSE 8080 

WORKDIR /app/cmd/protohedge.api

CMD [ "./protohedge-api" ]

