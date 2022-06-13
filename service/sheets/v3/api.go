// Package sheets code generated by oapi sdk gen
package sheets

import (
	"net/http"
	"context"
	
	"github.com/feishu/oapi-sdk-go/core"
)

/**
构建业务域服务实例
**/
func NewService(httpClient *http.Client, config *core.Config) *SheetsService {
	s := &SheetsService{httpClient:httpClient,config:config}
	s.Spreadsheets = &spreadsheets{service: s}
	s.Sheets = &sheets{service: s}
	s.SpreadsheetSheetFilter = &spreadsheetSheetFilter{service: s}
	s.FilterViews = &filterViews{service: s}
	s.Conditions = &conditions{service: s}
	s.FloatImages = &floatImages{service: s}
	return s
}

/**
业务域服务定义
**/
type SheetsService struct {
	httpClient *http.Client
	config *core.Config
	Spreadsheets *spreadsheets
	Sheets *sheets
	SpreadsheetSheetFilter *spreadsheetSheetFilter
	FilterViews *filterViews
	Conditions *conditions
	FloatImages *floatImages
}


/**
资源服务定义
**/
type spreadsheets struct {
   service *SheetsService
}
type sheets struct {
   service *SheetsService
}
type spreadsheetSheetFilter struct {
   service *SheetsService
}
type filterViews struct {
   service *SheetsService
}
type conditions struct {
   service *SheetsService
}
type floatImages struct {
   service *SheetsService
}
/**
资源服务方法定义
**/
func (s *spreadsheets) Create(ctx context.Context, req *CreateSpreadsheetReq, options ...core.RequestOptionFunc) (*CreateSpreadsheetResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateSpreadsheetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *sheets) Find(ctx context.Context, req *FindSpreadsheetSheetReq, options ...core.RequestOptionFunc) (*FindSpreadsheetSheetResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/find", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &FindSpreadsheetSheetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *sheets) MoveDimension(ctx context.Context, req *MoveDimensionSpreadsheetSheetReq, options ...core.RequestOptionFunc) (*MoveDimensionSpreadsheetSheetResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/move_dimension", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &MoveDimensionSpreadsheetSheetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *sheets) Replace(ctx context.Context, req *ReplaceSpreadsheetSheetReq, options ...core.RequestOptionFunc) (*ReplaceSpreadsheetSheetResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/replace", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &ReplaceSpreadsheetSheetResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilter) Create(ctx context.Context, req *CreateSpreadsheetSheetFilterReq, options ...core.RequestOptionFunc) (*CreateSpreadsheetSheetFilterResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,s.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFilterResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilter) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFilterReq, options ...core.RequestOptionFunc) (*DeleteSpreadsheetSheetFilterResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,s.service.config, http.MethodDelete,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFilterResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilter) Get(ctx context.Context, req *GetSpreadsheetSheetFilterReq, options ...core.RequestOptionFunc) (*GetSpreadsheetSheetFilterResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,s.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &GetSpreadsheetSheetFilterResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spreadsheetSheetFilter) Update(ctx context.Context, req *UpdateSpreadsheetSheetFilterReq, options ...core.RequestOptionFunc) (*UpdateSpreadsheetSheetFilterResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,s.service.config, http.MethodPut,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &UpdateSpreadsheetSheetFilterResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *filterViews) Create(ctx context.Context, req *CreateSpreadsheetSheetFilterViewReq, options ...core.RequestOptionFunc) (*CreateSpreadsheetSheetFilterViewResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *filterViews) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFilterViewReq, options ...core.RequestOptionFunc) (*DeleteSpreadsheetSheetFilterViewResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodDelete,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *filterViews) Get(ctx context.Context, req *GetSpreadsheetSheetFilterViewReq, options ...core.RequestOptionFunc) (*GetSpreadsheetSheetFilterViewResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &GetSpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *filterViews) Patch(ctx context.Context, req *PatchSpreadsheetSheetFilterViewReq, options ...core.RequestOptionFunc) (*PatchSpreadsheetSheetFilterViewResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodPatch,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &PatchSpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *filterViews) Query(ctx context.Context, req *QuerySpreadsheetSheetFilterViewReq, options ...core.RequestOptionFunc) (*QuerySpreadsheetSheetFilterViewResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/query", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &QuerySpreadsheetSheetFilterViewResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *conditions) Create(ctx context.Context, req *CreateSpreadsheetSheetFilterViewConditionReq, options ...core.RequestOptionFunc) (*CreateSpreadsheetSheetFilterViewConditionResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *conditions) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFilterViewConditionReq, options ...core.RequestOptionFunc) (*DeleteSpreadsheetSheetFilterViewConditionResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodDelete,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *conditions) Get(ctx context.Context, req *GetSpreadsheetSheetFilterViewConditionReq, options ...core.RequestOptionFunc) (*GetSpreadsheetSheetFilterViewConditionResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &GetSpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *conditions) Query(ctx context.Context, req *QuerySpreadsheetSheetFilterViewConditionReq, options ...core.RequestOptionFunc) (*QuerySpreadsheetSheetFilterViewConditionResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/query", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &QuerySpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *conditions) Update(ctx context.Context, req *UpdateSpreadsheetSheetFilterViewConditionReq, options ...core.RequestOptionFunc) (*UpdateSpreadsheetSheetFilterViewConditionResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPut,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &UpdateSpreadsheetSheetFilterViewConditionResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *floatImages) Create(ctx context.Context, req *CreateSpreadsheetSheetFloatImageReq, options ...core.RequestOptionFunc) (*CreateSpreadsheetSheetFloatImageResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodPost,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *floatImages) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFloatImageReq, options ...core.RequestOptionFunc) (*DeleteSpreadsheetSheetFloatImageResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodDelete,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *floatImages) Get(ctx context.Context, req *GetSpreadsheetSheetFloatImageReq, options ...core.RequestOptionFunc) (*GetSpreadsheetSheetFloatImageResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &GetSpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *floatImages) Patch(ctx context.Context, req *PatchSpreadsheetSheetFloatImageReq, options ...core.RequestOptionFunc) (*PatchSpreadsheetSheetFloatImageResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodPatch,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &PatchSpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *floatImages) Query(ctx context.Context, req *QuerySpreadsheetSheetFloatImageReq, options ...core.RequestOptionFunc) (*QuerySpreadsheetSheetFloatImageResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodGet,
		"/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/query", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &QuerySpreadsheetSheetFloatImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}