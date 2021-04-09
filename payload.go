package jpush

import (
	"encoding/json"
)

type PayLoad struct {
	Platform     interface{} `json:"platform"`               //推送平台设置
	Audience     interface{} `json:"audience"`               //推送设备指定
	Notification interface{} `json:"notification,omitempty"` //通知内容体。是被推送到客户端的内容。与 message 一起二者必须有其一，可以二者并存
	Message      interface{} `json:"message,omitempty"`      //消息内容体。是被推送到客户端的内容。与 notification 一起二者必须有其一，可以二者并存
	Options      *Option     `json:"options,omitempty"`      //可选参数
}

func NewPushPayLoad() *PayLoad {
	pl := &PayLoad{}
	o := &Option{}
	o.ApnsProduction = false
	pl.Options = o
	return pl
}

func (this *PayLoad) SetPlatform(pf *Platform) {
	this.Platform = pf.Os
}

func (this *PayLoad) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *PayLoad) SetOptions(o *Option) {
	this.Options = o
}

func (this *PayLoad) SetMessage(m *Message) {
	this.Message = m
}

func (this *PayLoad) SetNotice(notice *Notice) {
	this.Notification = notice
}

func (this *PayLoad) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
