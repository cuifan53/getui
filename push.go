package getui

import (
	"strings"

	"github.com/lithammer/shortuuid/v3"

	"github.com/cuifan53/getui/types"
)

// 执行cid单推
func (c *client) PushSingleCid(req *types.PushSingleReq) (*types.PushCommonResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	if req.RequestId == "" {
		req.RequestId = strings.ToLower(shortuuid.New())
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushCommonResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/single/cid",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 执行别名单推
func (c *client) PushSingleAlias(req *types.PushSingleReq) (*types.PushCommonResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	if req.RequestId == "" {
		req.RequestId = strings.ToLower(shortuuid.New())
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushCommonResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/single/alias",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 执行cid批量单推
func (c *client) PushSingleBatchCid(req *types.PushSingleBatchReq) (*types.PushCommonResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	for _, v := range req.MsgList {
		if v.RequestId == "" {
			v.RequestId = strings.ToLower(shortuuid.New())
		}
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushCommonResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/single/batch/cid",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 执行别名批量单推
func (c *client) PushSingleBatchAlias(req *types.PushSingleBatchReq) (*types.PushCommonResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	for _, v := range req.MsgList {
		if v.RequestId == "" {
			v.RequestId = strings.ToLower(shortuuid.New())
		}
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushCommonResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/single/batch/alias",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建消息
func (c *client) PushListMessage(req *types.PushListMessageReq) (*types.PushTaskResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	if req.RequestId == "" {
		req.RequestId = strings.ToLower(shortuuid.New())
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushTaskResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/list/message",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 执行cid批量推
func (c *client) PushListCid(req *types.PushListReq) (*types.PushCommonResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushCommonResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/list/cid",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 执行别名批量推
func (c *client) PushListAlias(req *types.PushListReq) (*types.PushCommonResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushCommonResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/list/alias",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 执行群推
func (c *client) PushAll(req *types.PushAllReq) (*types.PushTaskResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushTaskResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/all",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 根据条件筛选用户推送
func (c *client) PushTag(req *types.PushTagReq) (*types.PushTaskResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushTaskResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/tag",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 使用标签快速推送
func (c *client) PushFastCustomTag(req *types.PushFastCustomTagReq) (*types.PushTaskResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	body, err := c.postBody(req)
	if err != nil {
		return nil, err
	}
	resp := types.PushTaskResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/push/fast_custom_tag",
		Method: types.MethodPost,
		Body:   body,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 停止任务
func (c *client) DeleteTask(taskid string) (*types.CommonResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	resp := types.CommonResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/task/" + taskid,
		Method: types.MethodDelete,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 查询定时任务
func (c *client) GetScheduleTask(taskid string) (*types.GetScheduleTaskResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	resp := types.GetScheduleTaskResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/task/schedule/" + taskid,
		Method: types.MethodGet,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// 删除定时任务
func (c *client) DeleteScheduleTask(taskid string) (*types.CommonResp, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}
	resp := types.CommonResp{}
	if err := c.request(&resp, requestOptions{
		Url:    "/task/schedule/" + taskid,
		Method: types.MethodDelete,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}
