// Package acs code generated by oapi sdk gen
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

package larkacs

import (
	"io"

	"io/ioutil"

	"fmt"

	"context"
	"errors"

	"github.com/larksuite/oapi-sdk-go/v3/event"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

const (
	UserIdTypeUserId  = "user_id"  // 以 user_id 来识别用户
	UserIdTypeUnionId = "union_id" // 以 union_id 来识别用户
	UserIdTypeOpenId  = "open_id"  // 以 open_id 来识别用户
)

const (
	UserIdTypeGetUserUserId  = "user_id"  // 以 user_id 来识别用户
	UserIdTypeGetUserUnionId = "union_id" // 以 union_id 来识别用户
	UserIdTypeGetUserOpenId  = "open_id"  // 以 open_id 来识别用户
)

const (
	UserIdTypeListUserUserId  = "user_id"  // 以 user_id 来识别用户
	UserIdTypeListUserUnionId = "union_id" // 以 union_id 来识别用户
	UserIdTypeListUserOpenId  = "open_id"  // 以 open_id 来识别用户
)

const (
	UserIdTypePatchUserUserId  = "user_id"  // 以 user_id 来识别用户
	UserIdTypePatchUserUnionId = "union_id" // 以 union_id 来识别用户
	UserIdTypePatchUserOpenId  = "open_id"  // 以 open_id 来识别用户
)

const (
	UserIdTypeGetUserFaceUserId  = "user_id"  // 以 user_id 来识别用户
	UserIdTypeGetUserFaceUnionId = "union_id" // 以 union_id 来识别用户
	UserIdTypeGetUserFaceOpenId  = "open_id"  // 以 open_id 来识别用户
)

const (
	UserIdTypeUpdateUserFaceUserId  = "user_id"  // 以 user_id 来识别用户
	UserIdTypeUpdateUserFaceUnionId = "union_id" // 以 union_id 来识别用户
	UserIdTypeUpdateUserFaceOpenId  = "open_id"  // 以 open_id 来识别用户
)

type AccessRecord struct {
	AccessRecordId *string `json:"access_record_id,omitempty"` // 门禁记录 ID
	UserId         *string `json:"user_id,omitempty"`          // 门禁记录所属用户 ID
	DeviceId       *string `json:"device_id,omitempty"`        // 门禁设备 ID
	IsClockIn      *bool   `json:"is_clock_in,omitempty"`      // 是否是打卡
	AccessTime     *string `json:"access_time,omitempty"`      // 访问时间，单位秒
	AccessType     *string `json:"access_type,omitempty"`      // 识别方式
	AccessData     *string `json:"access_data,omitempty"`      // 识别相关数据，根据 access_type 不同，取值不同
	IsDoorOpen     *bool   `json:"is_door_open,omitempty"`     // 是否开门
}

type AccessRecordBuilder struct {
	accessRecordId     string // 门禁记录 ID
	accessRecordIdFlag bool
	userId             string // 门禁记录所属用户 ID
	userIdFlag         bool
	deviceId           string // 门禁设备 ID
	deviceIdFlag       bool
	isClockIn          bool // 是否是打卡
	isClockInFlag      bool
	accessTime         string // 访问时间，单位秒
	accessTimeFlag     bool
	accessType         string // 识别方式
	accessTypeFlag     bool
	accessData         string // 识别相关数据，根据 access_type 不同，取值不同
	accessDataFlag     bool
	isDoorOpen         bool // 是否开门
	isDoorOpenFlag     bool
}

func NewAccessRecordBuilder() *AccessRecordBuilder {
	builder := &AccessRecordBuilder{}
	return builder
}

// 门禁记录 ID
//
// 示例值：6939433228970082591
func (builder *AccessRecordBuilder) AccessRecordId(accessRecordId string) *AccessRecordBuilder {
	builder.accessRecordId = accessRecordId
	builder.accessRecordIdFlag = true
	return builder
}

// 门禁记录所属用户 ID
//
// 示例值：ou_7dab8a3d3cdcc9da365777c7ad535d62
func (builder *AccessRecordBuilder) UserId(userId string) *AccessRecordBuilder {
	builder.userId = userId
	builder.userIdFlag = true
	return builder
}

// 门禁设备 ID
//
// 示例值：6939433228970082593
func (builder *AccessRecordBuilder) DeviceId(deviceId string) *AccessRecordBuilder {
	builder.deviceId = deviceId
	builder.deviceIdFlag = true
	return builder
}

// 是否是打卡
//
// 示例值：true
func (builder *AccessRecordBuilder) IsClockIn(isClockIn bool) *AccessRecordBuilder {
	builder.isClockIn = isClockIn
	builder.isClockInFlag = true
	return builder
}

// 访问时间，单位秒
//
// 示例值：1624520221
func (builder *AccessRecordBuilder) AccessTime(accessTime string) *AccessRecordBuilder {
	builder.accessTime = accessTime
	builder.accessTimeFlag = true
	return builder
}

// 识别方式
//
// 示例值：FA
func (builder *AccessRecordBuilder) AccessType(accessType string) *AccessRecordBuilder {
	builder.accessType = accessType
	builder.accessTypeFlag = true
	return builder
}

// 识别相关数据，根据 access_type 不同，取值不同
//
// 示例值：{\"has_access_photo\":true}
func (builder *AccessRecordBuilder) AccessData(accessData string) *AccessRecordBuilder {
	builder.accessData = accessData
	builder.accessDataFlag = true
	return builder
}

// 是否开门
//
// 示例值：true
func (builder *AccessRecordBuilder) IsDoorOpen(isDoorOpen bool) *AccessRecordBuilder {
	builder.isDoorOpen = isDoorOpen
	builder.isDoorOpenFlag = true
	return builder
}

func (builder *AccessRecordBuilder) Build() *AccessRecord {
	req := &AccessRecord{}
	if builder.accessRecordIdFlag {
		req.AccessRecordId = &builder.accessRecordId

	}
	if builder.userIdFlag {
		req.UserId = &builder.userId

	}
	if builder.deviceIdFlag {
		req.DeviceId = &builder.deviceId

	}
	if builder.isClockInFlag {
		req.IsClockIn = &builder.isClockIn

	}
	if builder.accessTimeFlag {
		req.AccessTime = &builder.accessTime

	}
	if builder.accessTypeFlag {
		req.AccessType = &builder.accessType

	}
	if builder.accessDataFlag {
		req.AccessData = &builder.accessData

	}
	if builder.isDoorOpenFlag {
		req.IsDoorOpen = &builder.isDoorOpen

	}
	return req
}

type Device struct {
	DeviceId   *string `json:"device_id,omitempty"`   // 门禁设备 ID
	DeviceName *string `json:"device_name,omitempty"` // 设备名称
	DeviceSn   *string `json:"device_sn,omitempty"`   // 设备 SN 码
}

type DeviceBuilder struct {
	deviceId       string // 门禁设备 ID
	deviceIdFlag   bool
	deviceName     string // 设备名称
	deviceNameFlag bool
	deviceSn       string // 设备 SN 码
	deviceSnFlag   bool
}

func NewDeviceBuilder() *DeviceBuilder {
	builder := &DeviceBuilder{}
	return builder
}

// 门禁设备 ID
//
// 示例值：6939433228970082593
func (builder *DeviceBuilder) DeviceId(deviceId string) *DeviceBuilder {
	builder.deviceId = deviceId
	builder.deviceIdFlag = true
	return builder
}

// 设备名称
//
// 示例值：东门
func (builder *DeviceBuilder) DeviceName(deviceName string) *DeviceBuilder {
	builder.deviceName = deviceName
	builder.deviceNameFlag = true
	return builder
}

// 设备 SN 码
//
// 示例值：3X811621174000240
func (builder *DeviceBuilder) DeviceSn(deviceSn string) *DeviceBuilder {
	builder.deviceSn = deviceSn
	builder.deviceSnFlag = true
	return builder
}

func (builder *DeviceBuilder) Build() *Device {
	req := &Device{}
	if builder.deviceIdFlag {
		req.DeviceId = &builder.deviceId

	}
	if builder.deviceNameFlag {
		req.DeviceName = &builder.deviceName

	}
	if builder.deviceSnFlag {
		req.DeviceSn = &builder.deviceSn

	}
	return req
}

type Feature struct {
	Card         *int  `json:"card,omitempty"`          // 卡号
	FaceUploaded *bool `json:"face_uploaded,omitempty"` // 是否已上传人脸图片
}

type FeatureBuilder struct {
	card             int // 卡号
	cardFlag         bool
	faceUploaded     bool // 是否已上传人脸图片
	faceUploadedFlag bool
}

func NewFeatureBuilder() *FeatureBuilder {
	builder := &FeatureBuilder{}
	return builder
}

// 卡号
//
// 示例值：123456
func (builder *FeatureBuilder) Card(card int) *FeatureBuilder {
	builder.card = card
	builder.cardFlag = true
	return builder
}

// 是否已上传人脸图片
//
// 示例值：true
func (builder *FeatureBuilder) FaceUploaded(faceUploaded bool) *FeatureBuilder {
	builder.faceUploaded = faceUploaded
	builder.faceUploadedFlag = true
	return builder
}

func (builder *FeatureBuilder) Build() *Feature {
	req := &Feature{}
	if builder.cardFlag {
		req.Card = &builder.card

	}
	if builder.faceUploadedFlag {
		req.FaceUploaded = &builder.faceUploaded

	}
	return req
}

type File struct {
	Files    io.Reader `json:"files,omitempty"`     // 人脸图片内容
	FileType *string   `json:"file_type,omitempty"` // 文件类型，可选的类型有 jpg,png
	FileName *string   `json:"file_name,omitempty"` // 带后缀的文件名
}

type FileBuilder struct {
	files        io.Reader // 人脸图片内容
	filesFlag    bool
	fileType     string // 文件类型，可选的类型有 jpg,png
	fileTypeFlag bool
	fileName     string // 带后缀的文件名
	fileNameFlag bool
}

func NewFileBuilder() *FileBuilder {
	builder := &FileBuilder{}
	return builder
}

// 人脸图片内容
//
// 示例值：jpg 图片
func (builder *FileBuilder) Files(files io.Reader) *FileBuilder {
	builder.files = files
	builder.filesFlag = true
	return builder
}

// 文件类型，可选的类型有 jpg,png
//
// 示例值：jpg
func (builder *FileBuilder) FileType(fileType string) *FileBuilder {
	builder.fileType = fileType
	builder.fileTypeFlag = true
	return builder
}

// 带后缀的文件名
//
// 示例值：efeqz12f.jpg
func (builder *FileBuilder) FileName(fileName string) *FileBuilder {
	builder.fileName = fileName
	builder.fileNameFlag = true
	return builder
}

func (builder *FileBuilder) Build() *File {
	req := &File{}
	if builder.filesFlag {
		req.Files = builder.files
	}
	if builder.fileTypeFlag {
		req.FileType = &builder.fileType

	}
	if builder.fileNameFlag {
		req.FileName = &builder.fileName

	}
	return req
}

type User struct {
	Feature *Feature `json:"feature,omitempty"` // 用户特征
	UserId  *string  `json:"user_id,omitempty"` // 用户 ID
}

type UserBuilder struct {
	feature     *Feature // 用户特征
	featureFlag bool
	userId      string // 用户 ID
	userIdFlag  bool
}

func NewUserBuilder() *UserBuilder {
	builder := &UserBuilder{}
	return builder
}

// 用户特征
//
// 示例值：
func (builder *UserBuilder) Feature(feature *Feature) *UserBuilder {
	builder.feature = feature
	builder.featureFlag = true
	return builder
}

// 用户 ID
//
// 示例值：ou_7dab8a3d3cdcc9da365777c7ad535d62
func (builder *UserBuilder) UserId(userId string) *UserBuilder {
	builder.userId = userId
	builder.userIdFlag = true
	return builder
}

func (builder *UserBuilder) Build() *User {
	req := &User{}
	if builder.featureFlag {
		req.Feature = builder.feature
	}
	if builder.userIdFlag {
		req.UserId = &builder.userId

	}
	return req
}

type UserId struct {
	UserId  *string `json:"user_id,omitempty"`  //
	OpenId  *string `json:"open_id,omitempty"`  //
	UnionId *string `json:"union_id,omitempty"` //
}

type UserIdBuilder struct {
	userId      string //
	userIdFlag  bool
	openId      string //
	openIdFlag  bool
	unionId     string //
	unionIdFlag bool
}

func NewUserIdBuilder() *UserIdBuilder {
	builder := &UserIdBuilder{}
	return builder
}

//
//
// 示例值：
func (builder *UserIdBuilder) UserId(userId string) *UserIdBuilder {
	builder.userId = userId
	builder.userIdFlag = true
	return builder
}

//
//
// 示例值：
func (builder *UserIdBuilder) OpenId(openId string) *UserIdBuilder {
	builder.openId = openId
	builder.openIdFlag = true
	return builder
}

//
//
// 示例值：
func (builder *UserIdBuilder) UnionId(unionId string) *UserIdBuilder {
	builder.unionId = unionId
	builder.unionIdFlag = true
	return builder
}

func (builder *UserIdBuilder) Build() *UserId {
	req := &UserId{}
	if builder.userIdFlag {
		req.UserId = &builder.userId

	}
	if builder.openIdFlag {
		req.OpenId = &builder.openId

	}
	if builder.unionIdFlag {
		req.UnionId = &builder.unionId

	}
	return req
}

type ListAccessRecordReqBuilder struct {
	apiReq *larkcore.ApiReq
	limit  int // 最大返回多少记录，当使用迭代器访问时才有效
}

func NewListAccessRecordReqBuilder() *ListAccessRecordReqBuilder {
	builder := &ListAccessRecordReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 最大返回多少记录，当使用迭代器访问时才有效
func (builder *ListAccessRecordReqBuilder) Limit(limit int) *ListAccessRecordReqBuilder {
	builder.limit = limit
	return builder
}

// 分页大小
//
// 示例值：100
func (builder *ListAccessRecordReqBuilder) PageSize(pageSize int) *ListAccessRecordReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

// 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果
//
// 示例值：AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw=
func (builder *ListAccessRecordReqBuilder) PageToken(pageToken string) *ListAccessRecordReqBuilder {
	builder.apiReq.QueryParams.Set("page_token", fmt.Sprint(pageToken))
	return builder
}

// 记录开始时间，单位秒
//
// 示例值：1624520521
func (builder *ListAccessRecordReqBuilder) From(from int) *ListAccessRecordReqBuilder {
	builder.apiReq.QueryParams.Set("from", fmt.Sprint(from))
	return builder
}

// 记录结束时间，单位秒，;时间跨度不能超过 30 天
//
// 示例值：1624520521
func (builder *ListAccessRecordReqBuilder) To(to int) *ListAccessRecordReqBuilder {
	builder.apiReq.QueryParams.Set("to", fmt.Sprint(to))
	return builder
}

// 门禁设备 ID
//
// 示例值：7091146989218002577
func (builder *ListAccessRecordReqBuilder) DeviceId(deviceId string) *ListAccessRecordReqBuilder {
	builder.apiReq.QueryParams.Set("device_id", fmt.Sprint(deviceId))
	return builder
}

// 此次调用中使用的用户 ID 的类型
//
// 示例值：
func (builder *ListAccessRecordReqBuilder) UserIdType(userIdType string) *ListAccessRecordReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

func (builder *ListAccessRecordReqBuilder) Build() *ListAccessRecordReq {
	req := &ListAccessRecordReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.Limit = builder.limit
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	return req
}

type ListAccessRecordReq struct {
	apiReq *larkcore.ApiReq
	Limit  int // 最多返回多少记录，只有在使用迭代器访问时，才有效

}

type ListAccessRecordRespData struct {
	Items     []*AccessRecord `json:"items,omitempty"`      // -
	PageToken *string         `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	HasMore   *bool           `json:"has_more,omitempty"`   // 是否还有更多项
}

type ListAccessRecordResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *ListAccessRecordRespData `json:"data"` // 业务数据
}

func (resp *ListAccessRecordResp) Success() bool {
	return resp.Code == 0
}

type GetAccessRecordAccessPhotoReqBuilder struct {
	apiReq *larkcore.ApiReq
}

func NewGetAccessRecordAccessPhotoReqBuilder() *GetAccessRecordAccessPhotoReqBuilder {
	builder := &GetAccessRecordAccessPhotoReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 门禁访问记录 ID
//
// 示例值：6939433228970082591
func (builder *GetAccessRecordAccessPhotoReqBuilder) AccessRecordId(accessRecordId string) *GetAccessRecordAccessPhotoReqBuilder {
	builder.apiReq.PathParams.Set("access_record_id", fmt.Sprint(accessRecordId))
	return builder
}

func (builder *GetAccessRecordAccessPhotoReqBuilder) Build() *GetAccessRecordAccessPhotoReq {
	req := &GetAccessRecordAccessPhotoReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	return req
}

type GetAccessRecordAccessPhotoReq struct {
	apiReq *larkcore.ApiReq
}

type GetAccessRecordAccessPhotoResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	File     io.Reader `json:"-"`
	FileName string    `json:"-"`
}

func (resp *GetAccessRecordAccessPhotoResp) Success() bool {
	return resp.Code == 0
}

func (resp *GetAccessRecordAccessPhotoResp) WriteFile(fileName string) error {
	bs, err := ioutil.ReadAll(resp.File)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, bs, 0666)
	if err != nil {
		return err
	}
	return nil
}

type ListDeviceRespData struct {
	Items []*Device `json:"items,omitempty"` // -
}

type ListDeviceResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *ListDeviceRespData `json:"data"` // 业务数据
}

func (resp *ListDeviceResp) Success() bool {
	return resp.Code == 0
}

type GetUserReqBuilder struct {
	apiReq *larkcore.ApiReq
}

func NewGetUserReqBuilder() *GetUserReqBuilder {
	builder := &GetUserReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 用户 ID
//
// 示例值：ou_7dab8a3d3cdcc9da365777c7ad535d62
func (builder *GetUserReqBuilder) UserId(userId string) *GetUserReqBuilder {
	builder.apiReq.PathParams.Set("user_id", fmt.Sprint(userId))
	return builder
}

// 此次调用中使用的用户 ID 的类型
//
// 示例值：
func (builder *GetUserReqBuilder) UserIdType(userIdType string) *GetUserReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

func (builder *GetUserReqBuilder) Build() *GetUserReq {
	req := &GetUserReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	return req
}

type GetUserReq struct {
	apiReq *larkcore.ApiReq
}

type GetUserRespData struct {
	User *User `json:"user,omitempty"` // 门禁用户信息
}

type GetUserResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *GetUserRespData `json:"data"` // 业务数据
}

func (resp *GetUserResp) Success() bool {
	return resp.Code == 0
}

type ListUserReqBuilder struct {
	apiReq *larkcore.ApiReq
	limit  int // 最大返回多少记录，当使用迭代器访问时才有效
}

func NewListUserReqBuilder() *ListUserReqBuilder {
	builder := &ListUserReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 最大返回多少记录，当使用迭代器访问时才有效
func (builder *ListUserReqBuilder) Limit(limit int) *ListUserReqBuilder {
	builder.limit = limit
	return builder
}

// 分页大小
//
// 示例值：10
func (builder *ListUserReqBuilder) PageSize(pageSize int) *ListUserReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

// 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果
//
// 示例值：10
func (builder *ListUserReqBuilder) PageToken(pageToken string) *ListUserReqBuilder {
	builder.apiReq.QueryParams.Set("page_token", fmt.Sprint(pageToken))
	return builder
}

// 此次调用中使用的用户 ID 的类型
//
// 示例值：
func (builder *ListUserReqBuilder) UserIdType(userIdType string) *ListUserReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

func (builder *ListUserReqBuilder) Build() *ListUserReq {
	req := &ListUserReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.Limit = builder.limit
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	return req
}

type ListUserReq struct {
	apiReq *larkcore.ApiReq
	Limit  int // 最多返回多少记录，只有在使用迭代器访问时，才有效

}

type ListUserRespData struct {
	Items     []*User `json:"items,omitempty"`      // -
	PageToken *string `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	HasMore   *bool   `json:"has_more,omitempty"`   // 是否还有更多项
}

type ListUserResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *ListUserRespData `json:"data"` // 业务数据
}

func (resp *ListUserResp) Success() bool {
	return resp.Code == 0
}

type PatchUserReqBuilder struct {
	apiReq *larkcore.ApiReq
	user   *User
}

func NewPatchUserReqBuilder() *PatchUserReqBuilder {
	builder := &PatchUserReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 用户 ID
//
// 示例值：ou_7dab8a3d3cdcc9da365777c7ad535d62
func (builder *PatchUserReqBuilder) UserId(userId string) *PatchUserReqBuilder {
	builder.apiReq.PathParams.Set("user_id", fmt.Sprint(userId))
	return builder
}

// 此次调用中使用的用户 ID 的类型
//
// 示例值：
func (builder *PatchUserReqBuilder) UserIdType(userIdType string) *PatchUserReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

// 飞书智能门禁在人脸识别成功后会有韦根信号输出，输出用户的卡号。;对于使用韦根协议的门禁系统，企业可使用该接口录入用户卡号。
func (builder *PatchUserReqBuilder) User(user *User) *PatchUserReqBuilder {
	builder.user = user
	return builder
}

func (builder *PatchUserReqBuilder) Build() *PatchUserReq {
	req := &PatchUserReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	req.apiReq.Body = builder.user
	return req
}

type PatchUserReq struct {
	apiReq *larkcore.ApiReq
	User   *User `body:""`
}

type PatchUserResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
}

func (resp *PatchUserResp) Success() bool {
	return resp.Code == 0
}

type GetUserFaceReqBuilder struct {
	apiReq *larkcore.ApiReq
}

func NewGetUserFaceReqBuilder() *GetUserFaceReqBuilder {
	builder := &GetUserFaceReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 用户 ID
//
// 示例值：ou_7dab8a3d3cdcc9da365777c7ad535d62
func (builder *GetUserFaceReqBuilder) UserId(userId string) *GetUserFaceReqBuilder {
	builder.apiReq.PathParams.Set("user_id", fmt.Sprint(userId))
	return builder
}

// 裁剪图
//
// 示例值：true
func (builder *GetUserFaceReqBuilder) IsCropped(isCropped bool) *GetUserFaceReqBuilder {
	builder.apiReq.QueryParams.Set("is_cropped", fmt.Sprint(isCropped))
	return builder
}

// 此次调用中使用的用户 ID 的类型
//
// 示例值：
func (builder *GetUserFaceReqBuilder) UserIdType(userIdType string) *GetUserFaceReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

func (builder *GetUserFaceReqBuilder) Build() *GetUserFaceReq {
	req := &GetUserFaceReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	return req
}

type GetUserFaceReq struct {
	apiReq *larkcore.ApiReq
}

type GetUserFaceResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	File     io.Reader `json:"-"`
	FileName string    `json:"-"`
}

func (resp *GetUserFaceResp) Success() bool {
	return resp.Code == 0
}

func (resp *GetUserFaceResp) WriteFile(fileName string) error {
	bs, err := ioutil.ReadAll(resp.File)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, bs, 0666)
	if err != nil {
		return err
	}
	return nil
}

type UpdateUserFaceReqBuilder struct {
	apiReq *larkcore.ApiReq
	file   *File
}

func NewUpdateUserFaceReqBuilder() *UpdateUserFaceReqBuilder {
	builder := &UpdateUserFaceReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 用户 ID
//
// 示例值：ou_7dab8a3d3cdcc9da365777c7ad535d62
func (builder *UpdateUserFaceReqBuilder) UserId(userId string) *UpdateUserFaceReqBuilder {
	builder.apiReq.PathParams.Set("user_id", fmt.Sprint(userId))
	return builder
}

// 此次调用中使用的用户 ID 的类型
//
// 示例值：
func (builder *UpdateUserFaceReqBuilder) UserIdType(userIdType string) *UpdateUserFaceReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}

// 用户需要录入人脸图片才可以使用门禁考勤机。使用该 API 上传门禁用户的人脸图片。
func (builder *UpdateUserFaceReqBuilder) File(file *File) *UpdateUserFaceReqBuilder {
	builder.file = file
	return builder
}

func (builder *UpdateUserFaceReqBuilder) Build() *UpdateUserFaceReq {
	req := &UpdateUserFaceReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	req.apiReq.Body = builder.file
	return req
}

type UpdateUserFaceReq struct {
	apiReq *larkcore.ApiReq
	File   *File `body:""`
}

type UpdateUserFaceResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
}

func (resp *UpdateUserFaceResp) Success() bool {
	return resp.Code == 0
}

type P2AccessRecordCreatedV1Data struct {
	AccessRecordId *string `json:"access_record_id,omitempty"` // 门禁记录 ID
	UserId         *UserId `json:"user_id,omitempty"`          // 用户 ID
	DeviceId       *string `json:"device_id,omitempty"`        // 设备 ID
	IsClockIn      *bool   `json:"is_clock_in,omitempty"`      // 是否打卡
	IsDoorOpen     *bool   `json:"is_door_open,omitempty"`     // 是否开门
	AccessTime     *string `json:"access_time,omitempty"`      // 识别时间（单位：秒）
}

type P2AccessRecordCreatedV1 struct {
	*larkevent.EventV2Base                              // 事件基础数据
	*larkevent.EventReq                                 // 请求原生数据
	Event                  *P2AccessRecordCreatedV1Data `json:"event"` // 事件内容
}

func (m *P2AccessRecordCreatedV1) RawReq(req *larkevent.EventReq) {
	m.EventReq = req
}

type P2UserUpdatedV1Data struct {
	UserId       *UserId `json:"user_id,omitempty"`       // 用户 ID
	Card         *int    `json:"card,omitempty"`          // 卡号
	FaceUploaded *bool   `json:"face_uploaded,omitempty"` // 是否上传人脸图片
}

type P2UserUpdatedV1 struct {
	*larkevent.EventV2Base                      // 事件基础数据
	*larkevent.EventReq                         // 请求原生数据
	Event                  *P2UserUpdatedV1Data `json:"event"` // 事件内容
}

func (m *P2UserUpdatedV1) RawReq(req *larkevent.EventReq) {
	m.EventReq = req
}

type ListAccessRecordIterator struct {
	nextPageToken *string
	items         []*AccessRecord
	index         int
	limit         int
	ctx           context.Context
	req           *ListAccessRecordReq
	listFunc      func(ctx context.Context, req *ListAccessRecordReq, options ...larkcore.RequestOptionFunc) (*ListAccessRecordResp, error)
	options       []larkcore.RequestOptionFunc
	curlNum       int
}

func (iterator *ListAccessRecordIterator) Next() (bool, *AccessRecord, error) {
	// 达到最大量，则返回
	if iterator.limit > 0 && iterator.curlNum >= iterator.limit {
		return false, nil, nil
	}

	// 为 0 则拉取数据
	if iterator.index == 0 || iterator.index >= len(iterator.items) {
		if iterator.index != 0 && iterator.nextPageToken == nil {
			return false, nil, nil
		}
		if iterator.nextPageToken != nil {
			iterator.req.apiReq.QueryParams.Set("page_token", *iterator.nextPageToken)
		}
		resp, err := iterator.listFunc(iterator.ctx, iterator.req, iterator.options...)
		if err != nil {
			return false, nil, err
		}

		if resp.Code != 0 {
			return false, nil, errors.New(fmt.Sprintf("Code:%d,Msg:%s", resp.Code, resp.Msg))
		}

		if len(resp.Data.Items) == 0 {
			return false, nil, nil
		}

		iterator.nextPageToken = resp.Data.PageToken
		iterator.items = resp.Data.Items
		iterator.index = 0
	}

	block := iterator.items[iterator.index]
	iterator.index++
	iterator.curlNum++
	return true, block, nil
}

func (iterator *ListAccessRecordIterator) NextPageToken() *string {
	return iterator.nextPageToken
}

type ListUserIterator struct {
	nextPageToken *string
	items         []*User
	index         int
	limit         int
	ctx           context.Context
	req           *ListUserReq
	listFunc      func(ctx context.Context, req *ListUserReq, options ...larkcore.RequestOptionFunc) (*ListUserResp, error)
	options       []larkcore.RequestOptionFunc
	curlNum       int
}

func (iterator *ListUserIterator) Next() (bool, *User, error) {
	// 达到最大量，则返回
	if iterator.limit > 0 && iterator.curlNum >= iterator.limit {
		return false, nil, nil
	}

	// 为 0 则拉取数据
	if iterator.index == 0 || iterator.index >= len(iterator.items) {
		if iterator.index != 0 && iterator.nextPageToken == nil {
			return false, nil, nil
		}
		if iterator.nextPageToken != nil {
			iterator.req.apiReq.QueryParams.Set("page_token", *iterator.nextPageToken)
		}
		resp, err := iterator.listFunc(iterator.ctx, iterator.req, iterator.options...)
		if err != nil {
			return false, nil, err
		}

		if resp.Code != 0 {
			return false, nil, errors.New(fmt.Sprintf("Code:%d,Msg:%s", resp.Code, resp.Msg))
		}

		if len(resp.Data.Items) == 0 {
			return false, nil, nil
		}

		iterator.nextPageToken = resp.Data.PageToken
		iterator.items = resp.Data.Items
		iterator.index = 0
	}

	block := iterator.items[iterator.index]
	iterator.index++
	iterator.curlNum++
	return true, block, nil
}

func (iterator *ListUserIterator) NextPageToken() *string {
	return iterator.nextPageToken
}
