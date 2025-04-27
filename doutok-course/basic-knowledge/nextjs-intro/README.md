# Next.js 入门教程

Next.js 是一个基于 React 的轻量级框架，用于构建静态和服务器渲染的应用程序。它提供了丰富的功能，如服务器端渲染、静态网站生成、API 路由、自动代码分割等。本教程基于 Next.js 13+ 的 App Router。

## 1. 项目目录与优雅实践

### Next.js 项目结构 (App Router)

一个典型的 Next.js 项目结构如下：

```
my-nextjs-app/
│
├── app/               # App Router 目录（基于文件约定的路由）
│   ├── layout.tsx     # 根布局组件
│   ├── page.tsx       # 首页 (/)
│   ├── about/         # 关于页面路由
│   │   └── page.tsx   # 关于页面 (/about)
│   ├── blogs/         # 博客路由
│   │   ├── [id]/      # 动态路由
│   │   │   └── page.tsx # 博客文章页面 
│   │   ├── new/       # 创建新博客
│   │   │   └── page.tsx # 创建博客页面
│   │   └── page.tsx   # 博客列表页面 
│   ├── api/           # API 路由
│   │   └── route.ts   # API 处理器 
│   ├── globals.css    # 全局样式
│   └── error.tsx      # 错误处理页面
│
├── components/        # React 组件
│   ├── ui/            # UI 组件
│   └── ClientComponent.tsx # 客户端组件示例
│
├── lib/               # 工具函数和库
│   └── utils.ts
│
├── public/            # 静态资源
│   ├── favicon.ico    
│   └── images/
│
├── .next/             # Next.js 构建输出 (git ignored)
├── node_modules/      # 依赖 (git ignored)
├── package.json       # 项目依赖和脚本
├── pnpm-lock.yaml     # pnpm 锁文件
├── next.config.js     # Next.js 配置
├── tsconfig.json      # TypeScript 配置
└── README.md          # 项目说明
```

### Next.js App Router 优雅实践

1. **文件系统路由约定**
   - `app/page.tsx` → `/` (首页)
   - `app/about/page.tsx` → `/about` (关于页面)
   - `app/blogs/[id]/page.tsx` → `/blogs/:id` (动态路由)
   - 特殊文件：
     - `layout.tsx`: 布局组件
     - `loading.tsx`: 加载状态
     - `error.tsx`: 错误处理
     - `not-found.tsx`: 404页面

2. **数据获取方法**
   - React Server Components 中的直接获取
   - `generateStaticParams`: 静态路径生成
   - `revalidatePath`/`revalidateTag`: 按需重新验证
   - 客户端数据获取: SWR 或 React Query

3. **API 路由**
   - `app/api/*/route.ts` 文件定义 API 端点
   - 使用 `NextResponse` 进行响应处理

4. **布局系统**
   - 嵌套布局
   - 平行路由和拦截路由
   - 模板和分组

5. **渲染策略**
   - 服务器组件（默认）
   - 客户端组件 ('use client')
   - 流式渲染和部分渲染

## 2. 快速启动案例前端

### 安装 Node.js 和 pnpm

首先，确保你已安装 Node.js 和 pnpm：

```bash
# 安装 pnpm (如果尚未安装)
npm install -g pnpm

# 验证安装
pnpm --version
```

### 创建 Next.js 项目

使用 pnpm 创建新的 Next.js 项目：

```bash
# 使用 create-next-app 创建TypeScript项目
pnpm create next-app my-nextjs-app --typescript

# 进入项目目录
cd my-nextjs-app
```

在创建项目过程中，会提示你选择一些选项。请确保选择 "Yes" 当询问是否使用 App Router 时。

### 项目结构设置

默认情况下，`create-next-app` 生成的项目已经包含基本结构。您可以根据需要添加额外的目录。

### 创建首页

首页是访问者首先看到的页面，在 `app/page.tsx` 文件中创建：

```tsx
// app/page.tsx
import Link from 'next/link';

export default function Home() {
  return (
    <div className="container mx-auto px-4 py-8">
      <section className="py-12 text-center">
        <h1 className="text-4xl font-bold mb-4">欢迎来到我的博客</h1>
        <p className="text-lg text-gray-600 mb-8">探索技术、设计和创意的世界</p>
        <Link 
          href="/blog" 
          className="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
        >
          浏览博客文章
        </Link>
      </section>
      
      <section className="py-8">
        <h2 className="text-2xl font-bold mb-6 text-center">最新文章</h2>
        <div className="grid md:grid-cols-3 gap-6">
          <div className="border rounded-lg p-6 hover:shadow-md transition-shadow">
            <h3 className="text-xl font-semibold mb-3">Next.js入门指南</h3>
            <p className="text-gray-600 mb-4">了解如何使用Next.js构建现代Web应用</p>
            <Link href="/blog/1" className="text-blue-600 hover:underline">
              阅读更多 &rarr;
            </Link>
          </div>
          <div className="border rounded-lg p-6 hover:shadow-md transition-shadow">
            <h3 className="text-xl font-semibold mb-3">React服务器组件详解</h3>
            <p className="text-gray-600 mb-4">深入理解React服务器组件的工作原理</p>
            <Link href="/blog/2" className="text-blue-600 hover:underline">
              阅读更多 &rarr;
            </Link>
          </div>
          <div className="border rounded-lg p-6 hover:shadow-md transition-shadow">
            <h3 className="text-xl font-semibold mb-3">TypeScript与Next.js</h3>
            <p className="text-gray-600 mb-4">如何在Next.js项目中充分利用TypeScript</p>
            <Link href="/blog/3" className="text-blue-600 hover:underline">
              阅读更多 &rarr;
            </Link>
          </div>
        </div>
      </section>
      
      <section className="py-8">
        <h2 className="text-2xl font-bold mb-6 text-center">关于客户端组件</h2>
        <div className="border rounded-lg p-6">
          <p className="text-gray-700 mb-4">
            Next.js的App Router架构区分<strong>服务器组件</strong>和<strong>客户端组件</strong>。
            默认情况下，所有组件都是服务器组件，在服务器上渲染并发送到客户端。
          </p>
          <p className="text-gray-700 mb-4">
            当需要使用浏览器API、添加交互性或使用React hooks时，应该使用客户端组件。
            通过在文件顶部添加 <code className="bg-gray-100 px-2 py-1 rounded">'use client'</code> 指令来声明客户端组件。
          </p>
          <div className="bg-gray-50 p-4 rounded-lg">
            <h3 className="font-semibold mb-2">客户端组件示例：</h3>
            <pre className="bg-gray-800 text-white p-4 rounded overflow-x-auto">
              <code>{`'use client'

import { useState, useEffect } from 'react'

export default function ClientComponent() {
  const [count, setCount] = useState(0)
  
  return (
    <div>
      <h3>计数器: {count}</h3>
      <button onClick={() => setCount(count + 1)}>
        增加
      </button>
    </div>
  )
}`}</code>
            </pre>
          </div>
        </div>
      </section>

      <section className="py-8">
        <h2 className="text-2xl font-bold mb-6 text-center">功能演示</h2>
        <div className="grid md:grid-cols-2 gap-6">
          <div className="border rounded-lg p-6 hover:shadow-md transition-shadow">
            <h3 className="text-xl font-semibold mb-3">博客列表</h3>
            <p className="text-gray-600 mb-4">查看使用服务器组件和模拟数据实现的博客列表页面</p>
            <Link href="/blog" className="text-blue-600 hover:underline">
              查看示例 &rarr;
            </Link>
          </div>
          <div className="border rounded-lg p-6 hover:shadow-md transition-shadow">
            <h3 className="text-xl font-semibold mb-3">客户端组件</h3>
            <p className="text-gray-600 mb-4">了解如何在Next.js中使用客户端组件实现交互功能</p>
            <Link href="/client-example" className="text-blue-600 hover:underline">
              查看示例 &rarr;
            </Link>
          </div>
        </div>
      </section>
    </div>
  );
}
```

首页包含了以下元素：
- 欢迎区域，包含标题和指向博客列表的链接
- 最新文章区域，展示最近的博客文章

### 创建页面

在 `app` 目录中创建 `page.tsx` 文件以添加新页面：

```tsx
// app/about/page.tsx
export default function AboutPage() {
  return (
    <div>
      <h1>关于我们</h1>
      <p>这是 Next.js 示例项目的关于页面。</p>
    </div>
  )
}
```

### 创建布局

使用 `layout.tsx` 文件创建布局：

```tsx
// app/layout.tsx
import type { ReactNode } from 'react';

interface RootLayoutProps {
  children: ReactNode;
}

export default function RootLayout({ children }: RootLayoutProps) {
  return (
    <html lang="zh">
      <body>
        <header>
          <nav>
            {/* 导航栏组件 */}
          </nav>
        </header>
        <main>{children}</main>
        <footer>© {new Date().getFullYear()} 我的 Next.js 应用</footer>
      </body>
    </html>
  )
}
```

### 创建动态路由

使用方括号语法创建动态路由：

```tsx
// app/blogs/[id]/page.tsx
import { notFound } from 'next/navigation';
import Link from 'next/link';

// 模拟博客数据
const blogPosts = [
  { 
    id: '1', 
    title: 'Next.js入门指南',
    content: `
      <p>Next.js是一个基于React的强大框架，它提供了许多内置功能，使得构建现代Web应用变得更加简单。</p>
      <h2>主要特性</h2>
      <ul>
        <li>服务器端渲染 (SSR)</li>
        <li>静态站点生成 (SSG)</li>
        <li>API路由</li>
        <li>文件系统路由</li>
        <li>内置CSS和Sass支持</li>
        <li>代码分割和打包优化</li>
      </ul>
      <p>使用Next.js，你可以快速开发出高性能的React应用，无需复杂的配置。</p>
    `,
    author: {
      name: '张三',
      avatar: 'https://randomuser.me/api/portraits/men/1.jpg'
    },
    publishedAt: '2023-05-15',
    tags: ['Next.js', 'React', '前端开发']
  },
  { 
    id: '2', 
    title: 'React服务器组件详解',
    content: `
      <p>React服务器组件是React的一项新特性，它允许开发者创建在服务器上渲染的组件，从而提高性能并减少客户端JavaScript的体积。</p>
      <h2>服务器组件的优势</h2>
      <ol>
        <li>减少客户端JavaScript包大小</li>
        <li>直接访问服务器资源（数据库、文件系统等）</li>
        <li>自动代码分割</li>
        <li>改善首次加载性能</li>
      </ol>
      <p>在Next.js的App Router中，所有组件默认都是服务器组件，除非你显式声明为客户端组件。</p>
    `,
    author: {
      name: '李四',
      avatar: 'https://randomuser.me/api/portraits/women/2.jpg'
    },
    publishedAt: '2023-06-22',
    tags: ['React', '服务器组件', '性能优化']
  },
  { 
    id: '3', 
    title: 'TypeScript与Next.js',
    content: `
      <p>TypeScript是JavaScript的超集，添加了静态类型检查，在Next.js项目中使用TypeScript可以带来诸多好处。</p>
      <h2>TypeScript的优势</h2>
      <ul>
        <li>静态类型检查，减少运行时错误</li>
        <li>更好的IDE支持，包括代码补全和智能提示</li>
        <li>更容易维护的代码库</li>
        <li>自文档化的代码</li>
      </ul>
      <h2>在Next.js中使用TypeScript</h2>
      <p>Next.js原生支持TypeScript，你可以直接创建.tsx或.ts文件，无需额外配置。</p>
      <p>对于页面和API路由，你可以使用TypeScript接口来定义props和请求参数的类型。</p>
    `,
    author: {
      name: '王五',
      avatar: 'https://randomuser.me/api/portraits/men/3.jpg'
    },
    publishedAt: '2023-07-10',
    tags: ['TypeScript', 'Next.js', '类型安全']
  }
];

// 获取博客帖子函数
const getBlogPost = (id: string) => {
  return blogPosts.find(post => post.id === id);
};

// 博客详情页面组件
export default function BlogPostPage({ params }: { params: { id: string } }) {
  const post = getBlogPost(params.id);
  
  // 如果没有找到文章，返回404
  if (!post) {
    notFound();
  }
  
  return (
    <div className="container mx-auto px-4 py-8">
      <article className="max-w-3xl mx-auto">
        <div className="mb-8">
          <Link href="/blog" className="text-blue-600 hover:underline mb-4 inline-block">
            &larr; 返回博客列表
          </Link>
          <h1 className="text-4xl font-bold mb-4">{post.title}</h1>
          
          <div className="flex items-center mb-6">
            <img 
              src={post.author.avatar} 
              alt={post.author.name}
              className="w-10 h-10 rounded-full mr-3"
            />
            <div>
              <p className="font-medium">{post.author.name}</p>
              <p className="text-gray-500 text-sm">发布于 {post.publishedAt}</p>
            </div>
          </div>
          
          <div className="flex flex-wrap gap-2 mb-8">
            {post.tags.map(tag => (
              <span key={tag} className="bg-blue-100 text-blue-800 text-sm px-3 py-1 rounded-full">
                {tag}
              </span>
            ))}
          </div>
        </div>
        
        <div 
          className="prose prose-lg max-w-none"
          dangerouslySetInnerHTML={{ __html: post.content }}
        />
        
        <div className="mt-12 pt-8 border-t border-gray-200">
          <h3 className="text-xl font-bold mb-4">分享这篇文章</h3>
          <div className="flex space-x-4">
            <button className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
              分享到微信
            </button>
            <button className="px-4 py-2 bg-blue-400 text-white rounded hover:bg-blue-500">
              分享到微博
            </button>
          </div>
        </div>
      </article>
    </div>
  );
} 
```

### 创建 API 路由

在 `app/api` 目录中创建 API 端点：

```ts
// app/api/hello/route.ts
// app/api/route.ts
import { NextRequest, NextResponse } from 'next/server'

// 模拟用户数据
const users = [
  { id: '1', name: '张三', email: 'zhangsan@example.com' },
  { id: '2', name: '李四', email: 'lisi@example.com' },
  { id: '3', name: '王五', email: 'wangwu@example.com' }
];

export async function GET() {
  return NextResponse.json({ 
    success: true,
    data: users,
    timestamp: new Date().toISOString()
  })
}

export async function POST(request: NextRequest) {
  try {
    const body = await request.json()
    
    // 验证请求数据
    if (!body.name) {
      return NextResponse.json(
        { success: false, error: '名称不能为空' },
        { status: 400 }
      )
    }
    
    // 模拟创建新用户
    const newUser = {
      id: (users.length + 1).toString(),
      name: body.name,
      email: body.email || null
    }
    
    // 在真实应用中，这里会将用户添加到数据库
    // 这里只是模拟
    
    return NextResponse.json({ 
      success: true, 
      data: newUser 
    }, { status: 201 })
  } catch (error) {
    return NextResponse.json(
      { success: false, error: '请求处理失败' },
      { status: 500 }
    )
  }
}
```

### 数据获取

在服务器组件中进行数据获取（默认情况下，`page.tsx` 是服务器组件）：

```tsx
// app/blog/page.tsx - 服务器组件
import Link from 'next/link';
import BlogActions from './components/BlogActions';

// 模拟博客数据
const blogs = [
  { 
    id: '1', 
    title: 'Next.js入门指南',
    excerpt: '了解如何使用Next.js构建现代Web应用',
    author: {
      name: '张三',
      avatar: 'https://randomuser.me/api/portraits/men/1.jpg'
    },
    publishedAt: '2023-05-15',
    tags: ['Next.js', 'React', '前端开发']
  },
  { 
    id: '2', 
    title: 'React服务器组件详解',
    excerpt: '深入理解React服务器组件的工作原理',
    author: {
      name: '李四',
      avatar: 'https://randomuser.me/api/portraits/women/2.jpg'
    },
    publishedAt: '2023-06-22',
    tags: ['React', '服务器组件', '性能优化']
  },
  { 
    id: '3', 
    title: 'TypeScript与Next.js',
    excerpt: '如何在Next.js项目中充分利用TypeScript',
    author: {
      name: '王五',
      avatar: 'https://randomuser.me/api/portraits/men/3.jpg'
    },
    publishedAt: '2023-07-10',
    tags: ['TypeScript', 'Next.js', '类型安全']
  }
];

// 模拟获取博客列表函数
async function getBlogs() {
  // 模拟网络延迟
  await new Promise(resolve => setTimeout(resolve, 500));
  return blogs;
}

export default async function BlogsPage() {
  // 获取博客数据
  const blogList = await getBlogs();
  
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold">博客列表</h1>
        <Link 
          href="/blog/new" 
          className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition"
        >
          创建新博客
        </Link>
      </div>
      
      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        {blogList.map(blog => (
          <div key={blog.id} className="border rounded-lg overflow-hidden hover:shadow-md transition">
            <div className="p-6">
              <h2 className="text-xl font-bold mb-3">{blog.title}</h2>
              <p className="text-gray-600 mb-4">{blog.excerpt}</p>
              
              <div className="flex flex-wrap gap-2 mb-4">
                {blog.tags.map(tag => (
                  <span 
                    key={tag} 
                    className="bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded"
                  >
                    {tag}
                  </span>
                ))}
              </div>
              
              <div className="flex items-center text-sm text-gray-500 mb-4">
                <img 
                  src={blog.author.avatar} 
                  alt={blog.author.name}
                  className="w-6 h-6 rounded-full mr-2"
                />
                <span>{blog.author.name}</span>
                <span className="mx-2">•</span>
                <span>{new Date(blog.publishedAt).toLocaleDateString('zh-CN')}</span>
              </div>
              
              <div className="flex space-x-2">
                <Link 
                  href={`/blog/${blog.id}`} 
                  className="px-3 py-1 bg-blue-500 text-white rounded hover:bg-blue-600 text-sm"
                >
                  查看全文
                </Link>
                <Link 
                  href={`/blog/edit/${blog.id}`} 
                  className="px-3 py-1 bg-green-500 text-white rounded hover:bg-green-600 text-sm"
                >
                  编辑
                </Link>
                <BlogActions blogId={blog.id} />
              </div>
            </div>
          </div>
        ))}
      </div>
      
      {blogList.length === 0 && (
        <div className="text-center py-12">
          <p className="text-gray-500 mb-4">暂无博客内容</p>
          <Link 
            href="/blog/new" 
            className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
          >
            创建第一篇博客
          </Link>
        </div>
      )}
      
      <div className="mt-8 text-center">
        <Link 
          href="/" 
          className="text-blue-600 hover:underline"
        >
          返回首页
        </Link>
      </div>
    </div>
  );
} 
```

使用客户端组件进行交互：

```tsx
// blog/components/BlogActions.tsx
'use client';

interface BlogActionsProps {
  blogId: string;
}

export default function BlogActions({ blogId }: BlogActionsProps) {
  const handleDelete = () => {
    alert(`删除功能尚未实现：${blogId}`);
    // 这里可以实现实际的删除逻辑
  };

  return (
    <button 
      className="px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600 text-sm"
      onClick={handleDelete}
    >
      删除
    </button>
  );
} 
```

### 安装依赖

使用 pnpm 安装项目依赖：

```bash
# 安装依赖
pnpm add axios swr

# 安装开发依赖
pnpm add -D typescript @types/react eslint
```

### 运行开发服务器

```bash
pnpm dev
```

### 构建和部署

```bash
# 构建应用
pnpm build

# 启动生产环境服务器
pnpm start
```

## pnpm 优势

pnpm相比npm和yarn有以下优势：

1. **磁盘空间效率**：pnpm使用硬链接和内容寻址存储，减少了重复的依赖
2. **安装速度快**：比npm和yarn快2-3倍
3. **更严格的依赖管理**：通过使用符号链接确保依赖访问更安全
4. **内置monorepo支持**：无需额外工具即可管理多包项目

## pnpm 常用命令

```bash
# 初始化项目
pnpm init

# 安装所有依赖
pnpm install

# 添加依赖
pnpm add [package]

# 添加开发依赖
pnpm add -D [package]

# 更新依赖
pnpm update

# 运行脚本
pnpm [script]

# 删除依赖
pnpm remove [package]
```

## Node.js 与 Next.js 的关系

### Node.js 与 Next.js 的基本关系

Next.js 是构建在 Node.js 之上的 React 框架。这种关系可以从多个方面理解：

1. **运行时环境**：Next.js 使用 Node.js 作为其服务器端运行时环境
2. **构建工具**：Next.js 利用 Node.js 生态系统中的工具（如 webpack、babel）进行代码构建和转换
3. **包管理**：Next.js 项目通过 npm、yarn 或 pnpm 等 Node.js 包管理器管理依赖
4. **API 实现**：Next.js 的服务器端 API 路由基于 Node.js 的 HTTP 模块实现

### Next.js 的运行时环境

Next.js 确实在 Node.js 环境中启动了一个服务器来接收来自浏览器的请求。这个过程在不同模式下有所不同：

#### 开发环境 (`pnpm dev`)

在开发模式下：

1. Next.js 启动一个 Node.js HTTP 服务器（默认监听 3000 端口）
2. 该服务器具有热模块替换(HMR)功能，允许实时更新
3. 当浏览器请求到达时，Next.js 服务器根据请求的路径：
   - 对于页面请求：执行服务器端渲染(SSR)或提供静态生成(SSG)的内容
   - 对于 API 请求：执行相应的 API 路由处理函数
   - 对于静态资源：提供 public 目录中的文件
4. 开发服务器还处理源代码编译、打包和监视文件变化

```
浏览器请求 → Node.js服务器(Next.js) → 路由解析 → 页面渲染/API处理 → 响应返回
```

#### 生产环境 (`pnpm build` 然后 `pnpm start`)

在生产模式下：

1. `pnpm build` 预先构建所有可能的页面和资源
   - 静态生成(SSG)的页面被预渲染为HTML
   - 服务器组件被优化和序列化
   - JavaScript包被优化和代码分割
   
2. `pnpm start` 启动一个优化的 Node.js 生产服务器
   - 这个服务器比开发服务器轻量得多
   - 它主要负责：
     - 提供预构建的静态资源
     - 处理动态SSR请求
     - 执行API路由

```
浏览器请求 → Node.js生产服务器 → 提供预构建资源/动态渲染 → 响应返回
```

### 渲染模式与Node.js的关系

1. **服务器端渲染(SSR)**
   - 每次请求都在 Node.js 环境中执行React组件渲染
   - 生成HTML并发送给浏览器
   - 适用于需要最新数据的页面

2. **静态站点生成(SSG)**
   - 在构建时在 Node.js 环境中预渲染HTML
   - 请求来临时直接提供静态HTML
   - 适用于内容不经常变化的页面

3. **增量静态再生成(ISR)**
   - 结合SSG和SSR的优点
   - 预渲染HTML，但在指定间隔后在Node.js环境中重新生成

4. **客户端渲染**
   - 初始HTML由服务器提供
   - 后续渲染和数据获取在浏览器中发生
   - 减轻Node.js服务器负载

### Node.js 环境的限制

在使用Next.js时，需要注意Node.js环境的一些特点和限制：

1. **服务器组件 vs 客户端组件**
   - 服务器组件在Node.js环境中运行，可以访问文件系统、环境变量等
   - 客户端组件无法访问Node.js特有的功能和API

2. **API路由的Node.js能力**
   - API路由在Node.js环境中执行，可以使用完整的Node.js功能
   - 包括数据库连接、文件系统操作、复杂计算等

3. **边缘运行时**
   - Next.js还支持Edge Runtime（一种轻量级运行时）
   - Edge Runtime比Node.js更受限，但部署和冷启动更快

### 部署架构

Next.js应用的部署涉及到Node.js服务器的管理：

1. **传统服务器**
   - 部署完整的Node.js服务器
   - 例如在AWS EC2、DigitalOcean等上运行

2. **无服务器函数**
   - 将Next.js应用部署为无服务器函数
   - 例如AWS Lambda、Vercel等

3. **静态导出**
   - 完全静态导出，不需要Node.js服务器
   - 使用 `next export` 命令
   - 适用于不需要SSR或API路由的项目

## 本目录示例代码说明

本目录包含以下示例文件：

1. **page.tsx**: 一个简单的 Next.js 首页示例，展示了：
   - 页面组件结构
   - 使用 `next/head` 管理头部元素
   - 使用 `next/link` 进行客户端导航
   - React hooks 在 Next.js 中的使用

2. **[id].tsx**: 展示动态路由的实现，包括：
   - 动态路由参数获取
   - 静态生成 (generateStaticParams)
   - 数据获取模式

3. **route.ts**: API 路由示例，展示了：
   - 基于请求方法的处理逻辑
   - 响应处理
   - 错误处理

注意：这些示例基于 App Router 模式实现。

## App Router 和 Pages Router 的区别

| 特性 | App Router (app/) | Pages Router (pages/) |
|------|-------------------|----------------------|
| 组件模型 | React Server Components | 客户端组件 |
| 数据获取 | 组件中的 fetch 函数 | getServerSideProps/getStaticProps |
| 布局 | layout.tsx | _app.tsx 和布局组件 |
| 嵌套布局 | 多个 layout.tsx | 需手动实现 |
| 加载状态 | loading.tsx | 需手动实现 |
| 错误处理 | error.tsx | 需手动实现或使用 Error Boundaries |
| API 路由 | route.ts 处理程序 | pages/api/*.ts |

## 进阶资源

- [Next.js 官方文档](https://nextjs.org/docs)
- [Next.js App Router 文档](https://nextjs.org/docs/app)
- [Next.js 学习课程](https://nextjs.org/learn)
- [pnpm 官方文档](https://pnpm.io/zh/)
- [Vercel 平台](https://vercel.com)（Next.js 的创建者提供的托管服务）
- [Next.js GitHub 仓库](https://github.com/vercel/next.js)

## 练习建议

1. 创建一个包含多个页面的 Next.js 应用（使用 App Router）
2. 实现动态路由和数据获取
3. 添加多级嵌套布局
4. 创建 API 路由
5. 实现错误处理和加载状态
6. 将项目部署到 Vercel 或其他托管平台 
