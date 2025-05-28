FROM golang:tip-alpine AS build

WORKDIR /app

COPY go.mod . 

COPY main.go .

RUN go mod tidy

RUN go build -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/app .

EXPOSE 8080
CMD ["./app"]
