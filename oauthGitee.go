package oauth

import (
	"encoding/json"
	"errors"
	"regexp"
)

//GetRurl 获取登录地址
func (e *AppConf) GetRurl(state string) string {
	return "https://gitee.com/oauth/authorize?client_id=" + e.Conf.AppId + "&redirect_uri=" + e.Conf.Rurl + "&response_type=code"
}

// GetToken 获取token
func (e *AppConf) GetToken(code string) (*AuthGiteeSuccRes, error) {

	str, err := HttpPost("https://gitee.com/oauth/token?grant_type=authorization_code&code=" + code + "&client_id=" + e.Conf.AppId + "&redirect_uri=" + e.Conf.Rurl + "&client_secret=" + e.Conf.AppKey)
	if err != nil {
		return nil, err
	}
	// 用正则表达式检查字符串中是否包含"error"
	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {
		return nil, errors.New(str)
	} else {
		p := &AuthGiteeSuccRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}
}

// GetUserInfo 获取第三方用户信息
func (e *AppConf) GetUserInfo(access_token string) (string, error) {

	str, err := HttpGet("https://gitee.com/api/v5/user?access_token=" + access_token)
	if err != nil {
		return "", err
	}
	return string(str), nil
	// 根据具体需要可以做一个，反序列化，导向用户信息结构体
}
