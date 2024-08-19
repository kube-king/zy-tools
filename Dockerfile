FROM golang:1.18.5 as builder
#Dockerfile 维护者 xiangjiqiang
MAINTAINER xiangjiqiang@qq.com
# 启用go module
ENV GO111MODULE=on
# 设置国内代理镜像地址
ENV GOPROXY=https://goproxy.cn,direct
# Copy 代码到编译环境
ADD . /build/
# 设置工作目录/build
WORKDIR /build
# CGO_ENABLED禁用cgo 然后指定OS等，并go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o main

FROM sit-registry.qm.cn/sit-fc2-qm-monitoring-kube-service/golang-alpine:0.1
# 复制编译的二进制至运行镜像
WORKDIR /app
COPY --from=builder /build/main /app/server
EXPOSE 8080
ENTRYPOINT ["/app/server"]