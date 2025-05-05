这个案例中，大部分代码取自 Eino 官方案例库：https://github.com/cloudwego/eino-examples
在此基础上进行了补充和文档的完善。

在运行案例之前，先确保：
1. 执行 docker-compose up -d 确保 redis 启动成功，可以查看：http://127.0.0.1:8001/redis-stack/browser
2. 执行 source .env 确保环境变量设置正确