FROM golang:alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g" /etc/apk/repositories && \
    apk add --no-cache upx

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -tags timetzdata -ldflags="-s -w" -o /go/bin/app
RUN upx /go/bin/app

FROM alpine

WORKDIR /app

ENTRYPOINT ["./go"]

COPY --from=builder /go/bin/app go
