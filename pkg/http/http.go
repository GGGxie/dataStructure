package dhttp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Resp struct {
	RespBody []byte
	Code     int
}

//发送GET请求
//url:请求地址
//response:请求返回的内容
func Get(url string, data interface{}, header map[string]string, query map[string]string) (*Resp, error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	//设置header参数
	for k, v := range header {
		req.Header.Add(k, v)
	}
	//设置query参数
	tempQuery := req.URL.Query()
	for k, v := range query {
		tempQuery.Add(k, v)
	}
	req.URL.RawQuery = tempQuery.Encode()

	//请求 https 网站跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//获取response
	result, err := ParseResp(resp)
	return &Resp{
		Code:     resp.StatusCode,
		RespBody: result,
	}, err
}

//发送POST请求
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
//content:请求放回的内容
func Post(url string, data interface{}, header map[string]string, query map[string]string) (*Resp, error) {
	var jsonStr []byte
	switch data.(type) {
	case []byte:
		jsonStr = data.([]byte)
	default:
		jsonStr, _ = json.Marshal(data)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	//设置header参数
	for k, v := range header {
		req.Header.Add(k, v)
	}
	//设置query参数
	tempQuery := req.URL.Query()
	for k, v := range query {
		tempQuery.Add(k, v)
	}
	req.URL.RawQuery = tempQuery.Encode()

	req.Header.Add("Content-Type", "application/json")
	//请求 https 网站跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ParseResp(resp)
	return &Resp{
		Code:     resp.StatusCode,
		RespBody: result,
	}, nil
}

func ParseResp(resp *http.Response) ([]byte, error) {
	if result, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
