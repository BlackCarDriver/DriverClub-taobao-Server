package router
/*
the function or method of file translation is writen in this file
*/
import(
	"net/http"
	"../config"
	"../tools"
	"io/ioutil"
	"fmt"
)

//user use get method to request an specified file by filename
//if the filename in map then send the file back directly
//example of imgurl:https://localhost:8090/source/files?name=wahtbaoutnow
func GetFiles(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	namemap := vars["name"]
	if len( namemap ) == 0 { //url is worng
		return
	}
	fmt.Println(namemap)
	filename := config.Name_Path[ namemap[0] ]
	if filename == "" {
		fmt.Println("file name do not in the map! ")
		return 
	}
	//find the specied file and return
	filepath := config.FileRoot + filename
	temp, err := ioutil.ReadFile(filepath)
	if tools.HandleError("GetFiles readfile error :", err, 1) {
		tools.WriteJson(w,"Can't not find it file !")
		return
	}
	w.Write(temp)
	return
}