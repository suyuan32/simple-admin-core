FROM golang:1.19.1-alpine3.16 as builder

WORKDIR /home
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o /home/core_api core.go

FROM alpine:latest

WORKDIR /home

COPY --from=builder /home/core_api ./
COPY --from=builder /home/etc/core.yaml ./

EXPOSE 9100
ENTRYPOINT ./core_api -f core.yaml