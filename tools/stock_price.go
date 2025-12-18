package tools

import (
	"context"
	"encoding/json"
	"strings"

	"jd-go-stock/backend/data"

	"github.com/mark3labs/mcp-go/mcp"
)

// HandleGetStockPrice 获取实时股价
func HandleGetStockPrice(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	stockCodes, err := request.RequireString("stockCodes")
	if err != nil {
		return mcp.NewToolResultError("stockCodes is required"), nil
	}

	codes := strings.Split(stockCodes, ",")
	var processedCodes []string
	for _, code := range codes {
		processedCodes = append(processedCodes, GetStockCode(strings.TrimSpace(code)))
	}

	realTimeData, err := data.NewStockDataApi().GetStockCodeRealTimeData(processedCodes...)
	if err != nil {
		return mcp.NewToolResultError("Failed to get stock data: " + err.Error()), nil
	}

	jsonData, err := json.MarshalIndent(realTimeData, "", "  ")
	if err != nil {
		return mcp.NewToolResultError("Failed to marshal data: " + err.Error()), nil
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}

// GetStockCode 处理股票代码，确保有正确的市场前缀
func GetStockCode(code string) string {
	code = strings.TrimSpace(code)
	code = strings.ToLower(code)

	// 已经有前缀的直接返回
	if strings.HasPrefix(code, "sh") || strings.HasPrefix(code, "sz") ||
		strings.HasPrefix(code, "hk") || strings.HasPrefix(code, "us") ||
		strings.HasPrefix(code, "gb_") {
		return code
	}

	// 根据代码规则添加前缀
	if len(code) == 6 {
		// A股
		if strings.HasPrefix(code, "6") {
			return "sh" + code
		}
		if strings.HasPrefix(code, "0") || strings.HasPrefix(code, "3") {
			return "sz" + code
		}
	}

	// 港股 (5位数字)
	if len(code) == 5 && isNumeric(code) {
		return "hk" + code
	}

	return code
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}
