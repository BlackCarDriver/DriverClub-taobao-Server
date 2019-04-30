package database
/*
some simple function of databse query should writen in it file
*/
import (
	"../tools"
)

//query the total number of user in databse
//sql commant: select count(id) from account; or : select max(id) from account;
func CountUser()(num int){
	commant := 	`select count(id) from account`
	err := db.QueryRow(commant).Scan(&num)
	if tools.HandleError("database->countuser ",err,1) {
		return -1
	}
	return num
}