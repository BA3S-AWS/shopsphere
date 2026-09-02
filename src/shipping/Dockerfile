FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o shipping .

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/shipping /app/shipping

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/shipping"]
