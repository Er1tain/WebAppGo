FROM golang:1.16 AS build-env

WORKDIR /WebbAppGo

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o WebbAppGo .

FROM alpine:latest

WORKDIR /WebbAppGo

COPY --from=build-env /WebbAppGo/WebbAppGo .

CMD ["./WebbAppGo"]