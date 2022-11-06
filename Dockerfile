FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o protohedge.api ./cmd/protohedge.api
EXPOSE 8080 

CMD [ "/app/protohedge.api" ]