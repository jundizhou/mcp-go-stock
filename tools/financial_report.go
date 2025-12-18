package tools

import (
	"context"
	"strings"

	"mcp-go-stock/backend/data"

	"github.com/mark3labs/mcp-go/mcp"
)

// HandleGetFinancialReport 获取财务报告
func HandleGetFinancialReport(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	stockCode, err := request.RequireString("stockCode")
	if err != nil {
		return mcp.NewToolResultError("stockCode is required"), nil
	}

	stockCode = GetStockCode(stockCode)

	messages := data.GetFinancialReportsByXUEQIU(stockCode, 30)
	if messages == nil || len(*messages) == 0 {
		return mcp.NewToolResultText("没有找到 " + stockCode + " 的财务报告"), nil
	}

	var md strings.Builder
	md.WriteString("## " + stockCode + " 财务报告\n\n")
	for _, s := range *messages {
		md.WriteString(s)
	}

	return mcp.NewToolResultText(md.String()), nil
}
