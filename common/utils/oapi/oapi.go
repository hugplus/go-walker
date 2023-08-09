package oapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hugplus/go-walker/common/utils/http_util"
)

var m = map[string]*OApi{}

type OApi struct {
	client *http_util.HTTPClient
}

func New(baseUrl string) *OApi {
	if oa, ok := m[baseUrl]; ok {
		return oa
	}
	oa := &OApi{
		client: &(http_util.HTTPClient{
			BaseURL: baseUrl,
			Headers: map[string]string{
				"Content-Type": "application/json",
				//"Authorization": "Bearer xxxxxxxxxxxx",
			},
		}),
	}
	m[baseUrl] = oa
	return oa
}

func (e *OApi) Header(m map[string]string) *OApi {
	for k, v := range m {
		e.client.Headers[k] = v
	}
	return e
}

var userUri = "/v2/sso/getUserinfo"

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Userinfo struct {
	UserId     string    `json:"userId" comment:"用户id"`
	Username   string    `json:"username" comment:"用户名"`
	Mobile     string    `json:"mobile" comment:"手机号"`
	Email      string    `json:"email" comment:"邮箱"`
	FirstName  string    `json:"firstName" comment:"名"`
	LastName   string    `json:"lastName" comment:"姓"`
	Nickname   string    `json:"nickname" comment:"昵称"`
	Avatar     string    `json:"avatar" comment:"头像"`
	Bio        string    `json:"bio" comment:"签名"`
	Gender     string    `json:"gender" comment:"性别 0 女 1 男 2 未知"`
	Birthday   time.Time `json:"birthday" comment:"生日"`
	Inviter    string    `json:"inviter" gorm:"type:varchar(32);default:(-);comment:邀请人"` //邀请人
	InviteType int       `json:"inviteType" gorm:"type:tinyint;comment:邀请类型"`             //邀请类型
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`                           //创建时间
}

func (u Userinfo) GetName() string {
	if u.LastName != "" && u.FirstName != "" {
		return u.LastName + u.FirstName
	} else if u.FirstName != "" {
		return u.FirstName
	} else if u.LastName != "" {
		return u.LastName
	} else {
		return u.Nickname
	}
}

func (e *OApi) GetUserInfo(userId string, user *Userinfo) error {
	m := map[string]string{
		"id": userId,
	}
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	response, err := e.client.Post(userUri, data)
	if err != nil {
		//fmt.Println("POST error:", err)
		return err
	}
	fmt.Println(string(response))
	var res Res
	if err := json.Unmarshal(response, &res); err != nil {
		return err
	}
	if res.Code == 200 {
		d, err := json.Marshal(res.Data)
		if err != nil {
			return err
		}
		return json.Unmarshal(d, user)
	} else {
		return errors.New(res.Msg)
	}
}

var vcUri = "/v2/sso/verify/code"

func (e *OApi) VerifyCode(username, code string, ret *bool) error {
	m := map[string]string{
		"username": username,
		"code":     code,
	}
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	response, err := e.client.Post(vcUri, data)
	if err != nil {
		return err
	}
	var res Res
	if err := json.Unmarshal(response, &res); err != nil {
		return err
	}
	if res.Code == 200 {
		*ret = true
	} else {
		*ret = false
	}
	return nil
}

var cciUri = "/v2/retail/createUserCompanyInvite"

func (e *OApi) CreateCompanyInvite(companyName, companyId, invitationCode string, ret *bool) error {
	m := map[string]string{
		"companyName":    companyName,
		"companyId":      companyId,
		"invitationCode": invitationCode,
	}
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	response, err := e.client.Post(cciUri, data)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	var res Res
	if err := json.Unmarshal(response, &res); err != nil {
		return err
	}
	if res.Code == 200 {
		*ret = true
	} else {
		*ret = false
	}
	return nil
}
