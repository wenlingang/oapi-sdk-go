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

const (
	UserIdTypeOpenId  = "open_id"  // 用户的 open id
	UserIdTypeUnionId = "union_id" // 用户的 union id
	UserIdTypeUserId  = "user_id"  // 用户的 user id
)

type Credentials struct {
	Email  *string `json:"email,omitempty"`   // 邮箱
	Mobile *string `json:"mobile,omitempty"`  // 手机号
	UserId *string `json:"user_id,omitempty"` // 用户 id
}

type CredentialsBuilder struct {
	email      string // 邮箱
	emailFlag  bool
	mobile     string // 手机号
	mobileFlag bool
	userId     string // 用户 id
	userIdFlag bool
}

func NewCredentialsBuilder() *CredentialsBuilder {
	builder := &CredentialsBuilder{}
	return builder
}

// 邮箱
//
// 示例值：q*****@qq.com
func (builder *CredentialsBuilder) Email(email string) *CredentialsBuilder {
	builder.email = email
	builder.emailFlag = true
	return builder
}

// 手机号
//
// 示例值：186*****01
func (builder *CredentialsBuilder) Mobile(mobile string) *CredentialsBuilder {
	builder.mobile = mobile
	builder.mobileFlag = true
	return builder
}

// 用户 id
//
// 示例值：
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

type IdpCredential struct {
}

type IdpCredentialId struct {
	IdpCredentialId *string `json:"idp_credential_id,omitempty"` // idp 侧租户唯一标识
}

type IdpCredentialIdBuilder struct {
	idpCredentialId     string // idp 侧租户唯一标识
	idpCredentialIdFlag bool
}

func NewIdpCredentialIdBuilder() *IdpCredentialIdBuilder {
	builder := &IdpCredentialIdBuilder{}
	return builder
}

// idp 侧租户唯一标识
//
// 示例值：
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

type MaskSession struct {
	CreateTime   *string `json:"create_time,omitempty"`   // 创建时间
	TerminalType *int    `json:"terminal_type,omitempty"` // 客户端类型
	UserId       *string `json:"user_id,omitempty"`       // 用户 ID
}

type MaskSessionBuilder struct {
	createTime       string // 创建时间
	createTimeFlag   bool
	terminalType     int // 客户端类型
	terminalTypeFlag bool
	userId           string // 用户 ID
	userIdFlag       bool
}

func NewMaskSessionBuilder() *MaskSessionBuilder {
	builder := &MaskSessionBuilder{}
	return builder
}

// 创建时间
//
// 示例值：
func (builder *MaskSessionBuilder) CreateTime(createTime string) *MaskSessionBuilder {
	builder.createTime = createTime
	builder.createTimeFlag = true
	return builder
}

// 客户端类型
//
// 示例值：
func (builder *MaskSessionBuilder) TerminalType(terminalType int) *MaskSessionBuilder {
	builder.terminalType = terminalType
	builder.terminalTypeFlag = true
	return builder
}

// 用户 ID
//
// 示例值：
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

type QuerySessionReqBodyBuilder struct {
	userIds     []string // 用户 ID
	userIdsFlag bool
}

func NewQuerySessionReqBodyBuilder() *QuerySessionReqBodyBuilder {
	builder := &QuerySessionReqBodyBuilder{}
	return builder
}

// 用户 ID
//
//示例值：["47f621ff"]
func (builder *QuerySessionReqBodyBuilder) UserIds(userIds []string) *QuerySessionReqBodyBuilder {
	builder.userIds = userIds
	builder.userIdsFlag = true
	return builder
}

func (builder *QuerySessionReqBodyBuilder) Build() *QuerySessionReqBody {
	req := &QuerySessionReqBody{}
	if builder.userIdsFlag {
		req.UserIds = builder.userIds
	}
	return req
}

type QuerySessionPathReqBodyBuilder struct {
	userIds     []string // 用户 ID
	userIdsFlag bool
}

func NewQuerySessionPathReqBodyBuilder() *QuerySessionPathReqBodyBuilder {
	builder := &QuerySessionPathReqBodyBuilder{}
	return builder
}

// 用户 ID
//
// 示例值：["47f621ff"]
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

type QuerySessionReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *QuerySessionReqBody
}

func NewQuerySessionReqBuilder() *QuerySessionReqBuilder {
	builder := &QuerySessionReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

//
//
// 示例值：open_id
func (builder *QuerySessionReqBuilder) UserIdType(userIdType string) *QuerySessionReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

// 该接口用于查询用户的登录信息
func (builder *QuerySessionReqBuilder) Body(body *QuerySessionReqBody) *QuerySessionReqBuilder {
	builder.body = body
	return builder
}

func (builder *QuerySessionReqBuilder) Build() *QuerySessionReq {
	req := &QuerySessionReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	req.apiReq.Body = builder.body
	return req
}

type QuerySessionReqBody struct {
	UserIds []string `json:"user_ids,omitempty"` // 用户 ID
}

type QuerySessionReq struct {
	apiReq *larkcore.ApiReq
	Body   *QuerySessionReqBody `body:""`
}

type QuerySessionRespData struct {
	MaskSessions []*MaskSession `json:"mask_sessions,omitempty"` // 用户登录信息
}

type QuerySessionResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *QuerySessionRespData `json:"data"` // 业务数据
}

func (resp *QuerySessionResp) Success() bool {
	return resp.Code == 0
}
