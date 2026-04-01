package goresty

import (
	"context"
	"net"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Feign struct {
	client *resty.Client
}

func NewFeign(r *resty.Client) *Feign {
	if r == nil {
		r = resty.New()
	}
	return &Feign{client: r}
}

func NewFeignWithClient(hc *http.Client) *Feign {
	return NewFeign(resty.NewWithClient(hc))
}

func NewFeignWithLocalAddr(localAddr net.Addr) *Feign {
	return NewFeign(resty.NewWithLocalAddr(localAddr))
}

func NewFeignWithWithLocalAddr(localAddr net.Addr) *Feign {
	return NewFeignWithLocalAddr(localAddr)
}

func (r *Feign) GetClient() *resty.Client {
	return r.client
}

func (r *Feign) Options(f func(*resty.Client)) *Feign {
	if f != nil {
		f(r.client)
	}
	return r
}

func (r *Feign) OnError(h resty.ErrorHook) *Feign {
	if h != nil {
		r.client.OnError(h)
	}
	return r
}

func (r *Feign) OnBeforeRequest(m resty.RequestMiddleware) *Feign {
	if m != nil {
		r.client.OnBeforeRequest(m)
	}
	return r
}

func (r *Feign) OnAfterResponse(m resty.ResponseMiddleware) *Feign {
	if m != nil {
		r.client.OnAfterResponse(m)
	}
	return r
}

func (r *Feign) Request() *resty.Request {
	return r.client.R()
}

func (r *Feign) RequestWithContext(ctx context.Context) *resty.Request {
	return r.client.R().SetContext(ctx)
}
