FROM golang:1.24-alpine AS builder

WORKDIR /Wifi_visit_v5

RUN apk add --no-cache git

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -Idflags="-s -w" -o wifi-sim .

FROM scratch

COPY --from=builder /app/wifi-sim /wifi-sim

ENTRYPOINT [ "/wifi-sim" ]