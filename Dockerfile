FROM golang:1.20-alpine as builder
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o "cd-docker" .


FROM docker:latest
COPY --from=builder /app/cd-docker /
COPY --from=builder /app/_configFiles/config.ini /_configFiles/
COPY --from=builder /app/templates /templates
CMD ["./cd-docker"]
