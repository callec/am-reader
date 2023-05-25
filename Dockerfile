FROM golang:1.20.1-alpine as builder
LABEL maintainer="Carl Bergman <carl@cbergman.se>"

WORKDIR /app

ENV PDFJS_VERSION=3.4.120
ENV SQLC_VERSION=1.18.0

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN apk add --no-cache curl tar git

RUN curl -L https://github.com/mozilla/pdf.js/archive/refs/tags/v${PDFJS_VERSION}.tar.gz | tar xz && \
    mv pdf.js-${PDFJS_VERSION}/ /app/_content/js/pdfjs/

RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@v${SQLC_VERSION}

WORKDIR /app/service
RUN sqlc generate

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /am-reader ./cmd/site/
RUN chmod +x /am-reader

FROM scratch
COPY --from=builder /am-reader /am-reader
COPY --from=builder /app/service /service
WORKDIR /database
WORKDIR /uploads

EXPOSE 8080

CMD ["/am-reader"]
