#build stage
FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

#run stage
RUN go build -o ./bin/main

FROM alpine:latest

WORKDIR /app

RUN mkdir template

COPY --from=build /app/template ./template

COPY --from=build /app/bin/main .

CMD ["./main"]
