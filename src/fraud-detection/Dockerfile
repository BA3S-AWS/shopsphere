FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o fraud-detection .

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/fraud-detection /app/fraud-detection

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/fraud-detection"]
