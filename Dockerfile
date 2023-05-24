FROM golang:1.20.1-alpine as builder
LABEL maintainer="Carl Bergman <carl@cbergman.se>"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /am-reader ./cmd/site/
RUN chmod +x /am-reader

FROM scratch
COPY --from=builder /am-reader /am-reader
COPY --from=builder /app/service /service
COPY --from=builder /app/database /database
COPY --from=builder /app/database /uploads

EXPOSE 8080

CMD ["/am-reader"]
