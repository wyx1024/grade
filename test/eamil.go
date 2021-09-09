package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"
)

func main() {
	e := email.NewEmail()
	e.From = "wuyanxiang1024@163.com"
	e.To = []string{"303076208@qq.com"}
	e.Subject = "测试邮箱"
	e.Text = []byte("测试")
	fmt.Println(time.Now().Second())
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "wuyanxiang1024@163.com", "CKRGDOVMLIDBETDJ", "smtp.163.com"))
	fmt.Println(time.Now().Second())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(err)
	time.Sleep(5*time.Second)
}