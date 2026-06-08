FROM golang:1.25.9-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/w3r5l3y-ssh ./cmd/server


FROM alpine:3.21

RUN addgroup -g 568 apps && adduser -D -H -u 568 -G apps appuser

WORKDIR /app

COPY --from=build /out/w3r5l3y-ssh /app/w3r5l3y-ssh

RUN mkdir -p /data && chown -R 568:568 /data

USER appuser

EXPOSE 23234

ENV SSH_ADDR=0.0.0.0:23234
ENV SSH_HOST_KEY_PATH=/data/ssh_host_ed25519_key

ENTRYPOINT ["/app/w3r5l3y-ssh"]