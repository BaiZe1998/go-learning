# 定义项目名称
PROJECT_NAME := mongodb

# 默认目标
.DEFAULT_GOAL := help

# 定义命令
DOCKER_COMPOSE := docker-compose

# 配置 Docker Compose 文件路径
DOCKER_COMPOSE_FILE := docker-compose.yaml

# 启动 MongoDB 服务
up:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up -d

# 停止 MongoDB 服务
down:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

# 清理 MongoDB 数据卷
clean:
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down -v

# 帮助信息
help:
	@echo "Usage:"
	@echo "  make up         Start MongoDB container"
	@echo "  make down       Stop MongoDB container"
	@echo "  make clean      Remove MongoDB data volume"
	@echo "  make help       Show this help message"
