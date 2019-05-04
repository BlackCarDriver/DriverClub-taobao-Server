package database

import(
	"../tools"
	"../mylog"
	"../data"
)
/*
the function of user Login writen in this file
*/

//check the password,when user login with password and username
//templace: select count(id) from account where uname='black' and upassword='123456';
func CheckLoginPassword(account data.Account2) int{
	commant := `select count(id) from account where uname=$1 and upassword=$2`
	row := db.QueryRow(commant, account.Name , account.Password)
	rownum := 0
	err := row.Scan(&rownum)
	if tools.HandleError("Database->CheckLoginPasswrod row.scan ",err,1) {
		return othererror
	}
	if rownum > 1 {	//the account with same name and same password more thant one 
		mylog.Log("Login data maybe worng with tow same account ! username is " + account.Name)
		return othererror
	}
	if rownum == 0 {
		return disable
	}
	return enable
} 

//updata the login times of an account