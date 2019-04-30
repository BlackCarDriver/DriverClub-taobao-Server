package data

//the most basely data of goods that show in homepage goods-list
type Goods1 struct {
	Imgurl  string  `json:"imgurl"`
	Owner   string  `json:"owner"`
	Date    string  `json:"date"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
	GoodsId int64   `json:"goodsid"`
	Gname   string  `json:"gname"`
}

//mean the sub type name and the number of goods belong to it type
type GoodsSubType struct {
	Name string `json:"name"`
	Num  int64  `json:"num"`
}

//the big type of goods
type GoodsType struct {
	TypeName string         `json:"typename"`
	List     []GoodsSubType `json:"list"`
}

//the state of an goods
type GoodsState struct {
	See     int `json:"see"`
	Like    int `json:"like"`
	Talk    int `json:"talk"`
	Collect int `json:"collect"`
}

//goods-data uploads from upload-goods-page
type UploadGoodsData struct {
	Username   string `json:"username"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	Price      int    `json:"price"`
	Imgurl     string `json:"imgurl"`
	Type       string `json:"type"`
	Tag        string `json:"tag"`
	Usenewtag  bool   `json:"usenewtag"`
	Newtagname string `json:"newtagname"`
	Text       string `json:"text"`
}

//商品封面的简单版，显示在个人主页
type SimppleGoodsDate struct {
	Imgurl  string  `json:"imgurl"`
	GoodsId string  `json:"goodsid"`
	Describ string  `json:"describ"`
	Price   float64 `json:"price"`
}

//商品详细页基本数据
var MockData3 = Goods1{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "一个好用不贵的电饭煲，九成新，欲购从速！", 100.3, 123123, "九阳电饭煲"}

//商品详细页商品状态
var MockData4 = GoodsState{See: 112233, Like: 33322, Talk: 100, Collect: 40}

var MockData5 = `<p><u><b style="background-color: rgb(156, 198, 239);">hahahahahah</b></u></p>`

//首页分类
var MockData2 = []GoodsType{
	{"stydy", []GoodsSubType{{"book", 111}, {"ruler", 123}, {"cap", 123}}},
	{"sport", []GoodsSubType{{"ball", 333}, {"ruler", 123}, {"cap", 123}}},
	{"live", []GoodsSubType{{"rice", 11}, {"ruler", 123}, {"cap", 123}}},
	{"electrit", []GoodsSubType{{"phone", 1}, {"ruler", 123}, {"cap", 123}}},
	{"handdiv", []GoodsSubType{{"wool", 1}, {"ruler", 123}, {"cap", 123}}},
	{"virutal", []GoodsSubType{{"link", 111}, {"ruler", 123}, {"cap", 123}}},
	{"other", []GoodsSubType{{"water", 111}, {"ruler", 123}, {"cap", 123}}},
}

//首页封面
var MockData = []Goods1{
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 100.3, 123123, ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 120.3, 1231, ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 140, 120, ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 100.3, 123123, ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 120.3, 1231, ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 140, 120, ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 100.3, 123123, ""},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "BlackCardriver", "2019-2-10", "adsfasdf阿斯顿发生大发大法第三方", 120.3, 1231, ""},
}

var MockData6 = []SimppleGoodsDate{
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "hahahahahhahahahaha", 123.123},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "hahahahahhahahahaha", 0.9},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "hahahahahhahahahaha", 14},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "hahahahahhahahahaha", 123},
	{"http://www.mycodes.net/upload_files/article/162/1_20190319070316_nu1ok.jpg", "1234567", "hahahahahhahahahaha", 13},
}
