FROM golang:1.22 as build
WORKDIR /build/src
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-a -s -w" -o app .

FROM scratch
COPY --from=build /build/src/app /usr/bin/app
ENTRYPOINT ["/usr/bin/app"]

