package goresty

import (
	"context"
	"github.com/go-resty/resty/v2"
	"sync"
)

var client *resty.Client
var clientMu sync.RWMutex

func InitClient(c *resty.Client) {
	clientMu.Lock()
	client = c
	clientMu.Unlock()
}

func GetClient() *resty.Client {
	clientMu.RLock()
	c := client
	clientMu.RUnlock()
	if c != nil {
		return c
	}

	clientMu.Lock()
	if client == nil {
		client = resty.New()
	}
	c = client
	clientMu.Unlock()
	return c
}

func Request() *resty.Request {
	return GetClient().R()
}

func RequestWithContext(ctx context.Context) *resty.Request {
	return GetClient().R().SetContext(ctx)
}

// ------------------------ GET -----------------------------------

func Get(url string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).Get(url)
}

func GetQueryString(url string, object interface{}, queryString string) (*resty.Response, error) {
	return GetClient().R().SetQueryString(queryString).SetResult(object).Get(url)
}

func GetQueryParams(url string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetQueryParams(uriVariables).SetResult(object).Get(url)
}

func GetPathParams(url string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetPathParams(uriVariables).SetResult(object).Get(url)
}

func GetHeader(url string, headers map[string]string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetResult(object).Get(url)
}

func GetQueryStringHeader(url string, headers map[string]string, object interface{}, queryString string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetQueryString(queryString).SetResult(object).Get(url)
}

func GetQueryParamsHeader(url string, headers map[string]string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetQueryParams(uriVariables).SetResult(object).Get(url)
}

func GetPathParamsHeader(url string, headers map[string]string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetPathParams(uriVariables).SetResult(object).Get(url)
}

// ------------------------ POST -----------------------------------

func Post(url string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).Post(url)
}

func PostBody(url string, requestBody interface{}, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).SetBody(requestBody).Post(url)
}

func PostBodyQueryString(url string, requestBody interface{}, object interface{}, queryString string) (*resty.Response, error) {
	return GetClient().R().SetBody(requestBody).SetQueryString(queryString).SetResult(object).Post(url)
}

func PostBodyQueryParams(url string, requestBody interface{}, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetQueryParams(uriVariables).SetBody(requestBody).SetResult(object).Post(url)
}

func PostBodyPathParams(url string, requestBody interface{}, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetPathParams(uriVariables).SetBody(requestBody).SetResult(object).Post(url)
}

func PostHeader(url string, headers map[string]string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetResult(object).Post(url)
}

func PostBodyHeader(url string, requestBody interface{}, headers map[string]string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetBody(requestBody).SetResult(object).Post(url)
}

func PostBodyQueryStringHeader(url string, requestBody interface{}, headers map[string]string, object interface{}, queryString string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetBody(requestBody).SetQueryString(queryString).SetResult(object).Post(url)
}

func PostBodyQueryParamsHeader(url string, requestBody interface{}, headers map[string]string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetBody(requestBody).SetQueryParams(uriVariables).SetResult(object).Post(url)
}

func PostBodyPathParamsHeader(url string, requestBody interface{}, headers map[string]string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetBody(requestBody).SetPathParams(uriVariables).SetResult(object).Post(url)
}

// ------------------------ PUT -----------------------------------

func Put(url string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).Put(url)
}

func PutBody(url string, requestBody interface{}, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).SetBody(requestBody).Put(url)
}

// ------------------------ DELETE -----------------------------------

func Delete(url string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).Delete(url)
}

func DeleteBody(url string, requestBody interface{}, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).SetBody(requestBody).Delete(url)
}
