package getui

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/cuifan53/getui/types"

	"github.com/monaco-io/request"
)

type ClientOptions struct {
	AppID        string
	AppKey       string
	AppSecret    string
	MasterSecret string
}

func NewClient(options ClientOptions) *client {
	return &client{
		appID:        options.AppID,
		appKey:       options.AppKey,
		appSecret:    options.AppSecret,
		masterSecret: options.MasterSecret,
	}
}

type client struct {
	appID        string
	appKey       string
	appSecret    string
	masterSecret string
	token        string
	expireTime   int64
	mu           sync.RWMutex
}

type requestOptions struct {
	Url    string
	Method string
	Query  map[string]string
	Body   interface{}
}

// 加签
func (c *client) sign() (string, string) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	original := c.appKey + timestamp + c.masterSecret
	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), timestamp
}

// 请求
func (c *client) request(resp interface{}, options requestOptions) error {
	rc := request.Client{
		URL:    types.BaseUrl + c.appID + options.Url,
		Method: options.Method,
		Header: map[string]string{
			"token": c.token,
		},
		Query: options.Query,
		JSON:  options.Body,
	}
	r := rc.Send().Scan(resp)
	if !r.OK() {
		return r.Error()
	}
	return nil
}

// 校验token 如果token过期 重新获取token
func (c *client) checkToken() error {
	if time.Now().UnixNano()/1e6 < c.expireTime {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if time.Now().UnixNano()/1e6 < c.expireTime {
		return nil
	}
	resp, err := c.Auth()
	if err != nil {
		return err
	}
	c.token = resp.Data.Token
	expireTime, err := strconv.ParseInt(resp.Data.ExpireTime, 10, 64)
	if err != nil {
		return err
	}
	c.expireTime = expireTime
	return nil
}

// 请求数据预处理
func (c *client) postBody(req interface{}) (map[string]interface{}, error) {
	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body := make(map[string]interface{})
	if err := json.Unmarshal(bytes, &body); err != nil {
		return nil, err
	}
	return body, nil
}
