#!/bin/bash

# 微服务的根目录
ROOT_PATH="/root/goProjects/gomall/app"

# 微服务列表，每个微服务的目录名称
declare -a SERVICES=("cart" "checkout" "frontend" "order" "payment" "product" "user")

# 启动所有微服务
for service in "${SERVICES[@]}"; do
    echo "Starting $service..."
    # 进入微服务目录
    cd "$ROOT_PATH/$service"
    # 使用 nohup 启动微服务，并将输出重定向到日志文件
    nohup air > "$service.log" 2>&1 &
    # 返回到脚本开始时的目录
    cd - > /dev/null
    echo "$service started."
done

echo "All services have been started."