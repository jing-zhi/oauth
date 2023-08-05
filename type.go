package oauth

type AppConf struct {
	AppId  string
	AppKey string // 等价于 secret
	Rurl   string
}

// AuthGitee gitee
type AuthGitee struct {
	Conf *AppConf
}
type AuthGiteeErrRes struct {
	Error            int    `json:"errorCode"`
	ErrorDescription string `json:"errMsg"`
}
type AuthGiteeSuccRes struct {
	AccessToken string `json:"accessToken"`
}
