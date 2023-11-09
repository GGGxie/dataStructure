# 使用 golang 官方镜像作为基础镜像
FROM golang:1.21.4 AS build

# 设置工作目录
WORKDIR /app

# 复制 Go 项目文件到镜像中
COPY go.mod go.sum ./
RUN go mod download
COPY ./main.go /app/

# 下载 Go 项目依赖（如果使用 Go Modules）
RUN go mod download

# 构建 Go 程序
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# 第二阶段：构建最终镜像
FROM scratch

# 设置工作目录
WORKDIR /app

# 从第一阶段复制构建好的 Go 二进制文件
COPY --from=build /app/myapp .