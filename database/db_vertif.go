package database

import(
	"../mylog"
	"../tools"
)
//KeyMap and IpMap is the data to vertify the of message of user
//the code reference is writen in it file
var KeyMap = make(map[string]string)
var IpMap = make(map[string]string)

//copy the data from database to KeyMap and IpMap,return rows
func GetDataFromDb(){
	res := GetRandstr(KeyMap)
	if res >= 0 {
		mylog.Log("Take keymap form database to map scuess! row number is ", res)
	}else{
		mylog.Log("Take keymap fall !")
	}
	res2 := GetIpMap(IpMap)
	if res2 >= 0 {
		mylog.Log("Take IPmap form database to map scuess! row number is ", res2)
	}else{
		mylog.Log("Take IPmap fall !")
	}
}
//add an recode into keymap and ipmap
func InsertMap(name, key, ip string){
	KeyMap[name] = key
	IpMap[name] = ip
}
//save the keyMap and Ipmap to teh database,return the effect rows
//templace:insert into vtf_rk(uname,rk)values('asdf','adsfasdfasdf');
//templace:update vtf_rk set rk = 'hahahahaha' where uname = 'asdf';
func SaveDataToDb(){
	commant1 := `insert into vtf_rk(uname,rk)values($1,$2)`
	commant2 := `update vtf_rk set rk = $2 where uname = $1`
	commant3 := `insert into vtf_ip(uname,ip)values($1,$2)`
	commant4 := `update vtf_ip set ip = $1 where uname = $2`
	errnum := 0
	for name,rk := range KeyMap {
		_,err := db.Exec(commant1,name,rk)	//insert
		//if cant not insert, mean there already have an colum with same name
		if err!=nil {	
			_,err2 := db.Exec(commant2,name,rk)
			if err2!=nil {
				mylog.Log("vtf_rk can't insert and updata ! : ",err)
				errnum ++
			}
		}
	}
	for name,ip := range IpMap {
		_,err := db.Exec(commant3,name,ip)
		if err!=nil {
			_,err2 := db.Exec(commant4,name,ip)
			if err2!=nil {
				mylog.Log("vtf_ip can't insert and updata ! : ",err)
				errnum++
			}
		}
	}
	if errnum == 0{
		mylog.Log("SaveData form map to databaes sucess!")
	}else{
		mylog.Log(errnum ," worng happen when save data from map to database!")
	}
}

//read the data from database to map, return effect rows number
//templace : select uname,ip from vtf_rk;
func GetRandstr( keymap map[string]string )(num int){
	commant := `select uname,rk from vtf_rk`
	rows, err := db.Query(commant)
	if tools.HandleError("read table vtf_rk fall :" , err, 1) {
		return othererror
	}
	uname,rk := "",""
	for rows.Next() {
		err = rows.Scan(&uname, &rk)
		if err!=nil {
			mylog.Println("row.Scan error at GetRandstr :",err)
			continue
		}
		keymap[uname] = rk
		num ++
	}
	return num
}

//read the data from database to map, return effect rows number
//templace : select uname,ip from vtf_ip;
func GetIpMap( ipmap map[string]string )(num int){
	commant := `select uname,ip from vtf_ip`
	rows, err := db.Query(commant)
	if tools.HandleError("read table vtf_ip fall :" , err, 1) {
		return othererror
	}
	uname,ip := "",""
	for rows.Next() {
		err = rows.Scan(&uname, &ip)
		if err!=nil {
			mylog.Println("row.Scan error at GetRandstr :",err)
			continue
		}
		ipmap[uname] = ip
		num ++
	}
	return num
}