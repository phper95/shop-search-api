# FROM 基于 golang:1.16-alpine
FROM golang:1.16-alpine AS builder

# ENV 设置环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOPRIVATE=*gitee.com

# COPY 源路径 目标路径
COPY ../ /usr/local/src/shop-search-api

# RUN 执行 go build .
RUN cd /usr/local/src/shop-search-api && go build main.go

# FROM 基于 alpine:latest
FROM alpine:latest

# RUN 设置代理镜像
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories

# RUN 设置 Asia/Shanghai 时区
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# COPY 源路径 目标路径 从镜像中 COPY
COPY --from=builder  /usr/local/src/shop-search-api /usr/local/shop-search-api

# EXPOSE 设置端口映射
EXPOSE 9999/tcp

# WORKDIR 设置工作目录
WORKDIR /opt

# CMD 设置启动命令
CMD ["./shop-search-api"]