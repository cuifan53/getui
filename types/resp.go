package types

type CommonResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type AuthResp struct {
	CommonResp
	Data struct {
		ExpireTime string `json:"expire_time"`
		Token      string `json:"token"`
	} `json:"data"`
}

type PushCommonResp struct {
	CommonResp
	Data map[string]map[string]string `json:"data"`
}

type PushTaskResp struct {
	CommonResp
	Data *struct {
		Taskid string `json:"taskid"`
	} `json:"data"`
}

type GetScheduleTaskResp struct {
	CommonResp
	Data map[string]*struct {
		CreateTime          string `json:"create_time"`
		Status              string `json:"status"`
		TransmissionContent string `json:"transmission_content"`
		PushTime            string `json:"push_time"`
	} `json:"data"`
}
