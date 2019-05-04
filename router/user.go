package router
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"../database"
	"../mylog"
	"../tools"
	"fmt"
)

//return the short message of user show in homepage 
func UserShortMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Println("usershortmsg")
	tools. SetHeader2(w,r)
	if r.Method != "POST" {
		return
	}
	postbody := make( map[string]string  )
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &postbody)
	if tools.HandleError("Unmarshal postbody fall:",err,1) {
		return
	}
	postname := postbody["name"]
	fmt.Println(tools.GetCookie(r))
	if safe:= vertify(postname, r); safe == 0 {
		mylog.Log("UserShortMsg() receive a unsafe require, username is : "+postname)
		tools.WriteJson(w, unsafe)
		return
	}
	fmt.Println("Comfrim user scuess! state !")
	data , res:= database.GetUserShortMsg(postname)
	if res == scuess{
		tools.WriteJson(w,data)
		return
	}
	return
}

