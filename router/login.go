package router

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"../data"
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
	res := database.CheckLoginPassword(account)
	fmt.Println(res)
	tools.WriteJson(w, res)
}