# go-rest-api

学习 Go 后端开发时写的 REST API 练习：Gin + SQLite + JWT，实现用户注册登录、活动增删改查，以及活动报名与取消报名。

仓库来自 2026-03 ~ 2026-04 的学习阶段，属于练习项目，不是线上服务。

## 技术栈

| 组件 | 说明 |
| --- | --- |
| Web 框架 | Gin |
| 数据库 | SQLite，仓库根目录的 `api.db` |
| 认证 | JWT，`utils/jwt.go` 签发，`middlewares/auth.go` 校验 |
| 密码 | `utils/hash.go` 处理 |
| Go 版本 | `go.mod` 声明 `go 1.26.1` |

## 目录结构

```
go-rest-api/
├── main.go            入口：初始化数据库、注册路由、监听 :8080
├── db/db.go           数据库连接与初始化
├── models/
│   ├── user.go        用户模型
│   └── event.go       活动与报名模型
├── routes/
│   ├── routes.go      路由注册
│   ├── users.go       注册 / 登录
│   ├── events.go      活动 CRUD
│   └── register.go    报名与取消报名
├── middlewares/auth.go  JWT 鉴权中间件
├── utils/
│   ├── jwt.go         JWT 生成与解析
│   └── hash.go        密码哈希
└── api-test/          以 .http 文件保存的接口调试用例，一个文件对应一个接口
```

## 接口

| 方法 | 路径 | 需要鉴权 | 用途 |
| --- | --- | --- | --- |
| POST | `/signup` | 否 | 注册 |
| POST | `/login` | 否 | 登录，返回 JWT |
| GET | `/events` | 否 | 活动列表 |
| GET | `/events/:id` | 否 | 活动详情 |
| POST | `/events` | 是 | 创建活动 |
| PUT | `/events/:id` | 是 | 修改活动 |
| DELETE | `/events/:id` | 是 | 删除活动 |
| POST | `/events/:id/register` | 是 | 报名 |
| DELETE | `/events/:id/register` | 是 | 取消报名 |

鉴权接口需要在请求头携带 `Authorization: Bearer <token>`，路由定义见 `routes/routes.go`。

## 本地运行

```bash
go mod download
go run .
```

服务默认监听 `:8080`。

## 已知情况

- `api.db` 是 SQLite 数据库文件，目前直接提交在仓库里；数据库内容会随本地调试变化，建议改为启动时生成并加入 `.gitignore`
- `utils/jwt.go` 中的签名密钥是硬编码常量，建议改为从环境变量读取
- `go.mod` 里包括 Gin 在内的所有依赖都标记为 `indirect`，建议后续执行 `go mod tidy` 修正
- 监听端口、数据库路径、密钥都没有配置文件或环境变量入口
- 没有测试与 CI

## 说明

本仓库作为 Go 后端入门练习保留。
