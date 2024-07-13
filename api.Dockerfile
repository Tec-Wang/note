# 构建api服务使用的dockerfile
FROM golang:1.21 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build
# 执行go build . 的时候 这个 . 代表了上下文路径
ADD /go.mod .
ADD /go.sum .
RUN go mod download
COPY .. .
RUN go build -ldflags="-s -w" -o /app/note ./api

# 二阶段构建，第一阶段构建的内容，最终只需要build出来的可执行文件就可以了
# 第二阶段是复制这个可执行文件，然后进行运行。

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/note /app/note

CMD ["./note"]
