# 容器知识入门教程

容器技术是现代应用开发和部署的基石，本教程将介绍 Docker 容器的基础知识。

## Docker 背景介绍

### 什么是 Docker

Docker 是一个开源的应用容器引擎，让开发者可以打包他们的应用以及依赖包到一个可移植的容器中，然后发布到任何流行的 Linux 或 Windows 操作系统的机器上。Docker 使用了 Linux 内核的多种功能（如 Namespaces、Cgroups）来创建独立的容器。

### Docker 的发展历史

- **2013年**：Docker 由 dotCloud 公司（后更名为 Docker Inc.）推出，最初是 dotCloud 平台的内部项目
- **2014年**：Docker 1.0 发布，正式进入生产环境
- **2015年**：Docker Compose、Docker Swarm 和 Docker Machine 等工具发布，生态系统开始繁荣
- **2016年**：引入内置的编排功能
- **2017年**：集成 Kubernetes 支持
- **至今**：持续迭代发展，成为容器化技术的事实标准

### 为什么需要 Docker

在 Docker 出现之前，开发者面临以下问题：

1. **环境不一致**：开发、测试、生产环境的差异导致"在我电脑上能运行"的问题
2. **部署复杂**：应用依赖安装复杂，配置繁琐
3. **资源利用率低**：传统虚拟化方案资源占用高，启动慢
4. **应用隔离困难**：不同应用之间相互影响
5. **扩展性差**：难以快速扩容和缩容

Docker 通过容器化技术解决了这些问题：

1. **一致的环境**：无论在哪里运行，容器内的环境都是一样的
2. **轻量级**：容器共享主机系统内核，比传统虚拟机占用资源少，启动更快
3. **隔离**：容器之间彼此隔离，不会相互影响
4. **可移植性**：构建一次，到处运行
5. **微服务支持**：适合现代微服务架构，每个服务独立容器化

### Docker vs 虚拟机

|特性|Docker 容器|虚拟机|
|---|---|---|
|启动时间|秒级|分钟级|
|存储空间|MB级|GB级|
|性能|接近原生|有所损耗|
|系统资源|共享宿主机内核|独立内核|
|隔离性|进程级隔离|完全隔离|
|运行密度|单机可运行数十至数百个容器|单机通常运行数个虚拟机|

## 1. Docker 安装

### Windows 安装

1. 下载 [Docker Desktop for Windows](https://www.docker.com/products/docker-desktop)
2. 双击安装程序并按照提示完成安装
3. 安装完成后，Docker Desktop 会自动启动
4. 在系统托盘中可以看到 Docker 图标，表示 Docker 服务正在运行

### macOS 安装

1. 下载 [Docker Desktop for Mac](https://www.docker.com/products/docker-desktop)
2. 将下载的 `.dmg` 文件拖到应用程序文件夹
3. 启动 Docker Desktop 应用
4. 等待 Docker 初始化完成，顶部状态栏会显示 Docker 图标

### Linux 安装 (Ubuntu)

```bash
# 更新软件包索引
sudo apt-get update

# 安装依赖
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# 添加 Docker 官方 GPG 密钥
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# 设置稳定版仓库
echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# 安装 Docker Engine
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io

# 将当前用户添加到 docker 组（避免每次都使用 sudo）
sudo usermod -aG docker $USER
# 需要重新登录使配置生效
```

### 验证安装

```bash
# 查看 Docker 版本
docker --version

# 运行测试容器
docker run hello-world
```

## 2. 常用 Docker 命令

### 镜像相关命令

```bash
# 搜索镜像
docker search ubuntu

# 拉取镜像
docker pull ubuntu:latest

# 列出本地镜像
docker images

# 删除镜像
docker rmi ubuntu:latest

# 构建镜像
docker build -t myapp:1.0 .
```

### 容器相关命令

```bash
# 创建并启动容器
docker run -d -p 8080:80 --name mywebserver nginx

# 列出所有运行中的容器
docker ps

# 列出所有容器（包括已停止的）
docker ps -a

# 停止容器
docker stop mywebserver

# 启动已停止的容器
docker start mywebserver

# 重启容器
docker restart mywebserver

# 删除容器
docker rm mywebserver

# 进入容器交互式终端
docker exec -it mywebserver bash

# 查看容器日志
docker logs mywebserver

# 查看容器资源使用情况
docker stats mywebserver
```

### Docker Compose 命令

```bash
# 启动所有服务
docker-compose up -d

# 停止所有服务
docker-compose down

# 查看服务状态
docker-compose ps

# 查看服务日志
docker-compose logs

# 重建服务
docker-compose build
```

### 网络和存储命令

```bash
# 创建网络
docker network create mynetwork

# 列出网络
docker network ls

# 创建卷
docker volume create mydata

# 列出卷
docker volume ls
```

## 示例项目说明

本目录包含一个简单的 Docker 示例项目，包括：

1. **Dockerfile**: 定义如何构建应用容器镜像
2. **docker-compose.yml**: 定义多容器应用的服务、网络和卷
3. **app.js**: 一个简单的 Node.js Express 应用
4. **package.json**: Node.js 应用依赖定义

### 如何运行示例项目

1. 确保已安装 Docker 和 Docker Compose
2. 在此目录下运行：
   ```bash
   docker-compose up -d
   ```
3. 访问应用：
   - Web 应用: http://localhost:3000
   - 健康检查: http://localhost:3000/health

### 项目文件说明

#### Dockerfile

这个文件定义了如何构建应用容器：
- 使用 Node.js 16 Alpine 作为基础镜像
- 设置工作目录
- 复制和安装依赖
- 配置环境变量
- 指定启动命令

#### docker-compose.yml

这个文件定义了完整的应用栈：
- Web 应用服务 (使用 Dockerfile 构建)
- PostgreSQL 数据库服务
- Redis 缓存服务
- 网络配置
- 卷配置（持久存储）

#### app.js

一个简单的 Express 服务器，展示容器环境信息和健康检查端点。

## Docker 网络详解

Docker 网络是容器化环境中的关键组件，提供了容器间的通信基础设施。

### Docker 网络的主要作用

1. **容器间通信**：允许不同容器在不暴露端口到主机的情况下相互通信
2. **隔离环境**：可以创建完全隔离的网络环境，提高应用安全性
3. **服务发现**：容器可以通过容器名称而非IP地址相互访问，简化服务发现
4. **多主机连接**：使用overlay网络可以连接不同主机上的容器
5. **网络策略控制**：可以精细控制哪些容器可以相互通信

### 常用网络类型

```bash
# 查看可用的网络驱动
docker info | grep "Network"

# 创建自定义网络
docker network create --driver bridge my-network

# 在创建容器时连接到指定网络
docker run --network=my-network -d --name container1 nginx

# 将已有容器连接到网络
docker network connect my-network container2
```

#### 网络驱动类型

- **bridge**: 默认网络驱动，适用于同一主机上的容器
- **host**: 直接使用主机网络，移除容器与主机间的网络隔离
- **overlay**: 用于Docker Swarm环境中，连接多个Docker守护进程
- **macvlan**: 允许容器拥有独立的MAC地址，直接连接到物理网络
- **none**: 禁用所有网络

## Docker 卷详解

Docker卷提供了容器数据的持久化存储解决方案，解决了容器销毁后数据丢失的问题。

### Docker 卷的主要作用

1. **数据持久化**：即使容器被删除，存储在卷中的数据依然保留
2. **数据共享**：多个容器可以挂载相同的卷，实现数据共享
3. **备份与恢复**：简化数据备份和恢复流程
4. **性能优化**：与容器内部存储相比，卷通常提供更好的I/O性能
5. **存储解耦**：将应用与数据分离，提高系统灵活性和可维护性

### 卷的使用方式

```bash
# 创建卷
docker volume create my-data

# 查看卷的详细信息
docker volume inspect my-data

# 在容器中使用卷
docker run -d --name my-container -v my-data:/app/data nginx

# 使用绑定挂载（挂载主机目录）
docker run -d --name my-container -v $(pwd):/app/data nginx

# 备份卷中的数据
docker run --rm -v my-data:/source -v $(pwd):/backup alpine tar -czvf /backup/my-data-backup.tar.gz -C /source .
```

#### 卷的类型

- **命名卷**: 由Docker管理的命名存储空间
- **绑定挂载**: 直接映射主机文件系统的路径到容器
- **tmpfs挂载**: 将数据存储在主机的内存中，不写入文件系统

使用卷时要考虑权限、备份策略和数据生命周期管理，以确保数据安全和可靠性。

## Docker 进阶概念

1. **Docker 多阶段构建**：优化镜像大小和构建过程
2. **Docker 网络**：bridge、host、overlay 等不同网络驱动
3. **Docker 卷**：持久化数据存储
4. **Docker Swarm**：Docker 原生集群和编排工具
5. **Docker 安全**：最佳实践和安全配置
6. **Docker Registry**：镜像仓库和分发

## 相关资源

- [Docker 官方文档](https://docs.docker.com/)
- [Docker Hub](https://hub.docker.com/)
- [Docker Compose 文档](https://docs.docker.com/compose/)
- [Dockerfile 最佳实践](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)

## 练习建议

1. 尝试为不同语言的应用创建 Dockerfile
2. 练习使用 Docker Compose 设置多容器应用
3. 探索 Docker 卷和网络
4. 学习如何在生产环境中安全地部署 Docker 容器