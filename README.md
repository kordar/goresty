# goresty

基于 [resty](https://github.com/go-resty/resty) 的轻量封装，提供：

- 包级函数：Get/Post 及常见参数/headers 组合
- Feign：链式配置 Client，并生成 Request

## 安装

```bash
go get github.com/kordar/goresty
```

## 全局 Client

```go
package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/kordar/goresty"
	"time"
)

func main() {
	c := resty.New().SetTimeout(3 * time.Second)
	goresty.InitClient(c)

	var out map[string]any
	_, _ = goresty.Get("https://example.com", &out)
}
```

## Feign（推荐）

```go
package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/kordar/goresty"
	"log"
	"time"
)

func main() {
	feign := goresty.NewFeign(nil).
		Options(func(c *resty.Client) {
			c.SetBaseURL("https://example.com")
			c.SetTimeout(3 * time.Second)
		}).
		OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
			r.SetHeader("Auth", "token")
			return nil
		}).
		OnError(func(r *resty.Request, err error) {
			log.Println("request error:", err)
		})

	var out map[string]any
	_, _ = feign.Request().SetResult(&out).Get("/api/ping")
}
```
