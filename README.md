# goresty

基于 [resty](https://github.com/go-resty/resty) 的轻量封装，提供：

- 包级函数：Get/Post/Put/Delete 快捷调用
- Feign：链式配置 Client，生成 Request；支持 Context

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
	"context"
	"time"
)

func main() {
	c := resty.New().SetTimeout(3 * time.Second)
	goresty.InitClient(c)

	var out map[string]any
	_, _ = goresty.Get("https://example.com", &out)

	// 也可直接拿到全局 Request（或注入 Context）
	_, _ = goresty.Request().SetResult(&out).Get("https://example.com")
	_, _ = goresty.RequestWithContext(context.Background()).SetResult(&out).Get("https://example.com")
}
```

## 包级函数

- Get/Post/Put/Delete 基础用法

```go
var out map[string]any
_, _ = goresty.Get("https://example.com", &out)
_, _ = goresty.PostBody("https://example.com/api", map[string]any{"k":"v"}, &out)
_, _ = goresty.PutBody("https://example.com/api/1", map[string]any{"name":"n"}, &out)
_, _ = goresty.Delete("https://example.com/api/1", &out)
```

## Feign（推荐）

```go
package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/kordar/goresty"
	"log"
	"context"
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
	_, _ = feign.RequestWithContext(context.Background()).SetResult(&out).Get("/api/ping")
}
```
