FROM golang:1.23.2 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /app/main /
CMD ["/main"]
