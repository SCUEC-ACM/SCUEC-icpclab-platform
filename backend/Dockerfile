# 第一阶段：编译
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# 即使这里 COPY 了 .env，它也只存在于 builder 阶段
RUN go build -o /main main.go

# 第二阶段：运行（最终生成的镜像只包含这个阶段）
FROM alpine:latest
WORKDIR /
# 只从编译阶段拷贝二进制文件，不拷贝任何源代码或 .env
COPY --from=builder /main /main
EXPOSE 3000
CMD ["/main"]