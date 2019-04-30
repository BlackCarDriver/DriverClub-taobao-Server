package mylog

import (
	"fmt"
	"log"
	"os"
)

//the file to recorde some important error
const error_log_path = "./mylog/logfile/error.log"

//the file to recorde some general message
const msg_log_path = "./mylog/logfile/msg.log"

//the file to record some ordinary output which don't care the position
const tag_log_path = "./mylog/logfile/tag.log"

//used to record some import error,and quite the program
var Errorlog *log.Logger

//used to recrod program output,
var Msglog *log.Logger

//record the date,time and output of program
var Taglog *log.Logger

func init() {
	efp, err1 := os.Create(error_log_path)
	mfp, err2 := os.Create(msg_log_path)
	tfp, err3 := os.Create(tag_log_path)
	if err1 != nil {
		log.Fatal("can't not open error_file,", err1)
	}
	if err2 != nil {
		log.Fatal("can't not open  msg.log,", err2)
	}
	if err3 != nil {
		log.Fatal("can't not open tag.log,", err3)
	}
	//if close the *File,the log can't be writen in file!
	Errorlog = log.New(efp, "", 18)
	Msglog = log.New(mfp, "", 18)
	Taglog = log.New(tfp, "", 3)
}

func Log(any ...interface{}) {
	for _,y := range any{
		fmt.Print(y)
		Taglog.Print(y)
	}
	fmt.Println("")
	Taglog.Println("")
}
func Println(any ...interface{}){
	for _,y := range any{
		fmt.Print(y)
		Msglog.Print(y)
	}
	fmt.Println("")
	Msglog.Println("")
}
/*
//if you write the method like the folowing links do,
//it would not record the position of error !
func Msg_writer(prefix string, err error) {
	Msglog.Println(prefix, err)
}
*/
