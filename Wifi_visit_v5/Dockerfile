FROM golang:1.24.11-alpine AS builder

WORKDIR /Wifi_visit_v5

RUN apk add --no-cache git

COPY go.mod ./

RUN go mod download

COPY . .dockerignore ./

RUN CGO_ENABLED=0 go build -o wifi-sim .

FROM scratch

COPY --from=builder /app/wifi-sim /wifi-sim

ENTRYPOINT ["/wifi-sim"]