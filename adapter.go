package gmadapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// OpenAPIToMCPAdapter 是一个适配器，用于将 OpenAPI 文档中的server转换为 MCP 工具
type OpenAPIToMCPAdapter struct {
	server         *server.MCPServer
	backendBaseUrl string
	addrs          string
	openAPI        map[string]interface{}
	tools          map[string]*mcp.Tool
	handlers       map[string]func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// NewOpenAPIToMCPAdapter 创建一个新的适配器
func NewOpenAPIToMCPAdapter(name, version, backendBaseUrl, myAddr string) (*OpenAPIToMCPAdapter, error) {
	s := server.NewMCPServer(name, version)

	return &OpenAPIToMCPAdapter{
		server:         s,
		backendBaseUrl: backendBaseUrl,
		addrs:          myAddr,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}, nil
}

// LoadOpenAPI 从 URL 或本地文件加载 OpenAPI 文档
func (a *OpenAPIToMCPAdapter) LoadOpenAPI(source string) error {
	var data []byte
	var err error

	if strings.HasPrefix(source, "http") {
		resp, err := http.Get(source)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	} else {
		data, err = ioutil.ReadFile(source)
		if err != nil {
			return err
		}
	}

	if err := yaml.Unmarshal(data, &a.openAPI); err != nil {
		return err
	}

	return nil
}

// GenerateTools 从 OpenAPI 文档生成 MCP 工具
func (a *OpenAPIToMCPAdapter) GenerateTools() error {
	paths, ok := a.openAPI["paths"].(map[string]interface{})
	if !ok {
		return errors.New("failed to extract paths from OpenAPI document")
	}

	for path, pathItem := range paths {
		pathItemMap, ok := pathItem.(map[string]interface{})
		if !ok {
			continue
		}

		for method, operation := range pathItemMap {
			operationMap, ok := operation.(map[string]interface{})
			if !ok {
				continue
			}

			toolName := fmt.Sprintf("%s_%s", strings.ReplaceAll(path, "/", "_"), method)
			toolDesc := ""
			if summary, ok := operationMap["summary"].(string); ok {
				toolDesc = summary
			}
			if toolDesc == "" {
				if desc, ok := operationMap["description"].(string); ok {
					toolDesc = desc
				}
			}

			var toolOpts []mcp.ToolOption
			toolOpts = append(toolOpts, mcp.WithDescription(toolDesc))

			// 处理路径参数和查询参数
			if params, ok := operationMap["parameters"].([]interface{}); ok {
				for _, param := range params {
					paramMap, ok := param.(map[string]interface{})
					if !ok {
						continue
					}

					paramName, _ := paramMap["name"].(string)
					paramDesc, _ := paramMap["description"].(string)
					required, _ := paramMap["required"].(bool)

					schemaMap := paramMap["schema"].(map[string]interface{})
					generator := map[string]interface{}{
						"type":        schemaMap["type"],
						"description": paramDesc,
						"required":    required,
					}
					paramDef, ok := paramMap["default"]
					if ok {
						generator["default"] = paramDef
					}
					paramProps, ok := schemaMap["properties"]
					if ok {
						generator["props"] = paramProps.(map[string]interface{})
					}
					paramItems, ok := schemaMap["items"]
					if ok {
						generator["items"] = paramItems
					}
					opt, err := a.getMCPPropertyOption(paramName, generator)
					if err != nil {
						log.Printf("failed to create property option for %s: %v", paramName, err)
						continue
					}
					toolOpts = append(toolOpts, opt)
				}
			}

			// 处理请求体
			if requestBody, ok := operationMap["requestBody"].(map[string]interface{}); ok {
				content, ok := requestBody["content"].(map[string]interface{})
				if !ok {
					continue
				}

				for _, schema := range content {
					schemaMap, ok := schema.(map[string]interface{})
					if !ok {
						continue
					}

					properties, ok := schemaMap["schema"].(map[string]interface{})["properties"].(map[string]interface{})
					if !ok {
						continue
					}

					for paramName, param := range properties {
						paramMap, ok := param.(map[string]interface{})
						if !ok {
							continue
						}

						paramType, ok := paramMap["type"].(string)
						if !ok {
							continue
						}
						paramDesc, _ := paramMap["description"].(string)

						generator := map[string]interface{}{
							"type":        paramType,
							"description": paramDesc,
						}
						paramDef, ok := paramMap["default"]
						if ok {
							generator["default"] = paramDef
						}
						paramProps, ok := paramMap["properties"]
						if ok {
							generator["props"] = paramProps.(map[string]interface{})
						}
						paramItems, ok := paramMap["items"]
						if ok {
							generator["items"] = paramItems
						}

						opt, err := a.getMCPPropertyOption(paramName, generator)
						if err != nil {
							log.Printf("failed to create property option for %s: %v", paramName, err)
							continue
						}
						toolOpts = append(toolOpts, opt)
					}
				}
			}

			tool := mcp.NewTool(toolName, toolOpts...)
			a.tools[toolName] = &tool
			a.handlers[toolName] = a.createHandler(path, method)

			log.Printf("create a tool for %s", toolName)
		}
	}

	return nil
}

// getMCPPropertyOption 根据参数类型返回对应的 MCP 属性选项
func (a *OpenAPIToMCPAdapter) getMCPPropertyOption(name string, param map[string]interface{}) (mcp.ToolOption, error) {
	paramType, ok := param["type"].(string)
	if !ok {
		return nil, errors.New("failed to extract type from parameter")
	}

	var propOpts []mcp.PropertyOption

	// 解析描述
	if desc, ok := param["description"].(string); ok {
		propOpts = append(propOpts, mcp.Description(desc))
	}

	// 解析是否必填
	if required, ok := param["required"].(bool); ok && required {
		propOpts = append(propOpts, mcp.Required())
	}

	// 解析默认值
	if def, ok := param["default"]; ok {
		switch do := def.(type) {
		case string:
			propOpts = append(propOpts, mcp.DefaultString(do))
		case bool:
			propOpts = append(propOpts, mcp.DefaultBool(do))
		case float64:
			propOpts = append(propOpts, mcp.DefaultNumber(float64(do)))
		case float32:
			propOpts = append(propOpts, mcp.DefaultNumber(float64(do)))
		case int64:
			propOpts = append(propOpts, mcp.DefaultNumber(float64(do)))
		case int:
			propOpts = append(propOpts, mcp.DefaultNumber(float64(do)))
		case []any:
			propOpts = append(propOpts, mcp.DefaultArray(do))
		default:
			propOpts = append(propOpts, mcp.DefaultString(do.(string)))
		}
	}

	if enums, ok := param["enum"].([]interface{}); ok && len(enums) > 0 {
		var enumValues []string
		for _, e := range enums {
			enumValues = append(enumValues, fmt.Sprintf("%v", e))
		}
		propOpts = append(propOpts, mcp.Enum(enumValues...))
	}

	switch paramType {
	case "string":
		return mcp.WithString(name, propOpts...), nil
	case "number":
		return mcp.WithNumber(name, propOpts...), nil
	case "integer":
		return mcp.WithNumber(name, propOpts...), nil
	case "boolean":
		return mcp.WithBoolean(name, propOpts...), nil
	case "object":
		props, ok := param["props"].(map[string]interface{})
		if ok {
			propOpts = append(propOpts, mcp.Properties(props))
		}
		return mcp.WithObject(name, propOpts...), nil
	case "array":
		items, ok := param["items"]
		if ok {
			propOpts = append(propOpts, mcp.Items(items))
		}
		return mcp.WithArray(name, propOpts...), nil
	default:
		return mcp.WithString(name, propOpts...), nil // 默认处理为字符串
	}
}

// createHandler 为工具生成处理函数
func (a *OpenAPIToMCPAdapter) createHandler(path, method string) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.Params.Arguments

		url := a.backendBaseUrl + path
		for key, value := range args {
			if strings.Contains(url, "{"+key+"}") {
				url = strings.Replace(url, "{"+key+"}", value.(string), 1)
			}
		}

		// 构建请求体
		var body []byte
		if method == "post" || method == "put" || method == "patch" {
			body, _ = json.Marshal(args)
		}

		// 发送请求
		client := &http.Client{}
		req, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(string(body)))
		if err != nil {
			return nil, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// 读取响应
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(respBody)), nil
	}
}

// Start 启动 MCP 服务器
func (a *OpenAPIToMCPAdapter) Start(ctx context.Context) error {
	for toolName, tool := range a.tools {
		handler := a.handlers[toolName]
		a.server.AddTool(*tool, handler)
	}

	log.Info().Msgf("start mcp adapter at %s", a.addrs)
	s := server.NewSSEServer(a.server)
	go func() {
		<-ctx.Done()
		s.Shutdown(ctx)
	}()

	return s.Start(a.addrs)
}
