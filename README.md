# Discord Notify API

一个轻量级的 Go + Gin API 服务，通过 HTTP POST 请求向指定 Discord 频道发送格式化消息。部署后，只需一条 curl 命令即可触发 Discord 消息推送。

## 功能

- 向指定 Discord 频道发送格式化消息
- 支持标题、内容、来源标签
- 消息自动格式化（Markdown 引用块、粗体标题、时间戳）
- API Key 鉴权（sk- 前缀）

## 配置

复制 `.env.example` 为 `.env`，填入实际值：

```
PORT=8080
ENV=production
API_KEY=sk-your-secret-key-here
DISCORD_BOT_TOKEN=your-discord-bot-token
```

## 运行

```bash
go mod tidy
make dev
```

## API

### 发送消息

```
POST /api/v1/send
```

**请求头：**

| 名称 | 必填 | 说明 |
|------|------|------|
| Authorization | 是 | API Key，格式：`Bearer sk-xxx` 或 `sk-xxx` |
| Content-Type | 是 | `application/json` |

**请求体：**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| channel_id | string | 是 | Discord 频道 ID |
| title | string | 是 | 消息标题 |
| content | string | 是 | 消息内容，支持多行 |
| source | string | 否 | 来源标签，如 monitoring、ci/cd、cron |

**响应：**

```json
{
  "success": true,
  "message": "Message sent successfully"
}
```

### 健康检查

```
GET /health
```

无需鉴权，返回 `{"status": "ok"}`。

## curl 调用示例

```bash
# Discord Notify API 调用示例
# --------------------------------------------------
# 接口地址: POST /api/v1/send
# 请求头:
#   Authorization: Bearer sk-xxx (必填, API Key 鉴权)
#   Content-Type: application/json (必填)
# 请求体参数:
#   channel_id (string, 必填) - Discord 频道 ID
#   title      (string, 必填) - 消息标题
#   content    (string, 必填) - 消息内容, 支持多行(\n换行)
#   source     (string, 可选) - 来源标签, 如 monitoring / ci/cd / cron
# --------------------------------------------------

curl -X POST https://your-server.com/api/v1/send \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer sk-your-secret-key-here" \
  -d '{
    "channel_id": "123456789012345678",
    "title": "部署通知",
    "content": "服务 user-api 部署完成\n版本: v2.3.1\n区域: us-east-1",
    "source": "ci/cd"
  }'
```

## Discord 消息效果

![demo](./docs/assets/demo.png)

## License

MIT