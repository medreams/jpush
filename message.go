package jpush

type Message struct {
	Content     string                 `json:"msg_content"`            //消息内容本身
	Title       string                 `json:"title,omitempty"`        //消息标题
	ContentType string                 `json:"content_type,omitempty"` //消息内容类型
	Extras      map[string]interface{} `json:"extras,omitempty"`       //JSON 格式的可选参数
}

func (this *Message) SetContent(c string) {
	this.Content = c

}

func (this *Message) SetTitle(title string) {
	this.Title = title
}

func (this *Message) SetContentType(t string) {
	this.ContentType = t
}

func (this *Message) AddExtras(key string, value interface{}) {
	if this.Extras == nil {
		this.Extras = make(map[string]interface{})
	}
	this.Extras[key] = value
}
