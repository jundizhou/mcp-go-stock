package main

import (
	"context"
	"fmt"
	"mcp-go-stock/backend/db"
	"mcp-go-stock/tools"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	os.MkdirAll("data", 0755)
	db.Init("data/stock.db")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]any{"stockCode": "SZ300058"},
		},
	}

	fmt.Println("Testing get_financial_report with 1s timeout...")
	result, err := tools.HandleGetFinancialReport(ctx, req)
	if err != nil {
		fmt.Printf("Expected error caught: %v\n", err)
		return
	}
	if result.IsError {
		fmt.Printf("Expected tool error: %v\n", result.Content)
	} else {
		fmt.Printf("Unexpected success! Result length: %d\n", len(fmt.Sprint(result.Content)))
		fmt.Printf("Result text: %s\n", fmt.Sprint(result.Content))
	}
}
