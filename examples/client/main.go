package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server with capabilities
	mcpServer := server.NewMCPServer(
		"test-server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithToolCapabilities(true),
	)

	// Add a test tool
	mcpServer.AddTool(mcp.NewTool(
		"test-tool",
		mcp.WithDescription("Test tool"),
		mcp.WithString("parameter-1", mcp.Description("A string tool parameter")),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "Test Tool Annotation Title",
			ReadOnlyHint:    true,
			DestructiveHint: false,
			IdempotentHint:  true,
			OpenWorldHint:   false,
		}),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Input parameter: " + request.Params.Arguments["parameter-1"].(string),
				},
			},
		}, nil
	})

	// Initialize
	c, err := client.NewSSEMCPClient("http://localhost:8081/sse")
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		return
	}
	defer c.Close()

	sseTransport := c.GetTransport().(*transport.SSE)
	if sseTransport.GetBaseURL() == nil {
		fmt.Println("Base URL should not be nil")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := c.Start(ctx); err != nil {
		fmt.Printf("Failed to start client: %v", err)
		return
	}

	// Initialize
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "test-client",
		Version: "1.0.0",
	}

	result, err := c.Initialize(ctx, initRequest)
	if err != nil {
		fmt.Printf("Failed to initialize: %v", err)
		return
	}
	if result.ServerInfo.Name != "OpenAPI MCP Server" {
		fmt.Printf(
			"Expected server name 'OpenAPI MCP Server', got '%s'",
			result.ServerInfo.Name,
		)
		return
	}
	// Test Ping
	if err := c.Ping(ctx); err != nil {
		fmt.Printf("Ping failed: %v", err)
		return
	}

	// Test ListTools
	toolsRequest := mcp.ListToolsRequest{}
	toolListResult, err := c.ListTools(ctx, toolsRequest)
	if err != nil {
		fmt.Printf("1.ListTools failed: %v", err)
		return
	}
	if toolListResult == nil || len((*toolListResult).Tools) == 0 {
		fmt.Printf("1.Expected one tool")
		return
	}
	fmt.Printf("1.Tool list result: %v\n", toolListResult.Tools)

	request2 := mcp.CallToolRequest{}
	request2.Params.Name = "_users_post"
	request2.Params.Arguments = map[string]interface{}{
		"name":  fmt.Sprintf("yourname%d", time.Now().Unix()),
		"email": "hhhaa@qq.com",
	}
	result2, err := c.CallTool(ctx, request2)
	if err != nil {
		fmt.Printf("2.CallTool failed: %v", err)
		return
	}
	if len(result2.Content) != 1 {
		fmt.Printf("2.Expected 1 content item, got %d", len(result2.Content))
		return
	}
	fmt.Printf("2.Tool call %s result: %s\n", request2.Params.Name, result2.Content[0].(mcp.TextContent).Text)

	request3 := mcp.CallToolRequest{}
	request3.Params.Name = "_users_get"
	result3, err := c.CallTool(ctx, request3)
	if err != nil {
		fmt.Printf("3.CallTool failed: %v", err)
		return
	}
	if len(result3.Content) != 1 {
		fmt.Printf("3.Expected 1 content item, got %d", len(result3.Content))
		return
	}
	fmt.Printf("3.Tool call %s result: %s\n", request3.Params.Name, result3.Content[0].(mcp.TextContent).Text)
}
