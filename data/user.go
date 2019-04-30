package data

//used in set personal message page
type Usermsg struct {
	Username string `json:"username"`
	Userid   string `json:"userid"`
	UserSex  string `json:"usersex"`
	Sign     string `json:"sign"`
	Grade    string `json:"grade"`
	Colleage string `json:"colleage"`
	Email    string `json:"email"`
	Qq       string `json:"qq"`
	Phone    string `json:"phone"`
}

//mainly data in personal page
type PersonalExpend struct {
	Username string `json:"username"`
	Userid   string `json:"userid"`
	UserSex  string `json:"usersex"`
	Sign     string `json:"sign"`
	Grade    string `json:"grade"`
	Colleage string `json:"colleage"`
	Email    string `json:"email"`
	Qq       string `json:"qq"`
	Phone    string `json:"phone"`
	Leave    int64  `json:"leave"`
	Sorce    int64  `json:"sorce"`
	Rank     int64  `json:"rank"`
	Becare   int64  `json:"becare"`
	Like     int64  `json:"like"`
	Lasttime int64  `json:"lasttime"`
	Goodsnum int64  `json:"goodsnum"`
	Scuess   int64  `json:"scuess"`
	Visnum   int64  `json:"visnum"`
	Care     int64  `json:"care"`
}

//user Rank data in personal page
type Rank struct {
	Rank   int64  `json:"rank"`
	Name   string `json:"name"`
	Userid string `json:"userid"`
}

//the data of user care me and i am care
type User struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Imgurl   string `json:"imgurl"`
}

//the struct to describe account that signin
type Account2 struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//the data to describe the account when register
type Account1 struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

var Mockpersondata = Usermsg{
	"BlackCarDriver",
	"1234567",
	"BOY",
	"Welcome to BlackCaDriver.cn!!!",
	"2017",
	"计算机学院",
	"13169346002@163.com",
	"1731328065",
	"15626195662",
}

var MorkMydata = PersonalExpend{
	"BlackCarDriver",
	"1234567",
	"BOY",
	"Welcome to BlackCaDriver.cn!!!",
	"2017",
	"计算机学院",
	"13169346002@163.com",
	"1731328065",
	"15626195662",
	999, 888, 1, 12, 88,
	200, 10, 10, 3000, 20,
}
var MorkRankData = []Rank{
	{1, "BlackCardriver", "123123"},
	{2, "BlackCardriver", "123123"},
	{3, "BlackCardriver", "123123"},
	{4, "BlackCardriver", "123123"},
	{5, "BlackCardriver", "123123"},
	{6, "BlackCardriver", "123123"},
	{7, "BlackCardriver", "123123"},
	{8, "BlackCardriver", "123123"},
	{9, "BlackCardriver", "123123"},
}
var MorkUserData = [2][]User{
	{
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/ed06777a6962696e3b00"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/ed06777a6962696e3b00"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/ed06777a6962696e3b00"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/ed06777a6962696e3b00"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/ed06777a6962696e3b00"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/ed06777a6962696e3b00"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/ed06777a6962696e3b00"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/ed06777a6962696e3b00"},
	},
	{
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/d0017a686967616e6777323030366404"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/d0017a686967616e6777323030366404"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/d0017a686967616e6777323030366404"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/d0017a686967616e6777323030366404"},
		{"1234556", "Blackcardirver", "http://tb.himg.baidu.com/sys/portrait/item/d0017a686967616e6777323030366404"},
	},
}
