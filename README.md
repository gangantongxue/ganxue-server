# Ganxue Server - Golang学习网页后端项目

## 项目概述
Ganxue Server 是一个专为Golang学习打造的网页后端项目，集成了丰富的功能，如用户认证、代码运行和文档管理等。该项目采用Go语言开发，结合了多种数据库（MySQL、Redis、MongoDB），使用Hertz作为HTTP框架，具备完善的日志记录和配置管理机制，为Golang学习提供了一个稳定且功能强大的后端支持。

## 网页地址
[敢学](https://ganxue.com)

## 前端项目
[Ganxue Web](https://github.com/gangantongxue/ganxue-web)

## 功能特性
### 用户认证
- 支持用户注册、登录、重置密码和忘记密码等功能。
- 使用JWT进行身份验证，保障用户信息安全。

### 代码运行
- 提供代码运行接口，用户可以提交代码并获取运行结果。
- 采用Seccomp进行系统调用过滤，增强代码运行的安全性。

### 文档管理
- 允许用户根据ID获取Golang学习相关的文档内容，支持Markdown格式。

### 日志记录
- 对不同级别的错误和信息进行详细的日志记录，方便后续问题排查和系统监控。

### 配置管理
- 通过YAML文件和环境变量进行配置管理，支持动态加载，方便项目部署和维护。

## 项目结构
```plaintext
ganxue-server/
├── config/              # 配置文件相关
│   ├── config.go        # 配置结构体定义
│   ├── mysql.go         # MySQL 配置
│   ├── redis.go         # Redis 配置
│   ├── mongodb.go       # MongoDB 配置
│   └── config.yaml      # 配置文件
├── docs/                # 文档相关
│   └── openapi/         # OpenAPI 规范文档
│       ├── sign_in.yaml
│       ├── sign_up.yaml
│       ├── reset_password.yaml
│       ├── run_code.yaml
│       └── ...
├── handler/             # 处理函数
│   ├── open/            # 开放接口处理
│   │   ├── sign_in.go   # 登录处理
│   │   ├── sign_up.go   # 注册处理
│   │   ├── refresh.go   # 刷新 token 处理
│   │   └── ...
│   └── auth/            # 认证接口处理
│       └── detailed_user_info.go
├── initialize/          # 初始化相关
│   ├── init_all.go      # 初始化所有组件
│   └── config.go        # 加载配置文件
├── middleware/          # 中间件
│   └── logMid.go        # 日志中间件
├── model/               # 数据模型
│   ├── answer_model/    # 答案模型
│   │   └── answer_model.go
│   └── md_model/        # Markdown 模型
│       └── md_model.go
├── utils/               # 工具函数
│   ├── db/              # 数据库操作
│   │   ├── init.go      # 数据库初始化
│   │   ├── mysql/       # MySQL 操作
│   │   ├── redis/       # Redis 操作
│   │   └── mongodb/     # MongoDB 操作
│   ├── log/             # 日志记录
│   │   └── log.go
│   ├── token/           # token 处理
│   │   └── token.go
│   ├── password/        # 密码处理
│   └── ver_code/        # 验证码处理
├── global/              # 全局变量
│   └── global.go
├── static/              # 静态文件
│   └── golang/          # Golang 学习资源
│       ├── 0000/
│       │   └── 0000.md
│       ├── 0006/
│       │   └── 0006.txt
│       ├── 0007/
│       │   └── 0007.txt
│       └── ...
├── .gitignore           # Git 忽略文件
├── go.mod               # Go 模块文件
├── go.sum               # Go 模块依赖文件
└── README.md            # 项目说明文档
```

## 环境要求
- **Go**：1.24.1 及以上版本
- **MySQL**：8.0 及以上版本
- **Redis**：6.0 及以上版本
- **MongoDB**：4.4 及以上版本

## 安装与配置
### 克隆项目
```sh
git clone https://github.com/your-repo/ganxue-server.git
cd ganxue-server
```

### 配置环境变量
在项目根目录下创建 `.env` 文件，并添加以下内容：
```plaintext
MYSQL_PASSWORD=your_mysql_password
JWT_SECRET=your_jwt_secret
```

### 配置文件
修改 `config.yaml` 文件，根据实际情况配置数据库和邮件服务信息。

### 安装依赖
```sh
go mod tidy
```

## 运行项目
```sh
go run main.go
```

## 日志记录
项目使用自定义的日志记录工具，支持不同级别的日志记录，如 `Debug`、`Error`、`Panic` 等。日志信息会存储在 Redis 中，方便后续分析。

## 贡献指南
如果你想为项目做出贡献，请遵循以下步骤：
1. Fork 项目
2. 创建新的分支
3. 提交代码并推送至你的分支
4. 发起 Pull Request

## 许可证
本项目采用 [MIT 许可证](LICENSE)。

## 联系信息
如果你有任何问题或建议，请联系 [gangantongxue@outlook.com](mailto:gangantongxue@outlook.com)。