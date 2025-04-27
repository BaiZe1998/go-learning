# 使用官方 Node.js 镜像作为基础镜像
FROM node:16-alpine

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json 到工作目录
COPY package*.json ./

# 安装依赖
RUN npm install

# 复制应用程序代码到工作目录
COPY . .

# 暴露应用程序运行的端口
EXPOSE 3000

# 定义环境变量
ENV NODE_ENV=production

# 运行应用程序
CMD ["npm", "start"] 