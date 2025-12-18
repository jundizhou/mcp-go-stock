package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"mcp-go-stock/backend/db"
	"mcp-go-stock/tools"

	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	// 初始化数据库
	os.MkdirAll("data", 0755)
	db.Init("data/stock.db")

	ctx := context.Background()

	fmt.Println("========================================")
	fmt.Println("测试 jd-go-stock MCP Server 工具")
	fmt.Println("测试股票: 600498 烽火通信")
	fmt.Println("========================================")

	results := make(map[string]string)

	// 1. 测试 get_stock_price
	fmt.Println("\n【1】测试 get_stock_price")
	req := createRequest(map[string]any{"stockCodes": "sh600498"})
	result, err := tools.HandleGetStockPrice(ctx, req)
	results["get_stock_price"] = checkResult(result, err)

	// 2. 测试 get_stock_kline
	fmt.Println("\n【2】测试 get_stock_kline")
	req = createRequest(map[string]any{"stockCode": "sh600498", "days": float64(5)})
	result, err = tools.HandleGetStockKLine(ctx, req)
	results["get_stock_kline"] = checkResult(result, err)

	// 3. 测试 search_stock
	fmt.Println("\n【3】测试 search_stock")
	req = createRequest(map[string]any{"keyword": "烽火"})
	result, err = tools.HandleSearchStock(ctx, req)
	results["search_stock"] = checkResult(result, err)

	// 4. 测试 get_stock_code
	fmt.Println("\n【4】测试 get_stock_code")
	req = createRequest(map[string]any{"stockName": "烽火通信"})
	result, err = tools.HandleGetStockCode(ctx, req)
	results["get_stock_code"] = checkResult(result, err)

	// 5. 测试 get_stock_news
	fmt.Println("\n【5】测试 get_stock_news")
	req = createRequest(map[string]any{"stockCode": "sh600498"})
	result, err = tools.HandleGetStockNews(ctx, req)
	results["get_stock_news"] = checkResult(result, err)

	// 6. 测试 get_financial_report
	fmt.Println("\n【6】测试 get_financial_report")
	req = createRequest(map[string]any{"stockCode": "sh600498"})
	result, err = tools.HandleGetFinancialReport(ctx, req)
	results["get_financial_report"] = checkResult(result, err)

	// 7. 测试 get_market_news
	fmt.Println("\n【7】测试 get_market_news")
	req = createRequest(map[string]any{})
	result, err = tools.HandleGetMarketNews(ctx, req)
	results["get_market_news"] = checkResult(result, err)

	// 8. 测试 get_economic_data
	fmt.Println("\n【8】测试 get_economic_data")
	req = createRequest(map[string]any{"type": "GDP"})
	result, err = tools.HandleGetEconomicData(ctx, req)
	results["get_economic_data"] = checkResult(result, err)

	// 9. 测试 get_industry_research
	fmt.Println("\n【9】测试 get_industry_research")
	req = createRequest(map[string]any{})
	result, err = tools.HandleGetIndustryResearch(ctx, req)
	results["get_industry_research"] = checkResult(result, err)

	// 10. 测试 choice_stock_by_indicators
	fmt.Println("\n【10】测试 choice_stock_by_indicators")
	req = createRequest(map[string]any{"indicators": "市盈率<20"})
	result, err = tools.HandleChoiceStockByIndicators(ctx, req)
	results["choice_stock_by_indicators"] = checkResult(result, err)

	// 打印汇总
	fmt.Println("\n========================================")
	fmt.Println("测试结果汇总")
	fmt.Println("========================================")
	passCount := 0
	for name, status := range results {
		fmt.Printf("%s: %s\n", name, status)
		if status == "✅ 成功" {
			passCount++
		}
	}
	fmt.Printf("\n总计: %d/10 工具测试通过\n", passCount)
}

func createRequest(args map[string]any) mcp.CallToolRequest {
	return mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: args,
		},
	}
}

func checkResult(result *mcp.CallToolResult, err error) string {
	if err != nil {
		fmt.Printf("❌ 错误: %v\n", err)
		return "❌ 错误"
	}
	if result == nil {
		fmt.Println("❌ 结果为空")
		return "❌ 结果为空"
	}
	if result.IsError {
		content, _ := json.Marshal(result.Content)
		fmt.Printf("❌ 工具错误: %s\n", string(content)[:min(200, len(string(content)))])
		return "❌ 工具错误"
	}

	content, _ := json.Marshal(result.Content)
	text := string(content)
	if len(text) > 300 {
		text = text[:300] + "..."
	}
	fmt.Printf("✅ 成功: %s\n", text)
	return "✅ 成功"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
