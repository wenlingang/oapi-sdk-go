// Package ehr code generated by oapi sdk gen
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

package larkehr

import (
	"bytes"
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *EhrService {
	e := &EhrService{config: config}
	e.Attachment = &attachment{service: e}
	e.Employee = &employee{service: e}
	return e
}

type EhrService struct {
	config     *larkcore.Config
	Attachment *attachment // 飞书人事（标准版)
	Employee   *employee   // 飞书人事（标准版)
}

type attachment struct {
	service *EhrService
}
type employee struct {
	service *EhrService
}

// 下载附件
//
// - 根据文件 token 下载文件。;;调用「批量获取员工花名册信息」接口的返回值中，「文件」类型的字段 id，即是文件 token
//
// - ![image.png](//sf1-ttcdn-tos.pstatp.com/obj/open-platform-opendoc/bed391d2a8ce6ed2d5985ea69bf92850_9GY1mnuDXP.png)
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/attachment/get
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/ehrv1/get_attachment.go
func (a *attachment) Get(ctx context.Context, req *GetAttachmentReq, options ...larkcore.RequestOptionFunc) (*GetAttachmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/ehr/v1/attachments/:token"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetAttachmentResp{ApiResp: apiResp}
	// 如果是下载，则设置响应结果
	if apiResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(apiResp.RawBody)
		resp.FileName = larkcore.FileNameByHeader(apiResp.Header)
		return resp, err
	}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 批量获取员工花名册信息
//
// - 根据员工飞书用户 ID / 员工状态 / 雇员类型等搜索条件 ，批量获取员工花名册字段信息。字段包括「系统标准字段 / system_fields」和「自定义字段 / custom_fields」
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/employee/list
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/ehrv1/list_employee.go
func (e *employee) List(ctx context.Context, req *ListEmployeeReq, options ...larkcore.RequestOptionFunc) (*ListEmployeeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/ehr/v1/employees"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListEmployeeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, e.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *employee) ListByIterator(ctx context.Context, req *ListEmployeeReq, options ...larkcore.RequestOptionFunc) (*ListEmployeeIterator, error) {
	return &ListEmployeeIterator{
		ctx:      ctx,
		req:      req,
		listFunc: e.List,
		options:  options,
		limit:    req.Limit}, nil
}
