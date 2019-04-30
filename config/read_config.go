package config

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"../mylog"
)
//the path of profile
const server_conf string = "./config/conf/server.conf"
//the path of ssl certificate,it can not read from config file!!
var (
	UseHttps = "false"
	Cert = "./config/conf/1_www.blackcardriver.cn_bundle.crt"
	Key  = "./config/conf/2_www.blackcardriver.cn.key"
)
//the config variaible and their default value (read from config file)
var (
	Listen_addr    = "localhost:8080"
	Read_time_out  = 5 * time.Second
	Write_time_out = 5 * time.Second
	Public_addr    = "localhost:8080"
)
// default config of database connectting(read from config file)
var (
	Dbhost     = "localhost"
	Dbport     = 5432
	Dbuser     = "username"
	Dbpassword = "password"
	Dbname     = "databsename"
)
//the default value of config of Email server(read from cinfig file)
var (
	SendEmail = "false"
	Emailhost = "serverhost"
	Emailport = 0
	MyAccount = "emailadress"
	MyPassword = "password"
)
//the filename and the path of those file used by file translate server
var(
	FileRoot = "./source"
	Name_Path = make(map[string]string)
)
//the map to save the config message writen in the file
var confmap = make(map[string]string)

func init() {
	read_server_config()
	set_variaible(confmap)
	list_config()
}

//read the config from file into a map 
func read_server_config() {
	f, err := os.Open(server_conf)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	//begin to read the config file until the end of file
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		// filter annotation
		if len(b)>0 && b[0] == '#' { 
			continue
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 { //can not find '='
			continue
		}
		key := strings.TrimSpace(s[:index])
		value := strings.TrimSpace(s[index+1:])
		if len(key) == 0 || len(value) == 0 {
			continue
		}
		//if it key mean a filename, save if File_Path
		if key[0] == '_' {
			filename := key[1:]
			Name_Path[filename] = value
			continue
		}
		confmap[key] = value
	}
}

//read the config from map and write in variaible
func set_variaible(confmap map[string]string) {
	//if there is value in config file then use new one,else use default
	Listen_addr = notnullval(confmap["Listen_addr"], Listen_addr)
	Public_addr = notnullval(confmap["Public_addr"], Public_addr)
	Dbhost = notnullval(confmap["Dbhost"], Dbhost)
	Dbuser = notnullval(confmap["Dbuser"], Dbuser)
	Dbpassword = notnullval(confmap["Dbpassword"], Dbpassword)
	Dbname = notnullval(confmap["Dbname"], Dbname)
	UseHttps = notnullval(confmap["UseHttps"], UseHttps)
	FileRoot = notnullval(confmap["FileRoot"], FileRoot)
	SendEmail = notnullval(confmap["SendEmail"],SendEmail)
	Emailhost = notnullval(confmap["Emailhost"],Emailhost)
	MyAccount = notnullval(confmap["MyAccount"],MyAccount)
	MyPassword = notnullval(confmap["MyPassword"],MyPassword)
	//if the value is string type then should write in upsise
	//change the config-variable if they have specised in the config file
	//else use the default value
	if confmap["Read_time_out"] != "" {
		newint, err := strconv.Atoi(confmap["Read_time_out"])
		HandleConfigError("strconv readtimeout : ", err)
		Read_time_out = time.Duration(newint) * time.Second
	}
	if confmap["Write_time_out"] != "" {
		newint, err := strconv.Atoi(confmap["Write_time_out"])
		HandleConfigError("strconv write_time_out ", err)
		Write_time_out = time.Duration(newint) * time.Second
	}
	if confmap["Dbport"] != "" {
		newint, err := strconv.Atoi(confmap["Dbport"])
		HandleConfigError("strconv dbport :", err) 
		Dbport = newint
	}
	if confmap["Emailport"] != "" {
		newint, err:= strconv.Atoi(confmap["Emailport"])
		HandleConfigError("strconvEmailport :", err) 
		Emailport = newint
	}
}

//display config
func list_config() {
	mylog.Log("Listen_addr  " + Listen_addr)
	mylog.Log("Public_addr  " + Public_addr)
	mylog.Log("Read_time_out  " + Read_time_out.String())
	mylog.Log("Write_time_out  " + Write_time_out.String())
	mylog.Log("UseHttps "+ UseHttps)
	mylog.Log("DataBaseHost " + Dbhost)
	mylog.Log("DataBasePort " + strconv.Itoa(Dbport))
	mylog.Log("DataBaseUserName " + Dbuser)
	mylog.Log("UsingDatabase " + Dbname)
	mylog.Log("Dbpassword " + Dbpassword)
	mylog.Log("SendEmail " + SendEmail)
	mylog.Log("Emailhost " + Emailhost )
	mylog.Log("Emailport " + strconv.Itoa(Emailport))
	mylog.Log("MyAccount " + MyAccount)
	mylog.Log("FileRoot " + FileRoot)
	mylog.Log("MyPassword " + MyPassword)
}

//equal to  a !="" ? a:b
func notnullval(a string, b string) string {
	if a == "" {
		return b
	}
	return a
}

func HandleConfigError(key string, err error) {
	temp := "Some thing worng in config " + key
	if err != nil {
		mylog.Errorlog.Panicf(temp,err)
	}
}
