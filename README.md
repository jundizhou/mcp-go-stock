# jd-go-stock MCP Server

基于 go-stock 项目的 MCP (Model Context Protocol) Server，提供股票分析相关的 AI 工具能力。

## 功能特性

- 🔌 支持多种传输协议：stdio、SSE、streamableHttp
- 📈 实时股票价格查询（A股、港股、美股）
- 📊 K线数据获取
- 📰 市场资讯和财经新闻
- 💹 宏观经济数据（GDP、CPI、PPI、PMI）
- 🔍 股票搜索和代码查询
- 📋 财务报告和行业研报
- 🎯 指标选股功能

## 安装

```bash
# 进入目录
cd jd-go-stock

# 下载依赖
go mod tidy

# 编译
go build -o jd-go-stock
```

## 使用方式

### 本地 stdio 模式

```bash
./jd-go-stock
```

### SSE 模式（远程部署）

```bash
./jd-go-stock --mode=sse --port=3000
```

### HTTP Stream 模式（远程部署）

```bash
./jd-go-stock --mode=http --port=3000
```

## 客户端配置

### Claude Desktop (stdio 模式)

编辑 `~/Library/Application Support/Claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "jd-go-stock": {
      "command": "/path/to/jd-go-stock"
    }
  }
}
```

### Claude Desktop (SSE 远程模式)

```json
{
  "mcpServers": {
    "jd-go-stock": {
      "type": "sse",
      "url": "http://your-server.com:3000"
    }
  }
}
```

### Claude Desktop (streamableHttp 远程模式)

```json
{
  "mcpServers": {
    "jd-go-stock": {
      "type": "streamableHttp",
      "url": "http://your-server.com:3000/mcp"
    }
  }
}
```

## 可用工具

| 工具名称 | 描述 | 参数 |
|---------|------|------|
| `get_stock_price` | 获取实时股价 | `stockCodes`: 股票代码，逗号分隔 |
| `get_stock_kline` | 获取K线数据 | `stockCode`: 股票代码, `days`: 天数(可选) |
| `get_market_news` | 获取市场资讯 | 无 |
| `get_economic_data` | 获取宏观经济数据 | `type`: GDP/CPI/PPI/PMI/all |
| `search_stock` | 搜索股票 | `keyword`: 搜索关键词 |
| `get_stock_code` | 获取股票代码 | `stockName`: 股票名称 |
| `get_stock_news` | 获取个股新闻 | `stockCode`: 股票代码 |
| `get_financial_report` | 获取财务报告 | `stockCode`: 股票代码 |
| `get_industry_research` | 获取行业研报 | `industry`: 行业名称(可选) |
| `choice_stock_by_indicators` | 指标选股 | `indicators`: 筛选条件 |

## Docker 部署

```bash
# 构建镜像
docker build -t jd-go-stock .

# 运行 SSE 模式
docker run -p 3000:3000 jd-go-stock

# 运行 HTTP 模式
docker run -p 3000:3000 jd-go-stock ./jd-go-stock --mode=http --port=3000
```

## 示例使用

在 Claude 中可以这样使用：

- "帮我查询贵州茅台的实时股价"
- "获取腾讯控股最近30天的K线数据"
- "今天有什么重要的市场资讯？"
- "查询最新的GDP数据"
- "搜索新能源相关的股票"
- "筛选PE小于30、ROE大于15%的股票"

## 注意事项

1. 部分功能依赖网络请求，需要确保网络可用
2. 新闻爬虫功能可能需要 Chrome/Chromium 支持
3. 远程部署时建议配合 HTTPS 使用

## License

Apache License 2.0
