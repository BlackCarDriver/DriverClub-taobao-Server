//function related to sigin in and register write in this file
package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../data"
	"../database"
	"../tools"
	"../mylog"
	"../config"
)


//check the message that user want to register, if every thing ok then 
//should create an comfirm code and send it code to user
func ConfirMsg(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	var account data.Account1
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &account)
	if tools.HandleError("unmarshal in ConfirMsg ", err, 1) {
		tools.WriteJson(w, unknowerr)
		return
	}
	//comfirm the register message again
	if tools.CheckFormat(account.Name, account.Email, account.Password) == false {
		mylog.Log("the register message format is worng when ConfrimMsg !")
		tools.WriteJson(w, unknowerr)
		return
	}
	//check if the name or email is already exist, 
	res := database.CheckExist(account.Name, account.Email)
	if res != enable {
		fmt.Println(" check exist enable ,res = ",res)
		//return result include: repectname=-20 , repectemail=-30,  worng=-1,  enable=2 
		tools.WriteJson(w, res)
		return
	}
	//after above check then create a comfirm and send to the user by email
	comfirmcode := database.CreateComfirmCode(account.Name, account.Email, account.Password)
	if comfirmcode == "" { //sql exec fall or 0 row effect
		tools.WriteJson(w, unknowerr)
		return
	}
	account.Code = comfirmcode
	mylog.Println( "create comfirm code scuess, code is : ", comfirmcode)
	//dont need to send email when in testing
	if config.SendEmail != "true" {
		tools.WriteJson(w,enable)
		return
	}
	userindex := database.CountUser()
	mylog.Println("count account result is :", userindex)
	//send the comfrim code to the user email ,return othererror or scuess
	if tools.SendConfrimEmail(account, userindex) != scuess{
		tools.WriteJson(w, othererror)
		return
	}
	tools.WriteJson(w,enable)
	return
}



//after send an email to user and now user send the confirm code back
//we should save the account in the database after check the confrim before
func ConfirmCode(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	var data data.Account1
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)
	if tools.HandleError("unmarshal in ConfirmCode ", err, 1) {
		tools.WriteJson(w, unknowerr)
		return
	}
	if tools.CheckFormat(data.Name, data.Email, data.Password) == false {
		mylog.Log("the register message format is worng when cofirmcode !")
		tools.WriteJson(w, unknowerr)
		return
	}
	res := database.CompareCode(data.Name, data.Email, data.Password, data.Code)
	if res == disable || res == worng {	//the code is worng or try too much times
		tools.WriteJson(w, res)
		return
	}
	//if the comfim code is right, then create account , it will autoly delete the comfrim in database
	result := database.CreateAccount(data.Name, data.Email, data.Password)
	tools.WriteJson(w, result)	//scuess or worng or othererror
}

//test the function of database
func Test1(w http.ResponseWriter, r *http.Request){
	fmt.Println("test1 have been call !")
	res := database.CreateComfirmCode("blackcardriver","teswtats@123.com","123123")
	fmt.Println("the result is :", res)
	tools.WriteJson(w, res)	
	return
}

func Test2(w http.ResponseWriter, r *http.Request){
	fmt.Println("test2 have been call !")
	res := database.CompareCode("blackcardriver","teswtats@123.com","123123","746501")
	fmt.Println("the result is :", res)
	tools.WriteJson(w, res)
	return
}