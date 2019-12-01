FROM  golang:1.12 AS builder
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demo .
RUN swag init


FROM  alpine:3.10

# RUN apk --no-cache add ca-certificates

WORKDIR /app

EXPOSE 3000
COPY --from=builder /build/demo /app/
COPY --from=builder /build/docs /app/docs

CMD ["./demo"]