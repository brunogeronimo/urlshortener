FROM golang:1.16.4-alpine as build
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 go build -o /app/urlshortener

FROM golang:1.16.4-alpine as bin
COPY --from=build /app/configs/ /app/configs/
COPY --from=build /app/urlshortener /app/urlshortener
ENTRYPOINT ["/app/urlshortener"]