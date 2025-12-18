package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"jd-go-stock/backend/data"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/random"
	"github.com/mark3labs/mcp-go/mcp"
)

// HandleChoiceStockByIndicators 指标选股
func HandleChoiceStockByIndicators(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	indicators, err := request.RequireString("indicators")
	if err != nil {
		return mcp.NewToolResultError("indicators is required"), nil
	}

	indicators = strings.TrimSpace(indicators)
	if indicators == "" {
		return mcp.NewToolResultError("indicators cannot be empty"), nil
	}

	res := data.NewSearchStockApi(indicators).SearchStock(random.RandInt(5, 20))

	if convertor.ToString(res["code"]) != "100" {
		return mcp.NewToolResultText("未找到符合条件的股票"), nil
	}

	resData, ok := res["data"].(map[string]any)
	if !ok {
		return mcp.NewToolResultText("数据格式错误"), nil
	}

	result, ok := resData["result"].(map[string]any)
	if !ok {
		return mcp.NewToolResultText("结果格式错误"), nil
	}

	dataList, ok := result["dataList"].([]any)
	if !ok || len(dataList) == 0 {
		return mcp.NewToolResultText("无符合条件的股票"), nil
	}

	columns, ok := result["columns"].([]any)
	if !ok {
		return mcp.NewToolResultText("列信息格式错误"), nil
	}

	// 构建表头映射
	headers := map[string]string{}
	for _, v := range columns {
		d := v.(map[string]any)
		title := convertor.ToString(d["title"])
		if convertor.ToString(d["dateMsg"]) != "" {
			title = title + "[" + convertor.ToString(d["dateMsg"]) + "]"
		}
		if convertor.ToString(d["unit"]) != "" {
			title = title + "(" + convertor.ToString(d["unit"]) + ")"
		}
		headers[d["key"].(string)] = title
	}

	// 构建表格数据
	table := &[]map[string]any{}
	for _, v := range dataList {
		d := v.(map[string]any)
		tmp := map[string]any{}
		for key, title := range headers {
			tmp[title] = convertor.ToString(d[key])
		}
		*table = append(*table, tmp)
	}

	jsonData, _ := json.Marshal(*table)
	markdownTable, err := choiceStockJSONToMarkdownTable(jsonData)
	if err != nil {
		return mcp.NewToolResultError("表格生成失败: " + err.Error()), nil
	}

	content := "## 指标选股结果\n\n" + markdownTable
	return mcp.NewToolResultText(content), nil
}

// choiceStockJSONToMarkdownTable 将JSON数据转换为Markdown表格
func choiceStockJSONToMarkdownTable(jsonData []byte) (string, error) {
	var data []map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		return "无数据", nil
	}

	// 获取表头
	headers := []string{}
	for key := range data[0] {
		headers = append(headers, key)
	}

	// 构建表头行
	headerRow := "|"
	for _, header := range headers {
		headerRow += fmt.Sprintf(" %s |", header)
	}
	headerRow += "\n"

	// 构建分隔行
	separatorRow := "|"
	for range headers {
		separatorRow += " --- |"
	}
	separatorRow += "\n"

	// 构建数据行
	bodyRows := ""
	for _, rowData := range data {
		bodyRow := "|"
		for _, header := range headers {
			value := rowData[header]
			bodyRow += fmt.Sprintf(" %v |", value)
		}
		bodyRows += bodyRow + "\n"
	}

	return headerRow + separatorRow + bodyRows, nil
}
