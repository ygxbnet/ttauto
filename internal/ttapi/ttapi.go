package ttapi

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

var client = &http.Client{}

// SendCaptchaRequest 获取图形验证码
func SendCaptchaRequest() (captchaID string, captchaURL string, reterr error) {
	// 发送请求
	req, _ := http.NewRequest(http.MethodGet, baseURL+getCaptchaRequestURL, nil)
	response, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer response.Body.Close()

	// 返回数据后期处理
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}
	return gjson.Get(string(data), "data.captchaId").String(),
		baseURL + gjson.Get(string(data), "data.captchaUrl").String(),
		nil
}

// SendSMS 发送短信验证码
func SendSMS(phone string, captchaID string, captchaCode string) (retResponse string, reterr error) {
	// 请求数据拼装
	// 短信验证码登陆结构体
	type sendSMSJsonType struct {
		Phone       string `json:"phone"`
		CaptchaID   string `json:"captchaId"`
		CaptchaCode string `json:"captchaCode"`
	}
	requestData := sendSMSJsonType{Phone: phone, CaptchaID: captchaID, CaptchaCode: captchaCode}
	jsonData, err := json.MarshalIndent(requestData, "", "  ")
	if err != nil {
		return "", err
	}

	// 发送请求
	req, _ := http.NewRequest(http.MethodPost, baseURL+sendSMSURL, strings.NewReader(string(jsonData)))
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 返回数据后期处理
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Login 短信验证码登陆
func Login(phone string, SMSCode string) (retResponse string, reterr error) {
	// 发送请求
	// url.URL{}
	req, _ := http.NewRequest(http.MethodPost,
		fmt.Sprintf(baseURL+loginURL+"?phone=%s&authCode=%s", phone, SMSCode),
		nil)
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 返回数据后期处理
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// RefreshToken 刷新 Token
func RefreshToken(unionID string) (retResponse string, reterr error) {
	// 请求数据拼装
	body := fmt.Sprintf("{\"union_id\":\"%s\"}", unionID)

	// 发送请求
	req, _ := http.NewRequest(http.MethodPost, baseURL+refreshTokenURL, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 返回数据后期处理
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(token string) (retResponse string, reterr error) {
	req, _ := http.NewRequest(http.MethodPost, baseURL+getUserInfoURL, nil)
	req.Header.Add("Authorization", "Bearer "+token) // 身份验证
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 返回数据后期处理
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// SignIn 签到
func SignIn(token string) (retResponse string, reterr error) {
	req, _ := http.NewRequest(http.MethodPost, baseURL+signInURL, nil)
	req.Header.Add("Authorization", "Bearer "+token) // 身份验证
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 返回数据后期处理
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
