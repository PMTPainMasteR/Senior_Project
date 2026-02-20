FROM golang:1.24.12-alpine AS builder

WORKDIR /app

RUN apk add --no-cache \
    git=2.52.0-r0 \
    curl=8.17.0-r1 \
    bash=5.3.3-r1

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o wifi-sim .

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/wifi-sim /wifi-sim

USER nonroot:nonroot

ENTRYPOINT ["/wifi-sim"]
