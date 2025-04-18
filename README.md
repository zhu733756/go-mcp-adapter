# OpenAPI to MCP Adapter

## Introduction

Thanks to `github.com/mark3labs/mcp-go`, I make this package!

The OpenAPI to MCP Adapter is a robust tool designed to swiftly transform servers adhering to the OpenAPI specification into MCP (Model Control Protocol) servers. By incorporating OpenAPI static configuration, it seamlessly bridges conventional RESTful APIs with the MCP ecosystem, enabling your services to interact with a variety of clients via the MCP protocol.

## Key Features

- **Protocol Transformation** : Converts interfaces defined by OpenAPI into MCP - compatible endpoints, achieving seamless communication between different protocols.
- **Tool Registration** : Automatically registers corresponding MCP tools based on the OpenAPI specification, streamlining the tool management process.
- **Flexible Extension** : Supports custom tool processing logic to meet specific business needs.

## Getting Started

### Initializing the Adapter

```go
adapter, err := openapimcp.NewOpenAPIToMCPAdapter("MyService", "v1", "http://backend:8080", ":9090")
if err != nil {
    log.Fatal(err)
}
```

## Example Code

Complete usage examples are available in the `examples` directory:

- **Server Example (`examples/server`)** : Demonstrates how to create an MCP server using the OpenAPI to MCP Adapter, including tool registration and custom processing logic.
- **Client Example (`examples/client`)** : Provides example code for interacting with the converted MCP service.

## Roadmap

**Management Tool Development** :

- **MCP Server Manager** : Build a dedicated MCP server management tool with a unified console interface, enabling centralized management and monitoring of multiple MCP services by operations personnel.
- **Service Registration and Discovery** : Implement automated service registration and discovery mechanisms so MCP services can easily locate each other in distributed environments, enhancing system scalability and fault tolerance.
- **Worker Pool** : Design an efficient worker pool system to optimize task distribution and resource utilization, improving the overall performance and response speed of MCP services.
- **Reverse Proxy** : Develop components with reverse proxy capabilities to achieve request forwarding and load balancing, enhancing system stability and availability.
- **Rate Limiting Protection** : Introduce rate limiting protection mechanisms to prevent server overload and ensure stable service operation under high - concurrency conditions.
