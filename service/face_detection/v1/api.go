// Package face_detection code generated by oapi sdk gen
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

package larkface_detection

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *FaceDetectionService {
	f := &FaceDetectionService{config: config}
	f.Image = &image{service: f}
	return f
}

type FaceDetectionService struct {
	config *larkcore.Config
	Image  *image // 图片
}

type image struct {
	service *FaceDetectionService
}

// 人脸检测和属性分析
//
// - 检测图片中的人脸属性和质量等信息
//
// - 注意：返回值为 -1 表示该功能还暂未实现
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/face_detection-v1/image/detect_face_attributes
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/face_detectionv1/detectFaceAttributes_image.go
func (i *image) DetectFaceAttributes(ctx context.Context, req *DetectFaceAttributesImageReq, options ...larkcore.RequestOptionFunc) (*DetectFaceAttributesImageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/face_detection/v1/image/detect_face_attributes"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DetectFaceAttributesImageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
