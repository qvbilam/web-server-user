package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func params(d map[string]interface{}) *url.Values {
	p := url.Values{}
	for k, v := range d {
		p.Set(k, v.(string))
		p.Set(k, v.(string))
	}
	return &p
}

func Get(u string, d map[string]interface{}) ([]byte, error) {
	p := params(d)
	urlWithParams := u + "?" + p.Encode()
	// 创建一个GET请求
	resp, err := http.Get(urlWithParams)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Post(u string, d map[string]interface{}) ([]byte, error) {
	ct := "application/x-www-form-urlencoded"
	return postByContentType(u, d, ct)
}

func PostJson(u string, d map[string]interface{}) ([]byte, error) {
	ct := "application/json"
	return postByContentType(u, d, ct)
}

func postByContentType(u string, d map[string]interface{}, contentType string) ([]byte, error) {
	p := params(d)

	resp, err := http.Post(u, contentType, strings.NewReader(p.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
