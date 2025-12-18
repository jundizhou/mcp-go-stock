package tools

import (
	"context"
	"strings"

	"mcp-go-stock/backend/data"
	"mcp-go-stock/backend/util"

	"github.com/mark3labs/mcp-go/mcp"
)

// HandleGetEconomicData 获取宏观经济数据
func HandleGetEconomicData(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	dataType := "all"
	if t, err := request.RequireString("type"); err == nil && t != "" {
		dataType = t
	}

	var market strings.Builder
	newsApi := data.NewMarketNewsApi()

	switch dataType {
	case "GDP":
		res := newsApi.GetGDP()
		if res.GDPResult.Data != nil {
			md := util.MarkdownTableWithTitle("国内生产总值(GDP)", res.GDPResult.Data)
			market.WriteString(md)
		}
	case "CPI":
		res := newsApi.GetCPI()
		if res.CPIResult.Data != nil {
			md := util.MarkdownTableWithTitle("居民消费价格指数(CPI)", res.CPIResult.Data)
			market.WriteString(md)
		}
	case "PPI":
		res := newsApi.GetPPI()
		if res.PPIResult.Data != nil {
			md := util.MarkdownTableWithTitle("工业品出厂价格指数(PPI)", res.PPIResult.Data)
			market.WriteString(md)
		}
	case "PMI":
		res := newsApi.GetPMI()
		if res.PMIResult.Data != nil {
			md := util.MarkdownTableWithTitle("采购经理人指数(PMI)", res.PMIResult.Data)
			market.WriteString(md)
		}
	default: // all
		res := newsApi.GetGDP()
		if res.GDPResult.Data != nil {
			md := util.MarkdownTableWithTitle("国内生产总值(GDP)", res.GDPResult.Data)
			market.WriteString(md)
			market.WriteString("\n")
		}

		res2 := newsApi.GetCPI()
		if res2.CPIResult.Data != nil {
			md2 := util.MarkdownTableWithTitle("居民消费价格指数(CPI)", res2.CPIResult.Data)
			market.WriteString(md2)
			market.WriteString("\n")
		}

		res3 := newsApi.GetPPI()
		if res3.PPIResult.Data != nil {
			md3 := util.MarkdownTableWithTitle("工业品出厂价格指数(PPI)", res3.PPIResult.Data)
			market.WriteString(md3)
			market.WriteString("\n")
		}

		res4 := newsApi.GetPMI()
		if res4.PMIResult.Data != nil {
			md4 := util.MarkdownTableWithTitle("采购经理人指数(PMI)", res4.PMIResult.Data)
			market.WriteString(md4)
		}
	}

	if market.Len() == 0 {
		return mcp.NewToolResultText("暂无宏观经济数据"), nil
	}

	return mcp.NewToolResultText(market.String()), nil
}
