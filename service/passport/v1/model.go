// Package passport code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkpassport

import (
	"fmt"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 生成枚举值

const (
	UserIdTypeOpenId  = "open_id"
	UserIdTypeUnionId = "union_id"
	UserIdTypeUserId  = "user_id"
)

// 生成数据类型

type Credentials struct {
	Email  *string `json:"email,omitempty"`
	Mobile *string `json:"mobile,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// builder开始
type CredentialsBuilder struct {
	email      string
	emailFlag  bool
	mobile     string
	mobileFlag bool
	userId     string
	userIdFlag bool
}

func NewCredentialsBuilder() *CredentialsBuilder {
	builder := &CredentialsBuilder{}
	return builder
}

func (builder *CredentialsBuilder) Email(email string) *CredentialsBuilder {
	builder.email = email
	builder.emailFlag = true
	return builder
}
func (builder *CredentialsBuilder) Mobile(mobile string) *CredentialsBuilder {
	builder.mobile = mobile
	builder.mobileFlag = true
	return builder
}
func (builder *CredentialsBuilder) UserId(userId string) *CredentialsBuilder {
	builder.userId = userId
	builder.userIdFlag = true
	return builder
}

func (builder *CredentialsBuilder) Build() *Credentials {
	req := &Credentials{}
	if builder.emailFlag {
		req.Email = &builder.email

	}
	if builder.mobileFlag {
		req.Mobile = &builder.mobile

	}
	if builder.userIdFlag {
		req.UserId = &builder.userId

	}
	return req
}

// builder结束

type IdpCredential struct {
}

// builder开始
// builder结束

type IdpCredentialId struct {
	IdpCredentialId *string `json:"idp_credential_id,omitempty"`
}

// builder开始
type IdpCredentialIdBuilder struct {
	idpCredentialId     string
	idpCredentialIdFlag bool
}

func NewIdpCredentialIdBuilder() *IdpCredentialIdBuilder {
	builder := &IdpCredentialIdBuilder{}
	return builder
}

func (builder *IdpCredentialIdBuilder) IdpCredentialId(idpCredentialId string) *IdpCredentialIdBuilder {
	builder.idpCredentialId = idpCredentialId
	builder.idpCredentialIdFlag = true
	return builder
}

func (builder *IdpCredentialIdBuilder) Build() *IdpCredentialId {
	req := &IdpCredentialId{}
	if builder.idpCredentialIdFlag {
		req.IdpCredentialId = &builder.idpCredentialId

	}
	return req
}

// builder结束

type MaskSession struct {
	CreateTime   *string `json:"create_time,omitempty"`
	TerminalType *int    `json:"terminal_type,omitempty"`
	UserId       *string `json:"user_id,omitempty"`
}

// builder开始
type MaskSessionBuilder struct {
	createTime       string
	createTimeFlag   bool
	terminalType     int
	terminalTypeFlag bool
	userId           string
	userIdFlag       bool
}

func NewMaskSessionBuilder() *MaskSessionBuilder {
	builder := &MaskSessionBuilder{}
	return builder
}

func (builder *MaskSessionBuilder) CreateTime(createTime string) *MaskSessionBuilder {
	builder.createTime = createTime
	builder.createTimeFlag = true
	return builder
}
func (builder *MaskSessionBuilder) TerminalType(terminalType int) *MaskSessionBuilder {
	builder.terminalType = terminalType
	builder.terminalTypeFlag = true
	return builder
}
func (builder *MaskSessionBuilder) UserId(userId string) *MaskSessionBuilder {
	builder.userId = userId
	builder.userIdFlag = true
	return builder
}

func (builder *MaskSessionBuilder) Build() *MaskSession {
	req := &MaskSession{}
	if builder.createTimeFlag {
		req.CreateTime = &builder.createTime

	}
	if builder.terminalTypeFlag {
		req.TerminalType = &builder.terminalType

	}
	if builder.userIdFlag {
		req.UserId = &builder.userId

	}
	return req
}

// builder结束

// 生成请求和响应结果类型，以及请求对象的Builder构造器

type QuerySessionReqBodyBuilder struct {
	userIds     []string
	userIdsFlag bool
}

// 生成body的New构造器
func NewQuerySessionReqBodyBuilder() *QuerySessionReqBodyBuilder {
	builder := &QuerySessionReqBodyBuilder{}
	return builder
}

// 1.2 生成body的builder属性方法
func (builder *QuerySessionReqBodyBuilder) UserIds(userIds []string) *QuerySessionReqBodyBuilder {
	builder.userIds = userIds
	builder.userIdsFlag = true
	return builder
}

// 1.3 生成body的build方法
func (builder *QuerySessionReqBodyBuilder) Build() *QuerySessionReqBody {
	req := &QuerySessionReqBody{}
	if builder.userIdsFlag {
		req.UserIds = builder.userIds
	}
	return req
}

// 上传文件path开始
type QuerySessionPathReqBodyBuilder struct {
	userIds     []string
	userIdsFlag bool
}

func NewQuerySessionPathReqBodyBuilder() *QuerySessionPathReqBodyBuilder {
	builder := &QuerySessionPathReqBodyBuilder{}
	return builder
}
func (builder *QuerySessionPathReqBodyBuilder) UserIds(userIds []string) *QuerySessionPathReqBodyBuilder {
	builder.userIds = userIds
	builder.userIdsFlag = true
	return builder
}

func (builder *QuerySessionPathReqBodyBuilder) Build() (*QuerySessionReqBody, error) {
	req := &QuerySessionReqBody{}
	if builder.userIdsFlag {
		req.UserIds = builder.userIds
	}
	return req, nil
}

// 上传文件path结束

// 1.4 生成请求的builder结构体
type QuerySessionReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *QuerySessionReqBody
}

// 生成请求的New构造器
func NewQuerySessionReqBuilder() *QuerySessionReqBuilder {
	builder := &QuerySessionReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *QuerySessionReqBuilder) UserIdType(userIdType string) *QuerySessionReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}
func (builder *QuerySessionReqBuilder) Body(body *QuerySessionReqBody) *QuerySessionReqBuilder {
	builder.body = body
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *QuerySessionReqBuilder) Build() *QuerySessionReq {
	req := &QuerySessionReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	req.apiReq.Body = builder.body
	return req
}

type QuerySessionReqBody struct {
	UserIds []string `json:"user_ids,omitempty"`
}

type QuerySessionReq struct {
	apiReq *larkcore.ApiReq
	Body   *QuerySessionReqBody `body:""`
}

type QuerySessionRespData struct {
	MaskSessions []*MaskSession `json:"mask_sessions,omitempty"`
}

type QuerySessionResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *QuerySessionRespData `json:"data"`
}

func (resp *QuerySessionResp) Success() bool {
	return resp.Code == 0
}

// 生成消息事件结构体

// 生成请求的builder构造器
// 1.1 生成body的builder结构体
