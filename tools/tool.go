package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"regexp"
	"time"

	"../mylog"
)
//the meaning of return value
var (
	worng       = -1
	scuess      = 1
	enable      = 2
	disable     = -2
	repectname  = -20
	repectemail = -30
	othererror  = -99
)
var r *rand.Rand

func init(){
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

//write json code to responsewriter
func WriteJson(w http.ResponseWriter, data interface{}) {
	jsondata, err := json.Marshal(data)
	HandleError("worng at tool.go writejson :", err, 1)
	w.Write(jsondata)
}

//simplely set ResponseWriter
func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
}

//the simple function to handle the err,use avariable to contral
//exist or containue when error happend
//return 1 if the error is not null
func HandleError(tag string, err error, method int) bool {
	if err != nil {
		fmt.Println(err)
		switch method {
		case 1:
			mylog.Errorlog.Println(tag, err)
		case -1:
			mylog.Errorlog.Fatal(tag, err)
		}
		return true
	}
	return false
}

//display the content in http.Request
func ShowPostData(r *http.Request) {
	var postbody map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	HandleError("showPostData readall() :", err, -1)
	json.Unmarshal(body, &postbody)
	for key, val := range postbody {
		fmt.Println(key, "  -->  ", val)
	}
}

//create an passphrase with length is six
func MakePassphrase() string {
	rand.Seed(time.Now().UnixNano())
	r := 100000 + rand.Intn(900000)
	return strconv.Itoa(r)
}

//translate the age() strign that database return,scush as 00:38:00.04091 into second
func ParseTimeLap(lap string) int {
	regexp2 := regexp.MustCompile(`\d{2}`)
	regexp  := regexp.MustCompile(`^(-)?\d{2}:\d{2}:\d{2}.\d*$`)
	if regexp.MatchString(lap) == false {
		mylog.Errorlog.Println("Format unright in  ParseTimeLap !!")
		return worng
	}else{
		subint := regexp2.FindAllString(lap, 3)
		hts,_ := strconv.Atoi(subint[0])
		mts,_ := strconv.Atoi(subint[1])
		sts,_ := strconv.Atoi(subint[2])
		totaltime := hts *3600 + mts *60 + sts
		if totaltime > 1800 {	//time lap bigger than 30 minute
			return disable
		}
	}
	return enable
}    
//when user register and send new account to server, should check the fromat again
func CheckFormat(name,email,password string) bool{
	namereg := regexp.MustCompile(`^[\x{4E00}-\x{9FA5}_a-zA-Z0-9]{2,15}$`)	//\u4e00-\u9fa5 
	emailreg := regexp.MustCompile(`^\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}$`)
	passwordreg := regexp.MustCompile(`^[a-zA-Z._0-9]{6,20}$`)
	nameres := namereg.MatchString(name)
	emailres := emailreg.MatchString(email)
	passwordres := passwordreg.MatchString(password)
	return(nameres && emailres && passwordres)
}

//create rande string 
func CreateRandString(len int) string{
	bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        b := r.Intn(26) + 97
        bytes[i] = byte(b)
	}
    return string(bytes)
}