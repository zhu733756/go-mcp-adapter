# OpenAPI to MCP Adapter

## 介绍

OpenAPI to MCP Adapter 是一个强大的工具，能够将符合 OpenAPI 规范的服务器快速转换为 MCP（Model Control Protocol）服务器。通过提供 OpenAPI 静态配置，它可以无缝桥接传统的 RESTful API 与 MCP 生态系统，使您的服务能够以 MCP 协议与各类客户端进行交互。

## 功能亮点

- **协议转换** ：将 OpenAPI 定义的接口转换为 MCP 兼容的服务端点，实现不同协议间的无缝通信。
- **工具注册** ：自动根据 OpenAPI 规范注册对应的 MCP 工具，简化工具管理流程。
- **灵活扩展** ：支持自定义工具处理逻辑，满足特定业务需求。

## 使用方法

### 初始化适配器

```go
adapter, err := openapimcp.NewOpenAPIToMCPAdapter("MyService", "v1", "http://backend:8080", ":9090")
if err != nil {
    log.Fatal(err)
}
```

## 示例代码

完整的使用示例位于 `examples` 目录下：

- **服务端示例（examples/server）** ：展示了如何使用 OpenAPI to MCP Adapter 创建一个 MCP 服务端，包括工具注册和自定义处理逻辑。
- **客户端示例（examples/client）** ：提供了如何与转换后的 MCP 服务进行交互的示例代码。

## 未来展望

**管理面工具开发** ：

- **MCP Server Manager** ：构建一个专门用于管理 MCP 服务器的工具，提供统一的控制台界面，方便运维人员对多个 MCP 服务进行集中管理和监控。
- **服务注册与发现** ：实现自动化的服务注册和发现机制，使 MCP 服务能够在分布式环境中轻松找到彼此，提高系统的可扩展性和容错性。
- **工作池** ：设计高效的工作池系统，优化任务分配和资源利用，提升 MCP 服务的整体性能和响应速度。
- **反向代理** ：开发具备反向代理功能的组件，实现请求的转发和负载均衡，增强系统的稳定性和可用性。
- **限流保护** ：加入限流保护机制，防止服务器过载，保障服务在高并发情况下的稳定运行。
