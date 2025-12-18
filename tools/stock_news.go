package tools

import (
	"context"
	"strings"

	"mcp-go-stock/backend/data"

	"github.com/mark3labs/mcp-go/mcp"
)

// HandleGetStockNews 获取个股新闻
func HandleGetStockNews(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	stockCode, err := request.RequireString("stockCode")
	if err != nil {
		return mcp.NewToolResultError("stockCode is required"), nil
	}

	stockCode = GetStockCode(stockCode)

	// 获取股票新闻
	// 注意：这里使用的是 go-stock 中的新闻获取逻辑
	// 实际实现可能需要根据股票代码过滤新闻

	newsApi := data.NewMarketNewsApi()
	news := newsApi.GetNewsList(stockCode, 50)

	if news == nil || len(*news) == 0 {
		return mcp.NewToolResultText("暂无该股票相关新闻"), nil
	}

	var sb strings.Builder
	sb.WriteString("## " + stockCode + " 相关新闻\n\n")

	count := 0
	for _, telegraph := range *news {
		if count >= 15 {
			break
		}
		sb.WriteString("### " + telegraph.Time + "\n")
		sb.WriteString(telegraph.Content + "\n\n")
		count++
	}

	return mcp.NewToolResultText(sb.String()), nil
}
