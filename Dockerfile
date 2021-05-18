FROM golang:1.16.4-alpine as build
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 go build -o /app/urlshortener

FROM scratch as bin
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/urlshortener /app/urlshortener
ENTRYPOINT ["/app/urlshortener"]