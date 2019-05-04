package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../data"
	"../mylog"
	"../tools"
)
//the state return to the user
var (
	worng     = -1
	scuess    = 1
	enable    = 2
	disable   = -2
	unknowerr = -3
	othererror  = -99
	unsafe 	= -999
)
var goodstext string = "undefine"

func Test_connect(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	tools.WriteJson(w, "version :5-1-6-45")
}

//send back goods message
func TestData(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	tools.WriteJson(w, data.MockData)
}

func Usershort(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	tools.WriteJson(w, data.Shortmsg)
}


//send back type message of goods
func TestData2(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	tools.WriteJson(w, data.MockData2)
}

//return data of goods detail page
func TestData3(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	tag := vars["tag"][0]
	// id := vars["goodsid"][0]
	switch tag {
	case "base":
		tools.WriteJson(w, data.MockData3)
	case "state":
		tools.WriteJson(w, data.MockData4)
	case "text":
		tools.WriteJson(w, goodstext)
	default:
		return
	}
	return
}

//receive the goods-descriobe-text user upload
func TestData4(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	var postbody map[string]string
	body, err := ioutil.ReadAll(r.Body)
	tools.HandleError("ioutil.readall error", err, 1)
	json.Unmarshal(body, &postbody)
	if postbody["goodsdata"] != "" {
		goodstext = postbody["goodsdata"]
		fmt.Println("goodstext = ", goodstext)
		tools.WriteJson(w, "scuess!")
	} else {
		tools.WriteJson(w, "fall")
	}

	return
}

//receive the goods-data from upload-goods-page
func GetGoodsData(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	goodsdata := data.UploadGoodsData{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &goodsdata)
	if err != nil {
		mylog.Errorlog.Println(err)
		tools.WriteJson(w, -1)
	}
	fmt.Println(goodsdata)
	tools.WriteJson(w, 1)
}

//write back the usermessage accordding to the given user id
func GetUserMsg(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	id := vars["id"]
	if len(id) != 1 {
		return
	}
	tools.WriteJson(w, data.Mockpersondata)
	return
}

//write back message of other peopel
func GetUserMsg2(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	id := vars["id"]
	if len(id) != 1 {
		return
	}
	tools.WriteJson(w, data.MorkMydata)
	return
}

//updata basemsg of user
func UpdataBaseMsg(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	basemsg := data.Usermsg{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &basemsg)
	if err != nil {
		mylog.Errorlog.Println(err)
		tools.WriteJson(w, -1)
	}
	fmt.Println(basemsg)
	tools.WriteJson(w, 1)
}

//updata contact message of user
func UpdataContactMsg(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	basemsg := data.Usermsg{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &basemsg)
	if err != nil {
		mylog.Errorlog.Println(err)
		tools.WriteJson(w, -1)
	}
	fmt.Println(basemsg)
	tools.WriteJson(w, 1)
}

//write back the mesage personal-page need
func GetPersonalMsg(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	var postbody map[string]string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tools.HandleError("ioutil.ReadAll(r.Body) ", err, 1)
		tools.WriteJson(w, -1)
		return
	}
	json.Unmarshal(body, &postbody)
	tag := postbody["tag"]
	if tag == "" {
		tools.WriteJson(w, -1)
		return
	}
	switch tag {
	case "mymsg":
		tools.WriteJson(w, data.MorkMydata)
	case "mygoods":
		tools.WriteJson(w, data.MockData6)
	case "mycollect":
		tools.WriteJson(w, data.MockData6)
	case "message":
		tools.WriteJson(w, data.MockMessageData)
	case "rank":
		tools.WriteJson(w, data.MorkRankData)
	case "care":
		tools.WriteJson(w, data.MorkUserData)
	default:
		tools.WriteJson(w, 0)
		return
	}
	return
}
