package main

import (
	//"errors"
	"net/http"

	"./config"
	"./database"
	"./router"
	"./tools"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", router.Test_connect)
	mux.HandleFunc("/data", router.TestData)
	mux.HandleFunc("/data2", router.TestData2)
	mux.HandleFunc("/getmsg/goods/souce", router.TestData3)
	mux.HandleFunc("/getmsg/usermsg", router.GetUserMsg)
	mux.HandleFunc("/getmsg/othermsg", router.GetUserMsg2)
	mux.HandleFunc("/upload/goods", router.TestData4)
	mux.HandleFunc("/upload/cover", router.GetCover)
	mux.HandleFunc("/upload/headimg", router.PostHeadImg)
	mux.HandleFunc("/source/files", router.GetFiles)
	mux.HandleFunc("/source/images", router.GetImages)
	mux.HandleFunc("/upload/upload/goodsdata", router.GetGoodsData)
	mux.HandleFunc("/updata/mymessage/basemsg", router.UpdataBaseMsg)
	mux.HandleFunc("/updata/mymessage/contactmsg", router.UpdataContactMsg)
	mux.HandleFunc("/getmsg/personal/mymessage", router.GetPersonalMsg)
	mux.HandleFunc("/signin", router.Login)
	mux.HandleFunc("/register/confirmmsg", router.ConfirMsg)
	mux.HandleFunc("/regeister/confirmcode", router.ConfirmCode)
	mux.HandleFunc("/test2", router.Test2)
	//connect to database
	database.Testdb()
	//set mux server
	server := &http.Server{
		Addr:           config.Listen_addr,
		ReadTimeout:    config.Read_time_out,
		WriteTimeout:   config.Write_time_out,
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
	}
	var err error
	if config.UseHttps == "true" {
		err = server.ListenAndServeTLS(config.Cert, config.Key);
	}else {
		err = server.ListenAndServe()
	}
	tools.HandleError("Worng at ListenAndServe,", err, -1)
	return
}
