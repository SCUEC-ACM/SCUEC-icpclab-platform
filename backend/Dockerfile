FROM golang:1.25-alpine AS builder

WORKDIR /app

# 1. 设置代理（必须在 tidy 之前）
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 2. 先拷贝依赖文件
COPY go.mod go.sum ./

# 3. 这一步非常重要：拷贝所有源代码
# 只有拷贝了源代码，go mod tidy 才能识别出你代码里用了 gorm
COPY . .

# 4. 强制更新 go.mod 并下载缺失的包
RUN go mod tidy

# 5. 执行编译
RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /main /main
EXPOSE 3000
CMD ["/main"]