FROM golang:1.17-alpine3.16 as builder
WORKDIR /build
COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux go build -o products-rest-api cmd/products-rest-api/main.go

FROM alpine:3.11.3
COPY --from=builder /build/products-rest-api /app/products-rest-api

EXPOSE 8080

CMD [ "/app/products-rest-api" ]
