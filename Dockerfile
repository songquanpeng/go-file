FROM golang AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -ldflags "-s -w -extldflags '-static'" -o go-file

FROM scratch

ENV PORT=3000
COPY --from=builder /build/go-file /
EXPOSE 3000
ENTRYPOINT ["/go-file"]
