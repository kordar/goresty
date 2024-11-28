package goresty_test

import (
	"github.com/go-resty/resty/v2"
	"github.com/kordar/goresty"
	"log"
	"testing"
	"time"
)

var feign = goresty.NewFeign(nil).OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
	request.EnableTrace()
	request.SetHeader("Auth2", "MMM")
	if request.Header.Get("Auth") != "" {
		log.Println("=====>>>>>>>>", request.Header.Get("Auth"))
	}
	return nil
}).Options(func(client *resty.Client) {
	client.SetBaseURL("https://www.sina22.com")
	client.SetDebug(true)
	client.SetTimeout(time.Second * 3)
	client.SetRetryCount(2)
	client.SetRetryWaitTime(3 * time.Second)
}).OnError(func(request *resty.Request, err error) {
	log.Println("xxxxxxxxxxxxxx", err)
})

func TestGetHeader(t *testing.T) {
	resp, err := feign.Request().SetHeaders(map[string]string{"Auth": "AAABB"}).Get("/")
	log.Println("=========", resp, err)
}

func TestGetBody(t *testing.T) {
	body := struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}{}
	resp, err := feign.Request().SetHeaders(map[string]string{"Auth": "AAABB"}).
		SetBody(&body).
		Get("http://192.168.48.125:82/prod-api/md/basic/list?pageNum=1&pageSize=10&foodNumId=&tenantId=101")
	log.Println("=========", resp, err, body)
}
