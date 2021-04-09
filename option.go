package jpush

type Option struct {
	SendNo            int   `json:"sendno,omitempty"`              //推送序号
	TimeLive          int   `json:"time_to_live,omitempty"`        //离线消息保留时长(秒)
	ApnsProduction    bool  `json:"apns_production"`               //APNs 是否生产环境
	OverrideMsgId     int64 `json:"override_msg_id,omitempty"`     //要覆盖的消息 ID
	BigPushDuration   int   `json:"big_push_duration,omitempty"`   //定速推送时长(分钟)
	ApnsCollapseId    int   `json:"apns_collapse_id,omitempty"`    //更新 iOS 通知的标识符
	ThirdPartyChannel int   `json:"third_party_channel,omitempty"` //推送请求下发通道
}

func (this *Option) SetSendno(no int) {
	this.SendNo = no
}

func (this *Option) SetTimelive(timelive int) {
	this.TimeLive = timelive
}

func (this *Option) SetOverrideMsgId(id int64) {
	this.OverrideMsgId = id
}

func (this *Option) SetApns(apns bool) {
	this.ApnsProduction = apns
}

func (this *Option) SetBigPushDuration(bigPushDuration int) {
	this.BigPushDuration = bigPushDuration
}
