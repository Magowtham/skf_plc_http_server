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

ENV LOG_FILE_NAME=logger.txt
ENV SERVER_ADDRESS=0.0.0.0:8000
ENV CACHE_URL=redis://default:cYwSyVUqS5EGME0XSUOWxqfmg5y4HYVj@redis-18967.c62.us-east-1-4.ec2.redns.redis-cloud.com:18967
ENV DATABASE_URL=postgresql://admin:7U45l0VZOdyQRvEIif0OmWnhW2j8UnY7@dpg-csb4kkrtq21c73980160-a.singapore-postgres.render.com/vsense
ENV SECRETE_KEY=vsense2024
ENV SMTP_USERNAME=cshubhanga@gmail.com
ENV SMTP_PASSWORD=yvzggalzaxjlwotv
ENV SMTP_SERVICE_HOST=smtp.gmail.com
ENV SMTP_SERVICE_PORT=587

CMD ["./main"]
