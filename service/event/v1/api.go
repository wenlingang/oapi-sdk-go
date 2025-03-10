// Package event code generated by oapi sdk gen
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

package larkevent

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *EventService {
	e := &EventService{config: config}
	e.OutboundIp = &outboundIp{service: e}
	return e
}

type EventService struct {
	config     *larkcore.Config
	OutboundIp *outboundIp // 事件订阅
}

type outboundIp struct {
	service *EventService
}

// 获取事件出口 IP
//
// - 飞书开放平台向应用配置的回调地址推送事件时，是通过特定的 IP 发送出去的，应用可以通过本接口获取所有相关的 IP 地址。
//
// - IP 地址有变更可能，建议应用每隔 6 小时定时拉取最新的 IP 地址，以免由于企业防火墙设置，导致应用无法及时接收到飞书开放平台推送的事件。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-v1/outbound_ip/list
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/eventv1/list_outboundIp.go
func (o *outboundIp) List(ctx context.Context, req *ListOutboundIpReq, options ...larkcore.RequestOptionFunc) (*ListOutboundIpResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/event/v1/outbound_ip"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, o.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListOutboundIpResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, o.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (o *outboundIp) ListByIterator(ctx context.Context, req *ListOutboundIpReq, options ...larkcore.RequestOptionFunc) (*ListOutboundIpIterator, error) {
	return &ListOutboundIpIterator{
		ctx:      ctx,
		req:      req,
		listFunc: o.List,
		options:  options,
		limit:    req.Limit}, nil
}
