package data

type Mymessage struct {
	Messageid    string `json:"messageid"`
	Senderimgurl string `json:"senderimgurl"`
	Sendername   string `json:"sendername"`
	Title        string `json:"title"`
	Time         string `json:"time"`
	Describe     string `json:"describ"`
}

//个人中心的我的消息数据
var MockMessageData = []Mymessage{
	{
		"123123",
		"https://gocn.vip/uploads/avatar/000/00/72/71_avatar_max.jpg",
		"BlacCarDriver",
		"欢迎使用过balckcardriver.cn",
		"2019-3-3",
		"https://img.alicdn.com/bao/uploaded/i1/TB1WlfKOFXXXXbJapXXXXXXXXXX_!!0-item_pic.jpg_160x160xz.jpg",
	},
	{
		"123123",
		"https://gocn.vip/uploads/avatar/000/00/72/71_avatar_max.jpg",
		"BlacCarDriver",
		"欢迎使用过balckcardriver.cn",
		"2019-3-3",
		"https://img.alicdn.com/bao/uploaded/i1/TB1WlfKOFXXXXbJapXXXXXXXXXX_!!0-item_pic.jpg_160x160xz.jpg",
	},
}
