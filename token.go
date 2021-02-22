package getui

import (
	"github.com/cuifan53/getui/types"
)

// 鉴权
func (c *client) Auth() (*types.AuthResp, error) {
	sign, timestamp := c.sign()
	resp := types.AuthResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/auth",
		Method: types.MethodPost,
		Body: map[string]string{
			"sign":      sign,
			"timestamp": timestamp,
			"appkey":    c.appKey,
		},
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}
