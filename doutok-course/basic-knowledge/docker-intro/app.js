const express = require('express');
const app = express();
const port = process.env.PORT || 3000;

app.get('/', (req, res) => {
  res.send(`
    <h1>Docker 示例应用</h1>
    <p>这是一个运行在 Docker 容器中的简单 Node.js 应用。</p>
    <p>环境: ${process.env.NODE_ENV || '开发'}</p>
    <p>主机名: ${process.env.HOSTNAME || '未知'}</p>
    <p>当前时间: ${new Date().toLocaleString()}</p>
  `);
});

app.get('/health', (req, res) => {
  res.status(200).json({ status: 'ok', timestamp: new Date().toISOString() });
});

app.listen(port, () => {
  console.log(`应用正在监听端口 ${port}`);
  console.log(`环境: ${process.env.NODE_ENV || '开发'}`);
}); 