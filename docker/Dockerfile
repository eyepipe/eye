FROM golang:1.24 AS builder
ARG LDFLAGS
ARG GIT_COMMIT
ARG VERSION

WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/ ./cmd
COPY internal/ ./internal
COPY pkg/ ./pkg
RUN go mod tidy
# enable cgo for go-sqlite3 and goose
RUN CGO_ENABLED=1 go build -o build/eye-server -ldflags="${LDFLAGS}" cmd/server/*.go
RUN CGO_ENABLED=0 go build -o build/eye-cli -ldflags="${LDFLAGS}" cmd/cli/*.go

FROM debian:bookworm-slim
ARG GIT_COMMIT
ARG VERSION
LABEL org.opencontainers.image.source="https://github.com/eyepipe/eye" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.version="${VERSION} commit ${GIT_COMMIT}" \
      org.opencontainers.image.description="God sees everything, except what's encrypted."

WORKDIR /app
RUN apt update && \
    apt install -y --no-install-recommends ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/build/eye-cli /usr/local/bin/cli
COPY --from=builder /app/build/eye-server /usr/local/bin/server
COPY config/config.example.yml ./config/config.yml
COPY db/migrations db/migrations
RUN mkdir db/databases
ENTRYPOINT ["/usr/local/bin/cli"]
