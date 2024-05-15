FROM golang:1.22-alpine as base

WORKDIR /app
COPY . .
RUN go mod download
RUN go build cmd/api/main.go
FROM alpine

WORKDIR /app
COPY --from=base /app/main .

EXPOSE 80
ENV GIN_MODE release

CMD [ "./main" ]'
