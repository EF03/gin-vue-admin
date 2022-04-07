# 启动, 第一次启动可能会稍微慢一点
docker-compose -f docker-compose-dev.yaml  up

#后台启动
docker-compose -f docker-compose-dev.yaml  up  -d

# 停止
docker-compose -f docker-compose-dev.yaml  stop


docker-compose -f docker-compose-dev.yaml mysql redis


# 使用docker-compose启动四个容器
docker-compose up
# 如果您修改了某些配置选项,可以使用此命令重新打包镜像
docker-compose up --build
# 使用docker-compose 后台启动
docker-compose up -d
# 使用docker-compose 重新打包镜像并后台启动
docker-compose up --build -d
# 服务都启动成功后,使用此命令行可清除none镜像
docker system prune