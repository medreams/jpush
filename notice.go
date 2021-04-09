package jpush

type Notice struct {
	Alert    string          `json:"alert,omitempty"`    //通知的内容
	Android  *AndroidNotice  `json:"android,omitempty"`  //Android 平台上的通知
	IOS      *IOSNotice      `json:"ios,omitempty"`      //iOS 平台上 APNs 通知结构
	QuickApp *QuickAppNotice `json:"quickapp,omitempty"` //快应用平台上通知结构
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"` //Windows Phone 平台上的通知
}

type AndroidNotice struct {
	Alert             string                 `json:"alert"`                        //通知内容
	Title             string                 `json:"title,omitempty"`              //通知标题
	BuilderId         int                    `json:"builder_id,omitempty"`         //通知栏样式 ID
	ChannelID         string                 `json:"channel_id,omitempty"`         //android通知channel_id
	Priority          int                    `json:"priority,omitempty"`           //通知栏展示优先级
	Category          string                 `json:"category,omitempty"`           //通知栏条目过滤或排序
	Style             int                    `json:"style,omitempty"`              //通知栏展示优先级
	AlertType         int                    `json:"alert_type,omitempty"`         //通知提醒方式
	BigText           string                 `json:"big_text,omitempty"`           //大文本通知栏样式,当 style = 1 时可用
	Inbox             map[string]interface{} `json:"index,omitempty"`              //文本条目通知栏样式,当 style = 2 时可用 厂商inbox
	BigPicPath        string                 `json:"big_pic_path,omitempty"`       //大图片通知栏样式,当 style = 3 时可用
	Extras            map[string]interface{} `json:"extras,omitempty"`             //通知栏样式类型
	LargeIcon         string                 `json:"large_icon,omitempty"`         //通知栏大图标,图标路径可以是以http或https开头的网络图片
	SmallIconURL      string                 `json:"small_icon_uri,omitempty"`     //通知栏小图标,图标路径可以是以http或https开头的网络图片
	Intent            map[string]interface{} `json:"intent,omitempty"`             //指定跳转页面
	UriActivity       string                 `json:"uri_activity,omitempty"`       //指定跳转页面
	UriAction         string                 `json:"uri_action,omitempty"`         //指定跳转页面
	BadgeAddNum       int                    `json:"badge_add_num,omitempty"`      //角标数字，取值范围1-99
	BadgeClass        string                 `json:"badge_class,omitempty"`        //桌面图标对应的应用入口Activity类
	Sound             string                 `json:"sound,omitempty"`              //填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀
	ShowBeginTime     string                 `json:"show_begin_time,omitempty"`    //定时展示开始时间（yyyy-MM-dd HH:mm:ss）
	ShowEndTime       string                 `json:"show_end_time,omitempty"`      //定时展示结束时间（yyyy-MM-dd HH:mm:ss）
	DisplayForeground string                 `json:"display_foreground,omitempty"` //APP在前台，通知是否展示 值为 "1" 时，APP 在前台会弹出通知栏消息；值为 "0" 时，APP 在前台不会弹出通知栏消息。注：默认情况下 APP 在前台会弹出通知栏消息。

}

type IOSNotice struct {
	Alert            interface{}            `json:"alert"`                       //通知内容
	Sound            interface{}            `json:"sound,omitempty"`             //通知提示声音或警告通知
	Badge            string                 `json:"badge,omitempty"`             //应用角标
	ContentAvailable bool                   `json:"content-available,omitempty"` //推送唤醒
	MutableContent   bool                   `json:"mutable-content,omitempty"`   //通知扩展
	Category         string                 `json:"category,omitempty"`          //IOS 8 才支持。设置 APNs payload 中的 "category" 字段值
	Extras           map[string]interface{} `json:"extras,omitempty"`            //附加字段
	ThreadId         string                 `json:"thread-id,omitempty"`         //通知分组
}

type QuickAppNotice struct {
	Alert  string                 `json:"alert"`            //通知内容
	Title  string                 `json:"title,omitempty"`  //通知标题
	Page   string                 `json:"page,omitempty"`   //快应用通知跳转地址。
	Extras map[string]interface{} `json:"extras,omitempty"` //扩展字段
}

type WinPhoneNotice struct {
	Alert    string                 `json:"alert"`                //通知内容
	Title    string                 `json:"title,omitempty"`      //通知标题
	OpenPage string                 `json:"_open_page,omitempty"` //点击打开的页面名称
	Extras   map[string]interface{} `json:"extras,omitempty"`     //扩展字段
}

func (this *Notice) SetAlert(alert string) {
	this.Alert = alert
}

func (this *Notice) SetAndroidNotice(n *AndroidNotice) {
	this.Android = n
}

func (this *Notice) SetIOSNotice(n *IOSNotice) {
	this.IOS = n
}

func (this *Notice) SetQuickAppNotice(n *QuickAppNotice) {
	this.QuickApp = n
}

func (this *Notice) SetWinPhoneNotice(n *WinPhoneNotice) {
	this.WINPhone = n
}
