FROM golang:1.13-alpine as builder
RUN apk --no-cache add ca-certificates
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main k8socopus/*.go
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/main /app/
WORKDIR /app
ENTRYPOINT ["./main"]
CMD []