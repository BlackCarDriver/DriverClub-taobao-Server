package tools
/*
the function related to take and set cookie in writen in thsi file
*/
import (
	"net/http"
	"../config"
	"fmt"
	"net"
)
//create an particul cookie
func MakeCookie(key string, value string, time int)(ck http.Cookie) {
	ck = http.Cookie{
        Name: key,
		Value: value,
		MaxAge: time,			
		HttpOnly:true,
		Secure:(config.UseHttps=="true"),
	}
	return ck
}
//set an randan string on uesr cookie as token and return it string
func SetVtfCookie(w http.ResponseWriter,)string{
	randkey := CreateRandString(18)
	ck := MakeCookie("carkey", randkey, 300)	//time
	http.SetCookie(w, &ck)
	return randkey
}

//reade cookie from user
func GetCookie(req *http.Request)string{
	ck, err := req.Cookie("carkey")
    if err != nil {
	   fmt.Println(err)
	   return ""
	}
	return ck.Value
}

//get client ip from request header and return as string
func GetIp(r *http.Request)string{
	remoteAddr := r.RemoteAddr
	XForwardedFor := "X-Forwarded-For"
    XRealIP       := "X-Real-IP"
    if ip := r.Header.Get(XRealIP); ip != "" {
        remoteAddr = ip
    } else if ip = r.Header.Get(XForwardedFor); ip != "" {
        remoteAddr = ip
    } else {
        remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
    }
    if remoteAddr == "::1" {
        remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}
//check if the user forbid the browser save the cookie, return flase if 
//can not save the cookie on user boswer
func TestCookie(w http.ResponseWriter,r *http.Request)bool{
	ck := MakeCookie("drivertest","hellow,I am BlackCarDriver!",10)
	http.SetCookie(w, &ck)
	_, err := r.Cookie("drivertest")
	return (err != nil)
}