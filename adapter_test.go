package gmadapter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestGenerateTools_NumberParameter(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testNumber": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Test Number Summary",
					"description": "Test Number Description",
					"parameters": []interface{}{
						map[string]interface{}{
							"name":        "testNumberParam",
							"description": "Test Number Parameter",
							"schema": map[string]interface{}{
								"type": "number",
							},
							"required": true,
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testNumber_get"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Number Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	assert.Equal(t, 1, len(tool.InputSchema.Required))

	prop := tool.InputSchema.Properties["testNumberParam"].(map[string]interface{})
	assert.Equal(t, "number", prop["type"])
	assert.Equal(t, "Test Number Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 1, len(required))
	assert.Equal(t, "testNumberParam", tool.InputSchema.Required[0])
}

func TestGenerateTools_IntegerParameter(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testInteger": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Test Integer Summary",
					"description": "Test Integer Description",
					"parameters": []interface{}{
						map[string]interface{}{
							"name":        "testIntegerParam",
							"description": "Test Integer Parameter",
							"schema": map[string]interface{}{
								"type": "integer",
							},
							"required": true,
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testInteger_get"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Integer Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	assert.Equal(t, 1, len(tool.InputSchema.Required))

	prop := tool.InputSchema.Properties["testIntegerParam"].(map[string]interface{})
	assert.Equal(t, "number", prop["type"])
	assert.Equal(t, "Test Integer Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 1, len(required))
	assert.Equal(t, "testIntegerParam", tool.InputSchema.Required[0])
}

func TestGenerateTools_BooleanParameter(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testBoolean": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Test Boolean Summary",
					"description": "Test Boolean Description",
					"parameters": []interface{}{
						map[string]interface{}{
							"name":        "testBooleanParam",
							"description": "Test Boolean Parameter",
							"schema": map[string]interface{}{
								"type": "boolean",
							},
							"required": true,
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testBoolean_get"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Boolean Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	assert.Equal(t, 1, len(tool.InputSchema.Required))

	prop := tool.InputSchema.Properties["testBooleanParam"].(map[string]interface{})
	assert.Equal(t, "boolean", prop["type"])
	assert.Equal(t, "Test Boolean Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 1, len(required))
	assert.Equal(t, "testBooleanParam", tool.InputSchema.Required[0])
}

func TestGenerateTools_ObjectParameter(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testObject": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Test Object Summary",
					"description": "Test Object Description",
					"parameters": []interface{}{
						map[string]interface{}{
							"name":        "testObjectParam",
							"description": "Test Object Parameter",
							"schema": map[string]interface{}{
								"type": "object",
								"properties": map[string]interface{}{
									"nestedObjectParam": map[string]interface{}{
										"type":        "string",
										"description": "Nested Object Parameter",
									},
								},
							},
							"required": true,
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testObject_get"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Object Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	assert.Equal(t, 1, len(tool.InputSchema.Required))

	prop := tool.InputSchema.Properties["testObjectParam"].(map[string]interface{})
	assert.Equal(t, "object", prop["type"])
	assert.Equal(t, "Test Object Parameter", prop["description"])

	assert.Equal(t, 3, len(prop))
	nestedProp, ok := prop["properties"].(map[string]interface{})
	assert.True(t, ok)
	assert.NotNil(t, nestedProp)
	assert.Equal(t, 1, len(nestedProp))

	nestedNestedProp, ok := nestedProp["nestedObjectParam"].(map[string]interface{})
	assert.True(t, ok)
	assert.NotNil(t, nestedNestedProp)
	assert.Equal(t, 2, len(nestedNestedProp))
	assert.Equal(t, "string", nestedNestedProp["type"])
	assert.Equal(t, "Nested Object Parameter", nestedNestedProp["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 1, len(required))
	assert.Equal(t, "testObjectParam", tool.InputSchema.Required[0])
}

func TestGenerateTools_ArrayParameter(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testArray": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "Test Array Summary",
					"description": "Test Array Description",
					"parameters": []interface{}{
						map[string]interface{}{
							"name":        "testArrayParam",
							"description": "Test Array Parameter",
							"schema": map[string]interface{}{
								"type": "array",
								"items": map[string]interface{}{
									"type": "string",
								},
							},
							"required": true,
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testArray_get"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Array Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	assert.Equal(t, 1, len(tool.InputSchema.Required))

	prop := tool.InputSchema.Properties["testArrayParam"].(map[string]interface{})
	assert.Equal(t, "array", prop["type"])
	assert.Equal(t, "Test Array Parameter", prop["description"])

	items, ok := prop["items"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "string", items["type"])

	required := tool.InputSchema.Required
	assert.Equal(t, 1, len(required))
	assert.Equal(t, "testArrayParam", tool.InputSchema.Required[0])
}

func TestGenerateTools_POSTMethod(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testPost": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Test Post Summary",
					"description": "Test Post Description",
					"parameters": []interface{}{
						map[string]interface{}{
							"name":        "testPostParam",
							"description": "Test Post Parameter",
							"schema": map[string]interface{}{
								"type": "string",
							},
							"required": true,
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testPost_post"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Post Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	assert.Equal(t, 1, len(tool.InputSchema.Required))

	prop := tool.InputSchema.Properties["testPostParam"].(map[string]interface{})
	assert.Equal(t, "string", prop["type"])
	assert.Equal(t, "Test Post Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 1, len(required))
	assert.Equal(t, "testPostParam", tool.InputSchema.Required[0])
}

func TestGenerateTools_PUTMethod(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testPut": map[string]interface{}{
				"put": map[string]interface{}{
					"summary":     "Test Put Summary",
					"description": "Test Put Description",
					"parameters": []interface{}{
						map[string]interface{}{
							"name":        "testPutParam",
							"description": "Test Put Parameter",
							"schema": map[string]interface{}{
								"type": "string",
							},
							"required": true,
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testPut_put"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Put Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	assert.Equal(t, 1, len(tool.InputSchema.Required))

	prop := tool.InputSchema.Properties["testPutParam"].(map[string]interface{})
	assert.Equal(t, "string", prop["type"])
	assert.Equal(t, "Test Put Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 1, len(required))
	assert.Equal(t, "testPutParam", tool.InputSchema.Required[0])
}

func TestGenerateTools_DELETEMethod(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testDelete": map[string]interface{}{
				"delete": map[string]interface{}{
					"summary":     "Test Delete Summary",
					"description": "Test Delete Description",
					"parameters": []interface{}{
						map[string]interface{}{
							"name":        "testDeleteParam",
							"description": "Test Delete Parameter",
							"schema": map[string]interface{}{
								"type": "string",
							},
							"required": true,
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testDelete_delete"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Delete Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	assert.Equal(t, 1, len(tool.InputSchema.Required))

	prop := tool.InputSchema.Properties["testDeleteParam"].(map[string]interface{})
	assert.Equal(t, "string", prop["type"])
	assert.Equal(t, "Test Delete Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 1, len(required))
	assert.Equal(t, "testDeleteParam", tool.InputSchema.Required[0])
}

func TestGenerateTools_RequestBody_String(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testRequestBodyString": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Test Request Body String Summary",
					"description": "Test Request Body String Description",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"testRequestBodyStringParam": map[string]interface{}{
											"type":        "string",
											"description": "Test Request Body String Parameter",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testRequestBodyString_post"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Request Body String Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	assert.Equal(t, 0, len(tool.InputSchema.Required))

	prop := tool.InputSchema.Properties["testRequestBodyStringParam"].(map[string]interface{})
	assert.Equal(t, "string", prop["type"])
	assert.Equal(t, "Test Request Body String Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 0, len(required))
}

func TestGenerateTools_RequestBody_Number(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testRequestBodyNumber": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Test Request Body Number Summary",
					"description": "Test Request Body Number Description",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"testRequestBodyNumberParam": map[string]interface{}{
											"type":        "number",
											"description": "Test Request Body Number Parameter",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testRequestBodyNumber_post"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Request Body Number Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))

	prop := tool.InputSchema.Properties["testRequestBodyNumberParam"].(map[string]interface{})
	assert.Equal(t, "number", prop["type"])
	assert.Equal(t, "Test Request Body Number Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 0, len(required))
}

func TestGenerateTools_RequestBody_Integer(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testRequestBodyInteger": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Test Request Body Integer Summary",
					"description": "Test Request Body Integer Description",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"testRequestBodyIntegerParam": map[string]interface{}{
											"type":        "integer",
											"description": "Test Request Body Integer Parameter",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testRequestBodyInteger_post"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Request Body Integer Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))

	prop := tool.InputSchema.Properties["testRequestBodyIntegerParam"].(map[string]interface{})
	assert.Equal(t, "number", prop["type"])
	assert.Equal(t, "Test Request Body Integer Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 0, len(required))
}

func TestGenerateTools_RequestBody_Boolean(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testRequestBodyBoolean": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Test Request Body Boolean Summary",
					"description": "Test Request Body Boolean Description",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"testRequestBodyBooleanParam": map[string]interface{}{
											"type":        "boolean",
											"description": "Test Request Body Boolean Parameter",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testRequestBodyBoolean_post"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Request Body Boolean Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))
	prop := tool.InputSchema.Properties["testRequestBodyBooleanParam"].(map[string]interface{})
	assert.Equal(t, "boolean", prop["type"])
	assert.Equal(t, "Test Request Body Boolean Parameter", prop["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 0, len(required))
}

func TestGenerateTools_RequestBody_Object(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testRequestBodyObject": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Test Request Body Object Summary",
					"description": "Test Request Body Object Description",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"testRequestBodyObjectParam": map[string]interface{}{
											"type":        "object",
											"description": "Test Request Body Object Parameter",
											"properties": map[string]interface{}{
												"nestedObjectParam": map[string]interface{}{
													"type":        "string",
													"description": "Nested Object Parameter",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testRequestBodyObject_post"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Request Body Object Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))

	prop := tool.InputSchema.Properties["testRequestBodyObjectParam"].(map[string]interface{})
	assert.Equal(t, "object", prop["type"])
	assert.Equal(t, "Test Request Body Object Parameter", prop["description"])

	nestedProp := prop["properties"].(map[string]interface{})["nestedObjectParam"].(map[string]interface{})
	assert.Equal(t, "string", nestedProp["type"])
	assert.Equal(t, "Nested Object Parameter", nestedProp["description"])

	required := tool.InputSchema.Required
	assert.Equal(t, 0, len(required))
}

func TestGenerateTools_RequestBody_Array(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testRequestBodyArray": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Test Request Body Array Summary",
					"description": "Test Request Body Array Description",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"testRequestBodyArrayParam": map[string]interface{}{
											"type": "array",
											"items": map[string]interface{}{
												"type": "string",
											},
											"description": "Test Request Body Array Parameter",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testRequestBodyArray_post"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Request Body Array Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))

	prop := tool.InputSchema.Properties["testRequestBodyArrayParam"].(map[string]interface{})
	assert.Equal(t, "array", prop["type"])
	assert.Equal(t, "Test Request Body Array Parameter", prop["description"])

	items := prop["items"].(map[string]interface{})
	assert.Equal(t, "string", items["type"])

	required := tool.InputSchema.Required
	assert.Equal(t, 0, len(required))
}

func TestGenerateTools_RequestBody_StringWithDefault(t *testing.T) {
	openAPI := map[string]interface{}{
		"paths": map[string]interface{}{
			"/testRequestBodyStringWithDefault": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Test Request Body String With Default Summary",
					"description": "Test Request Body String With Default Description",
					"requestBody": map[string]interface{}{
						"content": map[string]interface{}{
							"application/json": map[string]interface{}{
								"schema": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"testRequestBodyStringWithDefaultParam": map[string]interface{}{
											"type":        "string",
											"description": "Test Request Body String With Default Parameter",
											"default":     "default_value",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	adapter := &OpenAPIToMCPAdapter{
		backendBaseUrl: "http://localhost:8080",
		addrs:          "localhost:8080",
		openAPI:        openAPI,
		tools:          make(map[string]*mcp.Tool),
		handlers:       make(map[string]func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}

	err := adapter.GenerateTools()
	assert.NoError(t, err)
	assert.Len(t, adapter.tools, 1)

	toolName := "_testRequestBodyStringWithDefault_post"
	tool, exists := adapter.tools[toolName]
	assert.True(t, exists)
	assert.NotNil(t, tool)
	assert.Equal(t, toolName, tool.Name)
	assert.Equal(t, "Test Request Body String With Default Summary", tool.Description)

	assert.Equal(t, 1, len(tool.InputSchema.Properties))

	prop := tool.InputSchema.Properties["testRequestBodyStringWithDefaultParam"].(map[string]interface{})
	assert.Equal(t, "string", prop["type"])
	assert.Equal(t, "Test Request Body String With Default Parameter", prop["description"])
	assert.Equal(t, "default_value", prop["default"])

	required := tool.InputSchema.Required
	assert.Equal(t, 0, len(required))
}
