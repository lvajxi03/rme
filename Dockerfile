FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/rme .

FROM alpine:3.21

RUN adduser -D -H -u 10001 appuser
COPY --from=builder /out/rme /usr/local/bin/rme

USER 10001

EXPOSE 9101

ENTRYPOINT ["/usr/local/bin/rme"]
