package router

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"../data"
	"../mylog"
	"../database"
	"../tools"
)



//user login, check the user name and password return state or token
func Login(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	var account data.Account2
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &account)
	if tools.HandleError("unmarshal in singin ", err, 1) {
		tools.WriteJson(w, unknowerr)
		return
	}
	//return result include othererror,disable,enable
	res := database.CheckLoginPassword(account)
	if res == enable {
		createVtify(account.Name,w,r)
	}
	//reminder user if his broswer can not save the cookie
	if  tools.TestCookie(w,r) == false {
		tools.WriteJson(w,worng)
	}
	tools.WriteJson(w, res)
}

//after an user login, we should save an randan string as token in 
//user cookie , and also recode the it and user ip in map, 
func createVtify(username string, w http.ResponseWriter, r *http.Request){
	userkey := tools.SetVtfCookie(w)
	userip := tools.GetIp(r)
	database.InsertMap(username, userkey, userip)
	mylog.Println("Create an vertify on user ",username," key is ",userkey," ip is ",userip)
}

//evaluate the security of message that user loging,both the ip and key of user
//are same to the  map will return 2, if just the key right return 1
func vertify(username string, r *http.Request)int {
	mrk := database.KeyMap[username] 
	mip := database.IpMap[username]
	if mrk == "" || mip =="" {
		fmt.Println("both is clear")
		return 0	//need to login 
	}
	userip := tools.GetIp(r)
	userrk := tools.GetCookie(r) 
	//fmt.Println(userip , "   ", userrk)
	if userrk == mrk && userip == mip {
		return 2	//it user can be certain
	}
	if userrk ==mrk{
		return 1	//the ip have change
	}
	return 0	//need to login again
}
