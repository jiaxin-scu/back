// description: 发送邮件的工具包
//
// author: vignetting
// time: 2021/5/10

package mail

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"strconv"
	"structure/pkg/setting"
)

var pool *email.Pool

// description: 用于初始化邮箱连接池
func Setup() {
	var err error
	if pool, err = email.NewPool(setting.MailSetting.Host+":"+strconv.Itoa(setting.MailSetting.Port), setting.MailSetting.PoolSize, smtp.PlainAuth("", setting.MailSetting.UserName, setting.MailSetting.Password, setting.MailSetting.Host)); err != nil {
		panic("构建由此池失败，错误原因为" + err.Error())
	}
}

// description: 	发送普通邮件，无超时限制，但会阻塞，直至发送完成
// param: toMail	目标邮箱
// param: subject	邮件主题
// param: text		邮件内容
// return: error	邮件发送失败原因
func SendTextMail(toMail, subject, text string) error {
	// 1. 构建邮件
	newMail := email.NewEmail()
	newMail.From = setting.MailSetting.UserName
	newMail.To = []string{toMail}
	newMail.Subject = subject
	newMail.Text = []byte(text)

	// 2. 发送邮件
	return pool.Send(newMail, setting.MailSetting.Timeout)
}

// description: 	发送 HTML 格式的邮件，无超时限制，但会阻塞，直至发送完成
// param: toMail  	目标邮箱
// param: subject	邮件主题
// param: html		邮件内容
// return: error	邮件发送失败原因
func SendHTMLMail(toMail, subject, html string) error {
	// 1. 构建邮件
	newMail := email.NewEmail()
	newMail.From = setting.MailSetting.UserName
	newMail.To = []string{toMail}
	newMail.Subject = subject
	newMail.HTML = []byte(html)

	// 2. 发送邮件
	return pool.Send(newMail, setting.MailSetting.Timeout)
}
