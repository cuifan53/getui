package types

type PushSingleReq struct {
	RequestId   string       `json:"request_id,omitempty"`
	Audience    *Audience    `json:"audience,omitempty"`
	Settings    *Setting     `json:"settings,omitempty"`
	PushMessage *PushMessage `json:"push_message,omitempty"`
	PushChannel *PushChannel `json:"push_channel,omitempty"`
}

type PushSingleBatchReq struct {
	IsAsync bool             `json:"is_async,omitempty"`
	MsgList []*PushSingleReq `json:"msg_list,omitempty"`
}

type PushListMessageReq struct {
	RequestId   string       `json:"request_id,omitempty"`
	GroupName   string       `json:"group_name,omitempty"`
	Settings    *Setting     `json:"settings,omitempty"`
	PushMessage *PushMessage `json:"push_message,omitempty"`
	PushChannel *PushChannel `json:"push_channel,omitempty"`
}

type PushListReq struct {
	Audience *Audience `json:"audience,omitempty"`
	IsAsync  bool      `json:"is_async,omitempty"`
	Taskid   string    `json:"taskid,omitempty"`
}

type PushAllReq struct {
	RequestId   string       `json:"request_id,omitempty"`
	GroupName   string       `json:"group_name,omitempty"`
	Audience    *Audience    `json:"audience,omitempty"`
	Settings    *Setting     `json:"settings,omitempty"`
	PushMessage *PushMessage `json:"push_message,omitempty"`
	PushChannel *PushChannel `json:"push_channel,omitempty"`
}

type PushTagReq struct {
	RequestId   string       `json:"request_id,omitempty"`
	GroupName   string       `json:"group_name,omitempty"`
	Audience    *Audience    `json:"audience,omitempty"`
	Settings    *Setting     `json:"settings,omitempty"`
	PushMessage *PushMessage `json:"push_message,omitempty"`
	PushChannel *PushChannel `json:"push_channel,omitempty"`
}

type PushFastCustomTagReq struct {
	RequestId   string       `json:"request_id,omitempty"`
	Audience    *Audience    `json:"audience,omitempty"`
	Settings    *Setting     `json:"settings,omitempty"`
	PushMessage *PushMessage `json:"push_message,omitempty"`
	PushChannel *PushChannel `json:"push_channel,omitempty"`
}
