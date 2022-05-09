package miniapp

import (
	"baby-daily-api/configs"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type miniAppService struct {
}

type MiniAppService interface {
	Code2Session(code string) (*Session, error)
}

func NewMiniAppService() MiniAppService {
	return &miniAppService{}
}

// Session code2session返回的会话信息
type Session struct {
	OpenId     string `json:"openid"`
	UnionId    string `json:"unionid"`
	SessionKey string `json:"session_key"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// Code2Session 换取session
func (m *miniAppService) Code2Session(code string) (*Session, error) {
	conf := configs.Get().Wechat
	session := &Session{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", conf.AppId,
		conf.Secret,
		code)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, session)
	if err != nil {
		return nil, err
	}

	if session.OpenId == "" || session.SessionKey == "" {
		return nil, errors.New(session.ErrMsg)
	}

	return session, nil
}
