// Package optical_char_recognition code generated by oapi sdk gen
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

package larkoptical_char_recognition

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *OpticalCharRecognitionService {
	o := &OpticalCharRecognitionService{config: config}
	o.Image = &image{service: o}
	return o
}

type OpticalCharRecognitionService struct {
	config *larkcore.Config
	Image  *image // 图片识别
}

type image struct {
	service *OpticalCharRecognitionService
}

// 基础图片识别 (OCR)
//
// - 可识别图片中的文字，按图片中的区域划分，分段返回文本列表
//
// - 单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/optical_char_recognition-v1/image/basic_recognize
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/optical_char_recognitionv1/basicRecognize_image.go
func (i *image) BasicRecognize(ctx context.Context, req *BasicRecognizeImageReq, options ...larkcore.RequestOptionFunc) (*BasicRecognizeImageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/optical_char_recognition/v1/image/basic_recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BasicRecognizeImageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
