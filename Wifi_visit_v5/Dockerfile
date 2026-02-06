FROM golang:1.24.12-alpine AS builder

WORKDIR /app

RUN apk add --no-cache "git" "curl=7.79.1-r0" "bash=5.1.8-r0"

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o wifi-sim .

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/wifi-sim /wifi-sim

USER nonroot:nonroot

ENTRYPOINT ["/wifi-sim"]