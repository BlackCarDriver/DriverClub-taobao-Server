package tools

import (
    "crypto/tls"
    "fmt"
    "log"
    "net"
    "net/smtp"
    "../mylog"
    "../config"
    "../data"
    //"../data" account data.Account2
)
//config variable can read from config file
var (
    serverhost = "smtp.exmail.qq.com"
    port = 465
    myemail = "youaccount"
    mypassword = "yourpassword"
)

//the templace of the email
var confrimBody = `
<div style="background-color:#4caf50;width: 400px;height30200px;padding: 10px;border-radius: 6px;font-weight: 500;margin: 20px;">
🚂  🚃  🚄  🚅  🚆  🚇  🚈  🚉  🚊  🚝  🚞  🚋 🚲 🚜<br>
你好！欢迎并感谢你即将成为本站的第<span style=" color: #E91E63; font-weight: 600;"> %d</span>个成员。<br>
本站仍在不断开发并更新,努力为你提供更好的体验。<br>
你刚刚注册的账号：<span style=" color: #E91E63; font-weight: 600;">%s </span> <br>
验证码为：<span style=" color: #E91E63; font-weight: 600;">%s</span> <br>
(30分钟内有效,若非本人操作,请忽略此邮件)<br>
 🚌 🚍  🚎  🚏  🚐 🚑  🚒  🚓  🚔 🚕 🚖 🚗 🚘 🚚 🚛 <br>
</div>
` 

//init the config variable
func init(){
    serverhost = config.Emailhost
    port = config.Emailport
    myemail = config.MyAccount
    mypassword = config.MyPassword
}

//send the comfirm email to user after it register
//index is the rank of it new account
func SendConfrimEmail(account data.Account1, index int) int {
    fmt.Println(account)
    toEmail  := account.Email
    username := account.Name
    code := account.Code
    message := createEmail(toEmail, index, username, code)
    auth := createAutn()
    err := SendMailUsingTLS(
        fmt.Sprintf("%s:%d", serverhost, port),    //address of email server
        auth,
        myemail,
        []string{toEmail},
        []byte(message),
    )
    if err != nil {
        HandleError("Send email fall,",err,1)
        return othererror
    }else{
       mylog.Log("Send Email scuess to " + toEmail)
       return scuess
    }
}
 
//create an auth
func createAutn() smtp.Auth {
    return smtp.PlainAuth(
        "",
        myemail,
        mypassword,
        serverhost,
    )
}
//create an emial by push the nessary varibale into the emil templace
func createEmail(toEmail string, num int, username string ,code string)( message string){
    header := make(map[string]string)
    header["From"] = "BlackCarDriver.cn" + "<" + myemail + ">"  
    header["To"] = toEmail
    header["Subject"] = " 你好，🚓 验证码到了！"
    header["Content-Type"] = "text/html; charset=UTF-8"
    for k, v := range header {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + fmt.Sprintf(confrimBody ,num, username, code)
    fmt.Println("--------------------")
    fmt.Println(message)
    fmt.Println("--------------------")
    return message
}

//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
    //problem : certificate signed by unknown authority
    var tr = &tls.Config{InsecureSkipVerify: true}
    conn, err := tls.Dial("tcp", addr, tr)
    if err != nil {
        log.Println("Dialing Error:", err)
        return nil, err
    }
    //分解主机端口字符串
    host, _, _ := net.SplitHostPort(addr)
    return smtp.NewClient(conn, host)
}
 
//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
    //create smtp client
    c, err := Dial(addr)
    if err != nil {
        log.Println("Create smpt client error:", err)
        return err
    }
    defer c.Close()

    if auth != nil {
        if ok, _ := c.Extension("AUTH"); ok {
            if err = c.Auth(auth); err != nil {
                log.Println("Error during AUTH", err)
                return err
            }
        }
    }
 
    if err = c.Mail(from); err != nil {
        return err
    }
 
    for _, addr := range to {
        if err = c.Rcpt(addr); err != nil {
            return err
        }
    }
 
    w, err := c.Data()
    if err != nil {
        return err
    }
 
    _, err = w.Write(msg)
    if err != nil {
        return err
    }
 
    err = w.Close()
    if err != nil {
        return err
    }
 
    return c.Quit()
}

