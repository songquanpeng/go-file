# Build stage
FROM golang AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -ldflags "-s -w -X 'go-file/common.Version=$(cat VERSION)' -extldflags '-static'" -o go-file

# Final stage
FROM alpine

RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates 2>/dev/null || true

ENV PORT=3000
COPY --from=builder /build/go-file /
WORKDIR /data
EXPOSE 3000
ENTRYPOINT ["/go-file"]
