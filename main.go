package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"jd-go-stock/backend/db"
	"jd-go-stock/tools"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	ServerName    = "jd-go-stock"
	ServerVersion = "1.0.0"
)

func main() {
	// 命令行参数
	mode := flag.String("mode", "stdio", "Transport mode: stdio, sse, http")
	port := flag.Int("port", 3000, "HTTP port for SSE/HTTP mode")
	dbPath := flag.String("db", "data/stock.db", "SQLite database path")
	flag.Parse()

	// 初始化数据库
	os.MkdirAll("data", 0755)
	log.Printf("Initializing database: %s\n", *dbPath)
	db.Init(*dbPath)
	log.Println("Database initialized")

	// 创建 MCP Server
	s := server.NewMCPServer(
		ServerName,
		ServerVersion,
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)

	// 注册所有工具
	registerAllTools(s)

	// 根据模式启动服务
	switch *mode {
	case "stdio":
		log.Println("Starting MCP server in stdio mode...")
		if err := server.ServeStdio(s); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	case "sse":
		log.Printf("Starting MCP server in SSE mode on port %d...\n", *port)
		sseServer := server.NewSSEServer(s)
		addr := fmt.Sprintf(":%d", *port)
		log.Printf("SSE server listening on %s\n", addr)
		if err := http.ListenAndServe(addr, sseServer); err != nil {
			log.Fatalf("SSE server error: %v", err)
		}
	case "http":
		log.Printf("Starting MCP server in HTTP stream mode on port %d...\n", *port)
		httpServer := server.NewStreamableHTTPServer(s)
		addr := fmt.Sprintf(":%d", *port)
		log.Printf("HTTP server listening on %s\n", addr)
		if err := http.ListenAndServe(addr, httpServer); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	default:
		log.Fatalf("Unknown mode: %s. Use stdio, sse, or http", *mode)
	}
}

// registerAllTools 注册所有 MCP 工具
func registerAllTools(s *server.MCPServer) {
	// 1. 获取实时股价
	s.AddTool(
		mcp.NewTool("get_stock_price",
			mcp.WithDescription("获取股票实时价格数据。支持A股(sh/sz)、港股(hk)、美股(us)"),
			mcp.WithString("stockCodes",
				mcp.Required(),
				mcp.Description("股票代码，多个用逗号分隔。例如: sz000001,sh600519,hk00700"),
			),
		),
		tools.HandleGetStockPrice,
	)

	// 2. 获取K线数据
	s.AddTool(
		mcp.NewTool("get_stock_kline",
			mcp.WithDescription("获取股票K线数据，返回日K数据"),
			mcp.WithString("stockCode",
				mcp.Required(),
				mcp.Description("股票代码。A股:sh/sz开头; 港股:hk开头; 美股:us开头"),
			),
			mcp.WithNumber("days",
				mcp.Description("K线数据条数，默认90天"),
			),
		),
		tools.HandleGetStockKLine,
	)

	// 3. 获取市场资讯
	s.AddTool(
		mcp.NewTool("get_market_news",
			mcp.WithDescription("获取市场资讯、财经电报、重要事件和会议信息"),
		),
		tools.HandleGetMarketNews,
	)

	// 4. 获取宏观经济数据
	s.AddTool(
		mcp.NewTool("get_economic_data",
			mcp.WithDescription("获取宏观经济数据：GDP、CPI、PPI、PMI"),
			mcp.WithString("type",
				mcp.Description("数据类型: all(全部), GDP, CPI, PPI, PMI。默认all"),
				mcp.Enum("all", "GDP", "CPI", "PPI", "PMI"),
			),
		),
		tools.HandleGetEconomicData,
	)

	// 5. 搜索股票
	s.AddTool(
		mcp.NewTool("search_stock",
			mcp.WithDescription("根据关键词搜索股票，返回匹配的股票列表"),
			mcp.WithString("keyword",
				mcp.Required(),
				mcp.Description("搜索关键词，可以是股票代码或名称"),
			),
		),
		tools.HandleSearchStock,
	)

	// 6. 获取股票代码
	s.AddTool(
		mcp.NewTool("get_stock_code",
			mcp.WithDescription("根据股票名称获取完整的股票代码（带市场前缀）"),
			mcp.WithString("stockName",
				mcp.Required(),
				mcp.Description("股票名称，如：贵州茅台、腾讯控股"),
			),
		),
		tools.HandleGetStockCode,
	)

	// 7. 获取个股新闻
	s.AddTool(
		mcp.NewTool("get_stock_news",
			mcp.WithDescription("获取指定股票的相关新闻资讯"),
			mcp.WithString("stockCode",
				mcp.Required(),
				mcp.Description("股票代码，如: sz000001, sh600519"),
			),
		),
		tools.HandleGetStockNews,
	)

	// 8. 获取财务报告
	s.AddTool(
		mcp.NewTool("get_financial_report",
			mcp.WithDescription("获取上市公司财务报告数据"),
			mcp.WithString("stockCode",
				mcp.Required(),
				mcp.Description("股票代码，如: sz000001, sh600519"),
			),
		),
		tools.HandleGetFinancialReport,
	)

	// 9. 获取行业研报
	s.AddTool(
		mcp.NewTool("get_industry_research",
			mcp.WithDescription("获取行业研究报告"),
			mcp.WithString("industry",
				mcp.Description("行业名称，如：新能源、半导体。不填则返回热门研报"),
			),
		),
		tools.HandleGetIndustryResearch,
	)

	// 10. 指标选股
	s.AddTool(
		mcp.NewTool("choice_stock_by_indicators",
			mcp.WithDescription("根据技术指标筛选股票"),
			mcp.WithString("indicators",
				mcp.Required(),
				mcp.Description("筛选指标条件，如: PE<20,ROE>15"),
			),
		),
		tools.HandleChoiceStockByIndicators,
	)
}

// contextKey 用于在 context 中传递数据
type contextKey string

const (
	ctxKeyStockCode contextKey = "stockCode"
)

func withStockCode(ctx context.Context, code string) context.Context {
	return context.WithValue(ctx, ctxKeyStockCode, code)
}
