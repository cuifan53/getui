package types

// 各字段含义见参考文档
// https://docs.getui.com/getui/server/rest_v2/common_args/?id=doc-title-6

type Audience struct {
	Cid           []string       `json:"cid,omitempty"`
	Alias         []string       `json:"alias,omitempty"`
	FastCustomTag string         `json:"fast_custom_tag,omitempty"`
	Tag           []*AudienceTag `json:"tag,omitempty"`
}

type AudienceTag struct {
	Key     string   `json:"key,omitempty"`
	Values  []string `json:"values,omitempty"`
	OptType string   `json:"opt_type,omitempty"`
}

type Setting struct {
	Ttl          int              `json:"ttl,omitempty"`
	Strategy     *SettingStrategy `json:"strategy,omitempty"`
	Speed        int              `json:"speed,omitempty"`
	ScheduleTime int              `json:"schedule_time,omitempty"`
}

type SettingStrategy struct {
	Default int `json:"default,omitempty"`
	Ios     int `json:"ios,omitempty"`
	St      int `json:"st,omitempty"`
}

type PushMessage struct {
	Duration     string                   `json:"duration,omitempty"`
	Notification *PushMessageNotification `json:"notification,omitempty"`
	Transmission string                   `json:"transmission,omitempty"`
	Revoke       *PushMessageRevoke       `json:"revoke,omitempty"`
}

type PushMessageNotification struct {
	Title        string `json:"title,omitempty"`
	Body         string `json:"body,omitempty"`
	BigText      string `json:"big_text,omitempty"`
	BigImage     string `json:"big_image,omitempty"`
	Logo         string `json:"logo,omitempty"`
	LogoUrl      string `json:"logo_url,omitempty"`
	ChannelId    string `json:"channel_id,omitempty"`
	ChannelName  string `json:"channel_name,omitempty"`
	ChannelLevel int    `json:"channel_level,omitempty"`
	ClickType    string `json:"click_type,omitempty"`
	Intent       string `json:"intent,omitempty"`
	Url          string `json:"url,omitempty"`
	Payload      string `json:"payload,omitempty"`
	NotifyId     int    `json:"notify_id,omitempty"`
	RingName     string `json:"ring_name,omitempty"`
	BadgeAddNum  int    `json:"badge_add_num,omitempty"`
}

type PushMessageRevoke struct {
	OldTaskId string `json:"old_task_id,omitempty"`
	Force     bool   `json:"force,omitempty"`
}

type PushChannel struct {
	Ios     *PushChannelIos     `json:"ios,omitempty"`
	Android *PushChannelAndroid `json:"android,omitempty"`
}

type PushChannelIos struct {
	Type           string                      `json:"type,omitempty"`
	Aps            *PushChannelIosAps          `json:"aps,omitempty"`
	AutoBadge      string                      `json:"auto_badge,omitempty"`
	Payload        string                      `json:"payload,omitempty"`
	Multimedia     []*PushChannelIosMultimedia `json:"multimedia,omitempty"`
	ApnsCollapseId string                      `json:"apns-collapse-id,omitempty"`
}

type PushChannelIosAps struct {
	Alert            *PushChannelIosApsAlert `json:"alert,omitempty"`
	ContentAvailable int                     `json:"content-available,omitempty"`
	Sound            string                  `json:"sound,omitempty"`
	Category         string                  `json:"category,omitempty"`
	ThreadId         string                  `json:"thread-id,omitempty"`
}

type PushChannelIosApsAlert struct {
	Title           string   `json:"title,omitempty"`
	Body            string   `json:"body,omitempty"`
	ActionLocKey    string   `json:"action-loc-key,omitempty"`
	LocKey          string   `json:"loc-key,omitempty"`
	LocArgs         []string `json:"loc-args,omitempty"`
	LaunchImage     string   `json:"launch-image,omitempty"`
	TitleLocKey     string   `json:"title-loc-key,omitempty"`
	TitleLocArgs    []string `json:"title-loc-args,omitempty"`
	Subtitle        string   `json:"subtitle,omitempty"`
	SubtitleLocKey  string   `json:"subtitle-loc-key,omitempty"`
	SubtitleLocArgs []string `json:"subtitle-loc-args,omitempty"`
}

type PushChannelIosMultimedia struct {
	Url      string `json:"url,omitempty"`
	Type     int    `json:"type,omitempty"`
	OnlyWifi bool   `json:"only_wifi,omitempty"`
}

type PushChannelAndroid struct {
	Ups *PushChannelAndroidUps `json:"ups,omitempty"`
}

type PushChannelAndroidUps struct {
	Notification *PushChannelAndroidUpsNotification `json:"notification,omitempty"`
	Transmission string                             `json:"transmission,omitempty"`
	Options      map[string]interface{}             `json:"options,omitempty"`
}

type PushChannelAndroidUpsNotification struct {
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	ClickType string `json:"click_type,omitempty"`
	Intent    string `json:"intent,omitempty"`
	Url       string `json:"url,omitempty"`
	Payload   string `json:"payload,omitempty"`
	NotifyId  int    `json:"notify_id,omitempty"`
}
