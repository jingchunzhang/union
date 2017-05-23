package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/tiantour/fetch"
)

var (
	// AppID appid
	AppID string

	// AppSecret app secret
	AppSecret string

	// Code code
	Code string

	// AccessToken access token
	AccessToken string

	// RefreshToken refresh token
	RefreshToken string

	// OpenID openID
	OpenID string
)

type (
	// Wechat wechat
	Wechat struct {
		OpenID     string   `json:"openid"`     // 用户的唯一标识
		NickName   string   `json:"nickname"`   // 用户昵称
		Sex        int      `json:"sex"`        // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
		Province   string   `json:"province"`   // 用户个人资料填写的省份
		City       string   `json:"city"`       // 普通用户个人资料填写的城市
		Country    string   `json:"country"`    // 国家，如中国为CN
		HeadImgURL string   `json:"headimgurl"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
		Privilege  []string `json:"privilege"`  // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
		UnionID    string   `json:"unionid"`    // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。详见：获取用户个人信息（UnionID机制）
		Language   string   `json:"language"`   // 语言
	}
)

// NewWechat new wechat
func NewWechat() *Wechat {
	return &Wechat{}
}

// User user
func (w Wechat) User() (Wechat, error) {
	result := Wechat{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
		AccessToken,
		OpenID,
	)
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    url,
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}