package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"jd-go-stock/backend/data"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/mark3labs/mcp-go/mcp"
)

// HandleGetStockKLine 获取K线数据
func HandleGetStockKLine(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	stockCode, err := request.RequireString("stockCode")
	if err != nil {
		return mcp.NewToolResultError("stockCode is required"), nil
	}

	stockCode = GetStockCode(stockCode)

	// 获取天数参数，默认90天
	days := 90
	if daysVal, err := request.RequireFloat("days"); err == nil {
		days = int(daysVal)
	}

	if !strutil.HasPrefixAny(stockCode, []string{"sz", "sh", "hk", "us", "gb_"}) {
		return mcp.NewToolResultError("无效的股票代码。A股:sh/sz开头; 港股:hk开头; 美股:us开头"), nil
	}

	var K *[]data.KLineData
	if strutil.HasPrefixAny(stockCode, []string{"sz", "sh"}) {
		K = data.NewStockDataApi().GetKLineData(stockCode, "240", int64(days))
	} else if strutil.HasPrefixAny(stockCode, []string{"hk", "us", "gb_"}) {
		K = data.NewStockDataApi().GetHK_KLineData(stockCode, "day", int64(days))
	}

	if K == nil || len(*K) == 0 {
		return mcp.NewToolResultError("无K线数据"), nil
	}

	// 构建 Markdown 表格
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("### %s %d日K线数据\n\n", stockCode, days))
	sb.WriteString("| 日期 | 开盘价 | 最高价 | 最低价 | 收盘价 | 成交量(万手) |\n")
	sb.WriteString("|------|--------|--------|--------|--------|-------------|\n")

	for _, kline := range *K {
		volume, _ := convertor.ToFloat(kline.Volume)
		volumeWan := volume / 10000.00 / 100.00
		sb.WriteString(fmt.Sprintf("| %s | %.2f | %.2f | %.2f | %.2f | %.2f |\n",
			kline.Day, kline.Open, kline.High, kline.Low, kline.Close, volumeWan))
	}

	return mcp.NewToolResultText(sb.String()), nil
}

// JSONToMarkdownTable 将 JSON 数组转换为 Markdown 表格
func JSONToMarkdownTable(jsonData []byte) (string, error) {
	var items []map[string]any
	if err := json.Unmarshal(jsonData, &items); err != nil {
		return "", err
	}

	if len(items) == 0 {
		return "No data", nil
	}

	// 获取所有列名
	var headers []string
	for key := range items[0] {
		headers = append(headers, key)
	}

	var sb strings.Builder
	// 表头
	sb.WriteString("| " + strings.Join(headers, " | ") + " |\n")
	// 分隔线
	sb.WriteString("|" + strings.Repeat("------|", len(headers)) + "\n")
	// 数据行
	for _, item := range items {
		var row []string
		for _, h := range headers {
			row = append(row, fmt.Sprintf("%v", item[h]))
		}
		sb.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}

	return sb.String(), nil
}
