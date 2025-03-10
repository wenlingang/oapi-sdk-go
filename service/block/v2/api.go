// Package block code generated by oapi sdk gen
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

package larkblock

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *BlockService {
	b := &BlockService{config: config}
	b.Entity = &entity{service: b}
	b.Message = &message{service: b}
	return b
}

type BlockService struct {
	config  *larkcore.Config
	Entity  *entity  // 服务端 API
	Message *message // 服务端 API
}

type entity struct {
	service *BlockService
}
type message struct {
	service *BlockService
}

// 创建 BlockEntity
//
// - 开发者可以通过该接口将部分或全部数据存放于 BlockEntity。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/block-v2/entity/create
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/blockv2/create_entity.go
func (e *entity) Create(ctx context.Context, req *CreateEntityReq, options ...larkcore.RequestOptionFunc) (*CreateEntityResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/block/v2/entities"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateEntityResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, e.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 更新 BlockEntity
//
// - 开发者通过该接口可以更新存储在 BlockEntity 中的数据，并实时推送到端侧。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/block-v2/entity/update
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/blockv2/update_entity.go
func (e *entity) Update(ctx context.Context, req *UpdateEntityReq, options ...larkcore.RequestOptionFunc) (*UpdateEntityResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/block/v2/entities/:block_id"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateEntityResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, e.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Block 协同数据推送
//
// - 根据 BlockID 向指定用户列表推送协同数据。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/block-v2/message/create
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/blockv2/create_message.go
func (m *message) Create(ctx context.Context, req *CreateMessageReq, options ...larkcore.RequestOptionFunc) (*CreateMessageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/block/v2/message"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateMessageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, m.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
