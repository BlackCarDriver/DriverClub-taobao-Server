package router

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"../config"
	"../tools"
)

var imgpath string = "./source/images" //the directory path of save and get images

//receive the cover-images of a goods and return the link of see it
func GetCover(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	//parse from and get the base message
	var username = r.PostFormValue("name")
	r.ParseMultipartForm(1002 << 10)
	//check if file is empty
	if len(r.MultipartForm.File["file"]) == 0 {
		return
	}
	//check if the type is image
	files := r.MultipartForm.File["file"][0]
	ext := strings.ToLower(path.Ext(files.Filename))
	if ext != ".jpg" && ext != ".png" {
		return
	}
	//save img to the host
	file, err := files.Open()
	defer file.Close()
	tools.HandleError("file.open", err, -1)
	var id = "0001"                                           //the temply id of goods
	savelocation := imgpath + `/cover/` + username + id + ext //the path of it images after saving
	os.Mkdir(imgpath, os.ModePerm)
	cur, err := os.Create(savelocation)
	defer cur.Close()
	tools.HandleError("os.Create err", err, -1)
	io.Copy(cur, file)
	fmt.Println("Save file :", savelocation)
	//return the imgurl of get it iamges
	imgrul := config.Public_addr + "/source/images?tag=cover&&name=" + username + id + ext
	tools.WriteJson(w, imgrul)
}

//just like GetCover function, but it is ready for headimg of user
//receive the cover-images of a goods and return the link of see it
func PostHeadImg(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	//parse from and get the base message
	var username = r.PostFormValue("name")
	r.ParseMultipartForm(1002 << 10)
	//check if file is empty
	if len(r.MultipartForm.File["file"]) == 0 {
		return
	}
	//check if the type is image
	files := r.MultipartForm.File["file"][0]
	ext := strings.ToLower(path.Ext(files.Filename))
	if ext != ".jpg" && ext != ".png" {
		return
	}
	//save img to the host
	file, err := files.Open()
	defer file.Close()
	tools.HandleError("file.open", err, -1)
	savelocation := imgpath + `/headimg/` + username + ext //the path of it images after saving
	os.Mkdir(imgpath+`/headimg/`, os.ModePerm)
	cur, err := os.Create(savelocation)
	defer cur.Close()
	tools.HandleError("os.Create err", err, -1)
	io.Copy(cur, file)
	fmt.Println("Save file :", savelocation)
	//return the imgurl of get it iamges
	imgrul := config.Public_addr + "/source/images?tag=headimg&&name=" + username + ext
	tools.WriteJson(w, imgrul)
}

//the return the images source by given imgurl
//example of imgurl:http://localhost:8090/source/images?tag=headimg&&name=testcover.jpg
func GetImages(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	tag := vars["tag"]
	name := vars["name"]
	if len(tag) == 0 || len(name) == 0 {
		return
	}
	//find the images and return []byte
	filepath := imgpath + `/` + tag[0] + "/" + name[0]
	temp, err := ioutil.ReadFile(filepath)
	if err != nil {
		tools.HandleError("GetImages readfile error :", err, 1)
		return
	}
	w.Write(temp)
	return
}
