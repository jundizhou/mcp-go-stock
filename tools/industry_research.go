package tools

import (
	"context"
	"strings"

	"jd-go-stock/backend/data"

	"github.com/duke-git/lancet/v2/strutil"
	"github.com/mark3labs/mcp-go/mcp"
)

// HandleGetIndustryResearch 获取行业研报
func HandleGetIndustryResearch(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	industry := ""
	if ind, err := request.RequireString("industry"); err == nil {
		industry = strings.TrimSpace(ind)
	}

	api := data.NewMarketNewsApi()

	// 如果没有指定行业，返回热门研报
	if industry == "" {
		// 获取热门行业研报
		resp := api.IndustryResearchReport("", 10)
		if len(resp) == 0 {
			return mcp.NewToolResultText("暂无行业研报"), nil
		}

		var md strings.Builder
		md.WriteString("## 热门行业研报\n\n")

		for _, a := range resp {
			d := a.(map[string]any)
			if infoCode, ok := d["infoCode"].(string); ok {
				reportInfo := api.GetIndustryReportInfo(infoCode)
				md.WriteString(reportInfo)
				md.WriteString("\n---\n\n")
			}
		}

		return mcp.NewToolResultText(md.String()), nil
	}

	// 尝试将行业名称转换为代码
	// 清理输入
	code := strutil.ReplaceWithMap(industry, map[string]string{
		"-":   "",
		"_":   "",
		"bk":  "",
		"BK":  "",
		"bk0": "",
		"BK0": "",
	})

	resp := api.IndustryResearchReport(code, 7)
	if len(resp) == 0 {
		return mcp.NewToolResultText("未找到 " + industry + " 相关的行业研报"), nil
	}

	var md strings.Builder
	md.WriteString("## " + industry + " 行业研报\n\n")

	for _, a := range resp {
		d := a.(map[string]any)
		if infoCode, ok := d["infoCode"].(string); ok {
			reportInfo := api.GetIndustryReportInfo(infoCode)
			md.WriteString(reportInfo)
			md.WriteString("\n---\n\n")
		}
	}

	return mcp.NewToolResultText(md.String()), nil
}
