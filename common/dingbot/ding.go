package dingbot

import "github.com/blinkbean/dingtalk"

func SendReg(secret, username, userid, keyword, inviter string) error {
	content := "[广播]新注册用户！！！用户手机号/邮箱： " + username
	if keyword != "" {
		content += "，来源关键词是：" + keyword
	}
	if inviter != "" {
		content += "，来源关键词是：" + inviter + "推荐的"
	}
	return send(secret, content)
}

func send(secret, content string) error {
	// 单个机器人有单位时间内消息条数的限制，如果有需要可以初始化多个token，发消息时随机发给其中一个机器人。
	var dingToken = []string{secret}
	cli := dingtalk.InitDingTalk(dingToken, ".")
	return cli.SendTextMessage(content)
}
