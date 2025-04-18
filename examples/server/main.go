package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	gmadapter "github.com/zhu733756/go-mcp-adapter"
)

type Config struct {
	OpenAPI        string `json:"openapi" yaml:"openapi"`
	ServerBaseUrl  string `json:"server_base_url" yaml:"server_base_url"`
	AdapterAddress string `json:"adapter_address" yaml:"adapter_address"`
}

func main() {
	config := Config{
		OpenAPI:        `D:\code\src\github.com\go-mcp-adapter\examples\server\openapi.yaml`,
		ServerBaseUrl:  "http://localhost:8080",
		AdapterAddress: "0.0.0.0:8081",
	}

	// backend servers
	go OpenAPIServer(config.ServerBaseUrl)

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-sigChan
		cancel()
	}()

	StartMCPAdapter(ctx, config.OpenAPI, config.ServerBaseUrl, config.AdapterAddress)
}

func OpenAPIServer(addrs string) {
	// User represents the user data structure
	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var users = []User{}
	var mu = sync.RWMutex{}

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		mu.RLock()
		defer mu.RUnlock()
		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		mu.Lock()
		defer mu.Unlock()

		newID := 1
		if len(users) > 0 {
			newID = users[len(users)-1].ID + 1
		}
		user.ID = newID

		users = append(users, user)
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	})

	gin.SetMode(gin.ReleaseMode)

	// Start the server
	r.Run(":8080")
}

func StartMCPAdapter(ctx context.Context, openapi, backendBaseUrl, myAddrs string) {
	adapter, err := gmadapter.NewOpenAPIToMCPAdapter("OpenAPI MCP Server", "1.0.0", backendBaseUrl, myAddrs)
	if err != nil {
		fmt.Printf("😡 init mcp server error: %v\n", err)
		return
	}

	funcs := []func() error{
		func() error { return adapter.LoadOpenAPI(openapi) },
		func() error { return adapter.GenerateTools() },
		func() error { return adapter.Start(ctx) },
	}

	WrapperServer(funcs)
}

func WrapperServer(funcs []func() error) (err error) {
	defer func() {
		if err != nil {
			log.Err(err).Msgf("😡 Server error: %v\n", err)
		}

		log.Info().Msg("👋 Server stopped")
	}()

	for _, f := range funcs {
		if err = f(); err != nil {
			return err
		}
	}

	return nil
}
