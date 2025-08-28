FROM golang:1.24-alpine AS builder
WORKDIR /src
RUN apk add --no-cache git ca-certificates
RUN addgroup -S app && adduser -S -G app app
ENV HOME=/home/app \
    GOPATH=/home/app/go \
    GOCACHE=/home/app/.cache/go-build
RUN mkdir -p /out "$GOPATH" "$GOCACHE" && chown -R app:app /src /out /home/app
USER app:app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/api

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --chown=nonroot:nonroot --from=builder /out/app /app/app
COPY --chown=nonroot:nonroot --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]