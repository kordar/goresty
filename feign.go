package goresty

import (
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
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

func NewFeignWithWithLocalAddr(localAddr net.Addr) *Feign {
	return NewFeign(resty.NewWithLocalAddr(localAddr))
}

func (r *Feign) GetClient() *resty.Client {
	return r.client
}

// ------------------- error ----------------------

func (r *Feign) Options(f func(*resty.Client)) *Feign {
	f(r.client)
	return r
}

// ------------------- error ----------------------

func (r *Feign) OnError(h resty.ErrorHook) *Feign {
	r.client.OnError(h)
	return r
}

//  ---------------------- middleware -----------------------------

func (r *Feign) OnBeforeRequest(m resty.RequestMiddleware) *Feign {
	r.client.OnBeforeRequest(m)
	return r
}

func (r *Feign) OnAfterResponse(m resty.ResponseMiddleware) *Feign {
	r.client.OnAfterResponse(m)
	return r
}

// Request Request请求调用方式
func (r *Feign) Request() *resty.Request {
	return r.client.R()
}
