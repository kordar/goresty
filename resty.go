package goresty

import "github.com/go-resty/resty/v2"

var client *resty.Client

func InitClient(c *resty.Client) {
	client = c
}

func GetClient() *resty.Client {
	if client == nil {
		client = resty.New()
	}
	return client
}

// ------------------------ GET -----------------------------------

// Get GET请求调用方式
func Get(url string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).Get(url)
}

// GetQueryString GET请求附加参数
func GetQueryString(url string, object interface{}, queryString string) (*resty.Response, error) {
	return GetClient().R().SetQueryString(queryString).SetResult(object).Get(url)
}

// GetQueryParams GET请求附加参数
func GetQueryParams(url string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetQueryParams(uriVariables).SetResult(object).Get(url)
}

// GetPathParams GET请求附加参数
// "/v1/users/{user}/details"
func GetPathParams(url string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetPathParams(uriVariables).SetResult(object).Get(url)
}

// GetHeader GET请求调用方式
func GetHeader(url string, headers map[string]string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetResult(object).Get(url)
}

// GetQueryStringHeader GET请求附加参数
func GetQueryStringHeader(url string, headers map[string]string, object interface{}, queryString string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetQueryString(queryString).SetResult(object).Get(url)
}

// GetQueryParamsHeader GET请求附加参数
func GetQueryParamsHeader(url string, headers map[string]string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetQueryParams(uriVariables).SetResult(object).Get(url)
}

// GetPathParamsHeader GET请求附加参数
// "/v1/users/{user}/details"
func GetPathParamsHeader(url string, headers map[string]string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetPathParams(uriVariables).SetResult(object).Get(url)
}

// ------------------------ POST -----------------------------------

// Post POST请求调用方式
func Post(url string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).Post(url)
}

// PostBody POST请求调用方式
func PostBody(url string, requestBody interface{}, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetResult(object).SetBody(requestBody).Post(url)
}

// PostBodyQueryString POST请求附加参数
func PostBodyQueryString(url string, requestBody interface{}, object interface{}, queryString string) (*resty.Response, error) {
	return GetClient().R().SetBody(requestBody).SetQueryString(queryString).SetResult(object).Post(url)
}

// PostBodyQueryParams GET请求附加参数
func PostBodyQueryParams(url string, requestBody interface{}, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetQueryParams(uriVariables).SetBody(requestBody).SetResult(object).Post(url)
}

// PostBodyPathParams GET请求附加参数
func PostBodyPathParams(url string, requestBody interface{}, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetPathParams(uriVariables).SetBody(requestBody).SetResult(object).Post(url)
}

// PostHeader GET请求调用方式
func PostHeader(url string, headers map[string]string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetResult(object).Post(url)
}

// PostBodyHeader GET请求调用方式
func PostBodyHeader(url string, requestBody interface{}, headers map[string]string, object interface{}) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetBody(requestBody).SetResult(object).Post(url)
}

// PostBodyQueryStringHeader GET请求附加参数
func PostBodyQueryStringHeader(url string, requestBody interface{}, headers map[string]string, object interface{}, queryString string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetBody(requestBody).SetQueryString(queryString).SetResult(object).Post(url)
}

// PostBodyQueryParamsHeader GET请求附加参数
func PostBodyQueryParamsHeader(url string, requestBody interface{}, headers map[string]string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetBody(requestBody).SetQueryParams(uriVariables).SetResult(object).Post(url)
}

// PostBodyPathParamsHeader GET请求附加参数
func PostBodyPathParamsHeader(url string, requestBody interface{}, headers map[string]string, object interface{}, uriVariables map[string]string) (*resty.Response, error) {
	return GetClient().R().SetHeaders(headers).SetBody(requestBody).SetPathParams(uriVariables).SetResult(object).Post(url)
}
