package router
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"../database"
	"../tools"
)

//return the short message of user show in homepage 
func UserShortMsg(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
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
	data , res:= database.GetUserShortMsg(postname)
	if res == scuess{
		tools.WriteJson(w,data)
		return
	}
	return
}