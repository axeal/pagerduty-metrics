FROM golang:1.13.5 AS builder

WORKDIR $GOPATH/src/github.com/axeal/pagerduty-metrics
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app ./
ENTRYPOINT ["./app"]