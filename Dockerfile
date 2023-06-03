FROM golang:1.20.1-alpine as builder
LABEL maintainer="Carl Bergman <carl@cbergman.se>"

WORKDIR /app

ENV PDFJS_VERSION=3.4.120
ENV SQLC_VERSION=1.18.0

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN --mount=type=cache \
    apk add --no-cache curl unzip git

RUN --mount=type=cache \
    go install github.com/kyleconroy/sqlc/cmd/sqlc@v${SQLC_VERSION}

RUN curl -LO https://github.com/mozilla/pdf.js/releases/download/v${PDFJS_VERSION}/pdfjs-${PDFJS_VERSION}-dist.zip && \
    unzip pdfjs-${PDFJS_VERSION}-dist.zip -d ./_content/js/pdfjs/ && \
    rm pdfjs-${PDFJS_VERSION}-dist.zip

WORKDIR /app/internal/service/
RUN sqlc generate

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /am-reader ./cmd/site/
RUN chmod +x /am-reader

FROM scratch as site
COPY --from=builder /am-reader /am-reader
COPY --from=builder /app/internal/service/init.sql /init.sql
COPY --from=builder /app/internal/service/setup.sql /setup.sql
WORKDIR /database
WORKDIR /uploads

EXPOSE 8080

CMD ["/am-reader"]
