package tools

import (
	"context"
	"encoding/json"
	"strings"

	"mcp-go-stock/backend/data"

	"github.com/duke-git/lancet/v2/random"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/tidwall/gjson"
)

// HandleGetMarketNews 获取市场资讯
func HandleGetMarketNews(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var md strings.Builder

	// 1. 获取财经日历
	res := data.NewMarketNewsApi().ClsCalendar()
	if len(res) > 0 {
		md.WriteString("## 📅 财经日历\n\n")
		for _, a := range res {
			bytes, err := json.Marshal(a)
			if err != nil {
				continue
			}
			date := gjson.Get(string(bytes), "calendar_day")
			md.WriteString("### 日期：" + date.String() + "\n")
			list := gjson.Get(string(bytes), "items")
			list.ForEach(func(key, value gjson.Result) bool {
				title := gjson.Get(value.String(), "title").String()
				md.WriteString("- " + title + "\n")
				return true
			})
			md.WriteString("\n")
		}
	}

	// 2. 获取市场新闻
	news := data.NewMarketNewsApi().GetNewsList("", random.RandInt(50, 100))
	if news != nil && len(*news) > 0 {
		md.WriteString("## 📰 市场资讯\n\n")
		count := 0
		for _, telegraph := range *news {
			if count >= 20 {
				break
			}
			md.WriteString("### " + telegraph.Time + "\n")
			md.WriteString(telegraph.Content + "\n\n")
			count++
		}
	}

	// 3. 获取 TradingView 全球新闻
	resp := data.NewMarketNewsApi().TradingViewNews()
	if resp != nil && len(*resp) > 0 {
		md.WriteString("## 🌍 全球新闻\n\n")
		count := 0
		for _, a := range *resp {
			if count >= 10 {
				break
			}
			md.WriteString("- " + a.Title + "\n")
			count++
		}
		md.WriteString("\n")
	}

	// 4. 获取路透社新闻
	reutersNew := data.NewMarketNewsApi().ReutersNew()
	if reutersNew.Result.Articles != nil && len(reutersNew.Result.Articles) > 0 {
		md.WriteString("## 📡 路透社资讯\n\n")
		count := 0
		for _, article := range reutersNew.Result.Articles {
			if count >= 10 {
				break
			}
			md.WriteString("### " + article.Title + "\n")
			if article.Description != "" {
				md.WriteString(article.Description + "\n")
			}
			md.WriteString("\n")
			count++
		}
	}

	if md.Len() == 0 {
		return mcp.NewToolResultText("暂无市场资讯"), nil
	}

	return mcp.NewToolResultText(md.String()), nil
}
