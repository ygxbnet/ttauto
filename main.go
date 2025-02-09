package main

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/tidwall/gjson"
	"log"
	"os"
	"path"
	"time"
	"ttauto/internal/ttapi"
)

var (
	phone       string
	captchaCode string
	SMSCode     string
	unionID     string
	token       string

	unionIDFile string
)

func main() {
	fmt.Println("欢迎使用 TTAuto\n")

	programPath, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	unionIDFile = path.Join(programPath, "union_id")

	_, err = os.Stat(unionIDFile) // 判断 union_id 文件是否存在
	if err != nil {
		fmt.Println("请登陆甜糖心愿，共3步")
		login()
		fmt.Println("登陆成功，开始执行后续程序...\n")
	}

	// 读取本地 union_id，并验证是否可用
	fmt.Println("已有 union_id，正在验证是否可用...")
	rawFile, err := os.ReadFile(unionIDFile)
	if err != nil {
		log.Panic("读取 union_id 文件失败：", err)
	}
	unionID = string(rawFile)
	response, err := ttapi.RefreshToken(unionID)
	if err != nil {
		log.Panic("刷新 token 失败，请重新登陆\n", err)
	}
	if gjson.Get(response, "errCode").Int() != 0 { // 判断是否刷新 token 成功
		log.Panic("无法刷新 token，union_id 无效，可能需要重新登陆\n", response)
	}
	fmt.Println("union_id 有效，token 刷新成功！")
	token = gjson.Get(response, "data.token").String()

	refreshTokensRegularly() // 定时刷新 token

	// 定时任务，每天 3:00 定时执行一次
	c := cron.NewWithLocation(time.FixedZone("CST", 8*3600))
	fmt.Println("【定时任务】已开启定时任务，每天 3:00 定时签到")
	c.AddFunc("0 3 * * *", func() {
		signInInfo, err := ttapi.SignIn(token)
		if err != nil || signInInfo != "{}" {
			fmt.Println("【定时签到】签到失败：", signInInfo, err)
		} else {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " 【定时签到】签到成功！")
		}
	})
	c.Start()

	// 阻塞
	select {}
}

// 登陆
func login() {
	// 获取手机号
	fmt.Println("1.请输入电话号码：")
	fmt.Scanln(&phone)

	// 获取4位数图形验证码
	captchaID, captchaURL, err := ttapi.SendCaptchaRequest()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(captchaURL, "\n2.请访问上面的网址，然后在此输入4位验证了码：")
	fmt.Scanln(&captchaCode)

	// 获取短信验证码
	retInfo, err := ttapi.SendSMS(phone, captchaID, captchaCode)
	if err != nil {
		log.Panic(err)
	}
	if retInfo != "{}" {
		log.Panic("短信发送失败，请检查手机号是否正确\n", retInfo)
	} else {
		fmt.Println("3.短信发送成功，请稍后输入短信验证码：")
	}
reLogin:
	fmt.Scanln(&SMSCode)

	// 短信验证码登陆
	userInfo, err := ttapi.Login(phone, SMSCode)
	if gjson.Get(userInfo, "errCode").Int() == 1 {
		fmt.Println("验证码错误，请重新输入验证码：")
		goto reLogin
	} else if gjson.Get(userInfo, "errCode").Int() != 0 {
		log.Panic("登陆失败，请检查其他设置是否正确\n", userInfo)
		return
	}
	fmt.Println("登陆成功！")
	fmt.Printf(`===================================
	用户名：%s
	手机号：%s
	等级：%s
	union_id：%s
	token：%s
===================================
`,
		gjson.Get(userInfo, "data.nickName"),
		gjson.Get(userInfo, "data.phoneNum"),
		gjson.Get(userInfo, "data.level"),
		gjson.Get(userInfo, "data.union_id"),
		gjson.Get(userInfo, "data.token"))

	// 保存 union_id
	err = os.WriteFile(unionIDFile, []byte(gjson.Get(userInfo, "data.union_id").String()), 0644)
	if err != nil {
		log.Panic("保存 union_id 失败，请手动保存\n", err)
	}
	fmt.Println("已将 union_id 保存至以下文件中\n", unionIDFile)
}

// 定时刷新 token
func refreshTokensRegularly() {
	ticker := time.NewTicker(40 * time.Hour) // 40小时刷新一次
	go func() {
		for {
			<-ticker.C

			response, err := ttapi.RefreshToken(unionID)
			if err != nil {
				log.Panic("【定时任务】定时刷新 token 失败，请重新登陆\n", err)
			}
			token = gjson.Get(response, "data.token").String()
			fmt.Println("【定时任务】定时刷新 token 成功")
		}
	}()
}
