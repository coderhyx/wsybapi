package controllers

import (
	"context"
	"wsybapi/models"
	"wsybapi/proto"

	"regexp"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserRpcController struct {
	beego.Controller
}

// 用户登录
// @router /login/do [*]
func (this *UserRpcController) LoginDo(ctx context.Context, req *proto.RequestLoginDo, res *proto.ResponseLoginDo) error {
	var (
		userLoginProto proto.LoginUser
		isorno         bool
		uid            int
		name           string
	)

	mobile := req.Mobile
	password := req.Password

	if mobile == "" {
		res.Code = 4001
		res.Msg = "手机号不能为空"
		goto ERR
	}
	isorno, _ = regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		res.Code = 4002
		res.Msg = "手机号格式不正确"
		goto ERR
	}
	if password == "" {
		res.Code = 4003
		res.Msg = "密码不能为空"
		goto ERR
	}
	uid, name = models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		userLoginProto.Uid = int64(uid)
		userLoginProto.Username = name
		res.Code = 0
		res.Msg = "登录成功"
		res.Items = &userLoginProto
		res.Count = 1
		return nil
	} else {
		res.Code = 4004
		res.Msg = "手机号或密码不正确"
		goto ERR
	}
ERR:
	res.Items = &userLoginProto
	res.Count = 0
	return nil
}
