FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o payment .

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/payment .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./payment"]
