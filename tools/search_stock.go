package tools

import (
	"context"
	"encoding/json"
	"strings"

	"jd-go-stock/backend/data"
	"jd-go-stock/backend/db"

	"github.com/mark3labs/mcp-go/mcp"
)

// HandleSearchStock 搜索股票
func HandleSearchStock(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	keyword, err := request.RequireString("keyword")
	if err != nil {
		return mcp.NewToolResultError("keyword is required"), nil
	}

	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return mcp.NewToolResultError("keyword cannot be empty"), nil
	}

	// 从数据库搜索
	var stocks []data.StockBasic
	result := db.Dao.Model(&data.StockBasic{}).
		Where("name LIKE ? OR symbol LIKE ? OR ts_code LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").
		Limit(20).
		Find(&stocks)

	if result.Error != nil {
		return mcp.NewToolResultError("Search failed: " + result.Error.Error()), nil
	}

	if len(stocks) == 0 {
		return mcp.NewToolResultText("未找到匹配的股票"), nil
	}

	// 构建结果
	var sb strings.Builder
	sb.WriteString("## 搜索结果\n\n")
	sb.WriteString("| 代码 | 名称 | 市场 | 行业 |\n")
	sb.WriteString("|------|------|------|------|\n")

	for _, stock := range stocks {
		sb.WriteString("| " + stock.TsCode + " | " + stock.Name + " | " + stock.Market + " | " + stock.Industry + " |\n")
	}

	return mcp.NewToolResultText(sb.String()), nil
}

// HandleGetStockCode 根据名称获取股票代码
func HandleGetStockCode(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	stockName, err := request.RequireString("stockName")
	if err != nil {
		return mcp.NewToolResultError("stockName is required"), nil
	}

	stockName = strings.TrimSpace(stockName)
	if stockName == "" {
		return mcp.NewToolResultError("stockName cannot be empty"), nil
	}

	// 从数据库查询
	var stock data.StockBasic
	result := db.Dao.Model(&data.StockBasic{}).
		Where("name = ?", stockName).
		First(&stock)

	if result.Error != nil {
		// 尝试模糊匹配
		result = db.Dao.Model(&data.StockBasic{}).
			Where("name LIKE ?", "%"+stockName+"%").
			First(&stock)

		if result.Error != nil {
			return mcp.NewToolResultText("未找到股票: " + stockName), nil
		}
	}

	// 转换为带前缀的代码
	code := stock.TsCode
	prefix := ""
	if strings.HasSuffix(code, ".SH") {
		prefix = "sh"
		code = prefix + strings.TrimSuffix(code, ".SH")
	} else if strings.HasSuffix(code, ".SZ") {
		prefix = "sz"
		code = prefix + strings.TrimSuffix(code, ".SZ")
	} else if strings.HasSuffix(code, ".HK") {
		prefix = "hk"
		code = prefix + strings.TrimSuffix(code, ".HK")
	}

	result_data := map[string]string{
		"name":     stock.Name,
		"code":     code,
		"ts_code":  stock.TsCode,
		"market":   stock.Market,
		"industry": stock.Industry,
	}

	jsonData, _ := json.MarshalIndent(result_data, "", "  ")
	return mcp.NewToolResultText(string(jsonData)), nil
}
