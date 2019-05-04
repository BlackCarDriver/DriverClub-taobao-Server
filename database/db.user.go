package database
import(
	"../data"
	"../tools"
)

//find and return the userhosrt message with specied name
//select * from short_msg where uname = 'black';
func GetUserShortMsg(name string)(data.UserShort ,  int){
	commant := `select * from short_msg where uname = $1 ;`
	row := db.QueryRow(commant, name)
	var us  data.UserShort
	lt, now, t := "", "",""
	err := row.Scan( &t , &us.Grade, &us.Score, &us.Goods, &us.Message, &lt, &now, &us.Imgurl)
	if tools.HandleError("Scan GetUserShortMsg error :",err,1) {
		return us,othererror
	}
	us.Lastime = tools.CountTimeLap(lt, now)
	us.Imgurl = tools.CreateImgUrl(us.Imgurl)
	return  us, scuess
}

//updata lasttime in usermsg2
//templace :update usermsg2 set lasttime = now() where userid = (select id from account where uname = 'black');
func UpdateLasttime(uname string){
	commant:=`update usermsg2 set lasttime = now() where userid = (select id from account where uname = $1);`
	db.Exec(commant, uname)
}