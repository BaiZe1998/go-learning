# React 入门教程

React 是由 Facebook 开发的一个用于构建用户界面的 JavaScript 库。React 使开发人员能够构建快速、可扩展的 Web 应用程序。

## 1. 下载与环境配置

要开始使用 React，您需要一个现代的 JavaScript 环境和 Node.js 开发环境。

### 安装 Node.js 和 pnpm

首先，安装 Node.js（包含 npm）：

- **Windows/macOS**: 从 [Node.js 官网](https://nodejs.org/) 下载并安装
- **macOS (使用 Homebrew)**: `brew install node`
- **Linux**: `sudo apt install nodejs npm` 或 `sudo yum install nodejs npm`

验证安装：

```bash
node -v
```

然后安装 pnpm (性能更好的包管理器)：

```bash
# 使用npm安装pnpm
npm install -g pnpm

# 验证pnpm安装
pnpm --version
```

### 创建新的 React 应用

使用 Create React App 工具创建新的 React 应用：

```bash
# 创建新项目
pnpm create react-app my-react-app

# 进入项目目录
cd my-react-app

# 启动开发服务器
pnpm start
```

或者使用 Vite 创建（更快的启动速度，推荐）：

```bash
# 使用 Vite 创建 React 项目
pnpm create vite my-react-app --template react

# 进入项目目录
cd my-react-app

# 安装依赖
pnpm install

# 启动开发服务器
pnpm dev
```

## 2. React 基本语法与包管理

### 基本概念

1. **组件 (Components)**：React 应用由组件构成
   - 函数组件（推荐）
   - 类组件

2. **JSX**：JavaScript 的语法扩展，允许在 JS 中编写类似 HTML 的代码

3. **Props**：向组件传递数据的方式

4. **State**：组件的内部状态

5. **Hooks**：在函数组件中使用状态和其他 React 特性的 API

### 函数组件示例

```jsx
import React from 'react';

function Greeting(props) {
  return <h1>你好，{props.name}！</h1>;
}

export default Greeting;
```

### 使用 State Hook

```jsx
import React, { useState } from 'react';

function Counter() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>你点击了 {count} 次</p>
      <button onClick={() => setCount(count + 1)}>
        点击我
      </button>
    </div>
  );
}
```

### 使用 Effect Hook

```jsx
import React, { useState, useEffect } from 'react';

function Timer() {
  const [seconds, setSeconds] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setSeconds(seconds => seconds + 1);
    }, 1000);
    
    return () => clearInterval(interval);
  }, []);

  return <div>计时器：{seconds} 秒</div>;
}
```

### 包管理与依赖

React 项目使用 pnpm 管理依赖（pnpm比npm和yarn更快、更高效）：

```bash
# 安装依赖
pnpm add react-router-dom

# 安装开发依赖
pnpm add -D typescript @types/react

# 更新所有依赖
pnpm update

# 运行脚本
pnpm run dev
```

pnpm的主要优势：
- **磁盘空间高效**：pnpm使用内容寻址存储来避免重复安装
- **快速安装**：比npm和yarn快2-3倍
- **严格的依赖管理**：更好的避免依赖地狱问题
- **支持monorepo**：内置对工作空间的支持

常用的包：

- **react-router-dom**: 路由管理
- **axios**: HTTP 请求
- **zustand** 或 **redux-toolkit**: 状态管理
- **styled-components** 或 **emotion**: CSS-in-JS 解决方案
- **MUI** 或 **Ant Design**: UI 组件库

## 本目录代码示例说明

本目录包含两个主要文件：

1. **App.jsx**: 包含三个示例组件
   - 计数器：展示基本的状态管理
   - 计时器：展示 useEffect 的使用
   - 待办事项列表：展示更复杂的状态管理

2. **App.css**: 为组件提供样式

### 如何运行示例代码

要运行本示例，需要将这些文件集成到一个 React 项目中：

1. 创建新的 React 应用：
   ```bash
   pnpm create vite react-demo --template react
   cd react-demo
   pnpm install
   ```

2. 替换 `src/App.jsx` 和 `src/App.css` 为本目录中的文件

3. 启动应用：
   ```bash
   pnpm dev
   ```

## pnpm 常用命令参考

```bash
# 初始化新项目
pnpm init

# 安装依赖
pnpm add [package]

# 安装开发依赖
pnpm add -D [package]

# 全局安装
pnpm add -g [package]

# 运行脚本
pnpm [script]

# 移除依赖
pnpm remove [package]

# 更新依赖
pnpm update

# 查看过时依赖
pnpm outdated
```

## 学习资源

- [React 官方文档](https://reactjs.org/docs/getting-started.html)
- [React Hooks 文档](https://reactjs.org/docs/hooks-intro.html)
- [pnpm 官方文档](https://pnpm.io/zh/)
- [Vite 官方文档](https://vitejs.dev/guide/)
- [React Router 文档](https://reactrouter.com/web/guides/quick-start)

## 练习建议

1. 修改计数器组件，添加最大值和最小值限制
2. 为待办事项添加优先级功能
3. 添加一个新的表单组件，练习表单处理
4. 尝试使用 Context API 在组件之间共享状态 