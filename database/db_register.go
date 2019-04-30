package database

/*
when s mre user reginter or login, we should check the message if correct or not
or check whether there already have an same uesrname in the database, the reference
code is writen in it file
*/
import (
	"fmt"

	"../mylog"
	"../tools"
)

//check the name and email whether have been used when have an new register
//example:  select count(id) from account where uname = 'blackcardriver2';
 //exalple2: select count(id) from account where email = '234234';
func CheckExist(name string, email string) int {
	commant := `select count(id) from account where uname = $1`
	commant2 := `select count(id) from account where email = $1`
	row := db.QueryRow(commant, name)
	row2 := db.QueryRow(commant2,email)
	namenum, emailnum := 0, 0
	row.Scan(&namenum)
	row2.Scan(&emailnum)
	fmt.Println(namenum, emailnum)
	if namenum != 0 {
		return repectname
	} else if emailnum != 0 {
		return repectemail
	}
	fmt.Println("check exist is ok ! return enable")
	return enable
}

//insert a new account data to database after checke the message and passpharse
//templace: insert into account (uname,email,upassword) values('blackcardriver','13169346002','123456');
func CreateAccount(name, email, password string) int {
	commant := ` insert into account (uname,email,upassword) values($1,$2,$3)`
	res, err := db.Exec(commant, name, email, password)
	if tools.HandleError("databaes->CreateAccount->Exec ", err, 1) {
		return othererror
	}
	rownum, _ := res.RowsAffected()
	if rownum == 0 {
		mylog.Errorlog.Println("Create Account Error! res.RowsAffected() is 0 ! name,email,password is:", name, email, password)
		return worng
	}
	fmt.Println("Create account scuess ", name)	//$$$$$$$$$$$$$$
	return scuess
}

//create an comfirm code message by given register message and return the comfrim
//if return an null string, then some thing worng may happen
//templace :insert into regcode(npe,code,stime)values('blackdriver@131313@111.com@123123','122333',CURRENT_TIMESTAMP);
func CreateComfirmCode(name, email, password string) string{
	nep := name + "@" + email + "@" + password
	code := tools.MakePassphrase()
	commant := `insert into regcode(npe,code,stime)values($1,$2,CURRENT_TIMESTAMP)`
	res, err := db.Exec(commant, nep, code)
	if tools.HandleError("databaes->CreateComfirmCode->Exec ", err, 1) {
		return ""
	}
	rownum, _ := res.RowsAffected()
	if rownum == 0 {
		mylog.Errorlog.Println("Insert into regcode Error! name,email,password is:", name, email, password)
		return ""
	}
	fmt.Println("Create comfirm code ", code)	//$$$$$$$$$$$$$$
	return code
}

//take the comfrim code from database and compare the given from user, anlelyse the time lap is also needed
//templace: select code, age( stime, now() ) from regcode where npe = 'blackdrive...' and state < 10 order by stime desc limit 1;
func CompareCode(name, email, password, code string) int {
	nep :=  name + "@" + email + "@" + password
	commant := `select code, age( stime, now() ) from regcode where npe = $1 and state < 10 order by stime desc limit 1`
	commant2 := `update regcode set state = state + 1 where npe = $1`
	row := db.QueryRow(commant, nep)
	ocode,lap := "","";
	err := row.Scan(&ocode, &lap)
	//no record in database
	if tools.HandleError("sql->CompareCode :",err,1) {
		return worng
	}
	//let the state  in regcode table +1
	db.Exec(commant2,nep)
	//time over or the format is unright or the code is worng
	if tools.ParseTimeLap(lap) == disable || ocode!=code {
		return disable
	}
	return enable
}