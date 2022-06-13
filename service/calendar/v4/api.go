// Package calendar code generated by oapi sdk gen
package calendar

import (
	"net/http"
	"context"
	
	"github.com/feishu/oapi-sdk-go/core"
)

/**
构建业务域服务实例
**/
func NewService(httpClient *http.Client, config *core.Config) *CalendarService {
	c := &CalendarService{httpClient:httpClient,config:config}
	c.Calendars = &calendars{service: c}
	c.CalendarAcls = &calendarAcls{service: c}
	c.CalendarEvents = &calendarEvents{service: c}
	c.CalendarEventAttendees = &calendarEventAttendees{service: c}
	c.CalendarEventAttendeeChatMembers = &calendarEventAttendeeChatMembers{service: c}
	c.ExchangeBindings = &exchangeBindings{service: c}
	c.Freebusy = &freebusy{service: c}
	c.Settings = &settings{service: c}
	c.TimeoffEvents = &timeoffEvents{service: c}
	return c
}

/**
业务域服务定义
**/
type CalendarService struct {
	httpClient *http.Client
	config *core.Config
	Calendars *calendars
	CalendarAcls *calendarAcls
	CalendarEvents *calendarEvents
	CalendarEventAttendees *calendarEventAttendees
	CalendarEventAttendeeChatMembers *calendarEventAttendeeChatMembers
	ExchangeBindings *exchangeBindings
	Freebusy *freebusy
	Settings *settings
	TimeoffEvents *timeoffEvents
}


/**
资源服务定义
**/
type calendars struct {
   service *CalendarService
}
type calendarAcls struct {
   service *CalendarService
}
type calendarEvents struct {
   service *CalendarService
}
type calendarEventAttendees struct {
   service *CalendarService
}
type calendarEventAttendeeChatMembers struct {
   service *CalendarService
}
type exchangeBindings struct {
   service *CalendarService
}
type freebusy struct {
   service *CalendarService
}
type settings struct {
   service *CalendarService
}
type timeoffEvents struct {
   service *CalendarService
}
/**
资源服务方法定义
**/
func (c *calendars) Create(ctx context.Context, req *CreateCalendarReq, options ...core.RequestOptionFunc) (*CreateCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendars) Delete(ctx context.Context, req *DeleteCalendarReq, options ...core.RequestOptionFunc) (*DeleteCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/calendars/:calendar_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &DeleteCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendars) Get(ctx context.Context, req *GetCalendarReq, options ...core.RequestOptionFunc) (*GetCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &GetCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendars) List(ctx context.Context, req *ListCalendarReq, options ...core.RequestOptionFunc) (*ListCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &ListCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendars) Patch(ctx context.Context, req *PatchCalendarReq, options ...core.RequestOptionFunc) (*PatchCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPatch,
		"/open-apis/calendar/v4/calendars/:calendar_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &PatchCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendars) Primary(ctx context.Context, req *PrimaryCalendarReq, options ...core.RequestOptionFunc) (*PrimaryCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/primary", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &PrimaryCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendars) Search(ctx context.Context, req *SearchCalendarReq, options ...core.RequestOptionFunc) (*SearchCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/search", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &SearchCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
/**如果是分页查询，则添加迭代器函数**/
func (c *calendars) SearchCalendar(ctx context.Context, req *SearchCalendarReq, options ...core.RequestOptionFunc) (*SearchCalendarIterator, error) {
   return &SearchCalendarIterator{
	  ctx:	  ctx,
	  req:	  req,
	  listFunc: c.Search,
	  options:  options}, nil
}
func (c *calendars) Subscribe(ctx context.Context, req *SubscribeCalendarReq, options ...core.RequestOptionFunc) (*SubscribeCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/subscribe", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &SubscribeCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendars) Subscription(ctx context.Context,  options ...core.RequestOptionFunc) (*SubscriptionCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/subscription", []core.AccessTokenType{core.AccessTokenTypeUser}, nil, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &SubscriptionCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendars) Unsubscribe(ctx context.Context, req *UnsubscribeCalendarReq, options ...core.RequestOptionFunc) (*UnsubscribeCalendarResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &UnsubscribeCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcls) Create(ctx context.Context, req *CreateCalendarAclReq, options ...core.RequestOptionFunc) (*CreateCalendarAclResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/acls", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcls) Delete(ctx context.Context, req *DeleteCalendarAclReq, options ...core.RequestOptionFunc) (*DeleteCalendarAclResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &DeleteCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcls) List(ctx context.Context, req *ListCalendarAclReq, options ...core.RequestOptionFunc) (*ListCalendarAclResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/acls", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &ListCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
/**如果是分页查询，则添加迭代器函数**/
func (c *calendarAcls) ListCalendarAcl(ctx context.Context, req *ListCalendarAclReq, options ...core.RequestOptionFunc) (*ListCalendarAclIterator, error) {
   return &ListCalendarAclIterator{
	  ctx:	  ctx,
	  req:	  req,
	  listFunc: c.List,
	  options:  options}, nil
}
func (c *calendarAcls) Subscription(ctx context.Context, req *SubscriptionCalendarAclReq, options ...core.RequestOptionFunc) (*SubscriptionCalendarAclResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/acls/subscription", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &SubscriptionCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvents) Create(ctx context.Context, req *CreateCalendarEventReq, options ...core.RequestOptionFunc) (*CreateCalendarEventResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvents) Delete(ctx context.Context, req *DeleteCalendarEventReq, options ...core.RequestOptionFunc) (*DeleteCalendarEventResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &DeleteCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvents) Get(ctx context.Context, req *GetCalendarEventReq, options ...core.RequestOptionFunc) (*GetCalendarEventResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &GetCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvents) List(ctx context.Context, req *ListCalendarEventReq, options ...core.RequestOptionFunc) (*ListCalendarEventResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/events", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &ListCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvents) Patch(ctx context.Context, req *PatchCalendarEventReq, options ...core.RequestOptionFunc) (*PatchCalendarEventResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPatch,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &PatchCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvents) Search(ctx context.Context, req *SearchCalendarEventReq, options ...core.RequestOptionFunc) (*SearchCalendarEventResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/search", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &SearchCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
/**如果是分页查询，则添加迭代器函数**/
func (c *calendarEvents) SearchCalendarEvent(ctx context.Context, req *SearchCalendarEventReq, options ...core.RequestOptionFunc) (*SearchCalendarEventIterator, error) {
   return &SearchCalendarEventIterator{
	  ctx:	  ctx,
	  req:	  req,
	  listFunc: c.Search,
	  options:  options}, nil
}
func (c *calendarEvents) Subscription(ctx context.Context, req *SubscriptionCalendarEventReq, options ...core.RequestOptionFunc) (*SubscriptionCalendarEventResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/subscription", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &SubscriptionCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendees) BatchDelete(ctx context.Context, req *BatchDeleteCalendarEventAttendeeReq, options ...core.RequestOptionFunc) (*BatchDeleteCalendarEventAttendeeResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/batch_delete", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &BatchDeleteCalendarEventAttendeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendees) Create(ctx context.Context, req *CreateCalendarEventAttendeeReq, options ...core.RequestOptionFunc) (*CreateCalendarEventAttendeeResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateCalendarEventAttendeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendees) List(ctx context.Context, req *ListCalendarEventAttendeeReq, options ...core.RequestOptionFunc) (*ListCalendarEventAttendeeResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &ListCalendarEventAttendeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
/**如果是分页查询，则添加迭代器函数**/
func (c *calendarEventAttendees) ListCalendarEventAttendee(ctx context.Context, req *ListCalendarEventAttendeeReq, options ...core.RequestOptionFunc) (*ListCalendarEventAttendeeIterator, error) {
   return &ListCalendarEventAttendeeIterator{
	  ctx:	  ctx,
	  req:	  req,
	  listFunc: c.List,
	  options:  options}, nil
}
func (c *calendarEventAttendeeChatMembers) List(ctx context.Context, req *ListCalendarEventAttendeeChatMemberReq, options ...core.RequestOptionFunc) (*ListCalendarEventAttendeeChatMemberResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/:attendee_id/chat_members", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &ListCalendarEventAttendeeChatMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
/**如果是分页查询，则添加迭代器函数**/
func (c *calendarEventAttendeeChatMembers) ListCalendarEventAttendeeChatMember(ctx context.Context, req *ListCalendarEventAttendeeChatMemberReq, options ...core.RequestOptionFunc) (*ListCalendarEventAttendeeChatMemberIterator, error) {
   return &ListCalendarEventAttendeeChatMemberIterator{
	  ctx:	  ctx,
	  req:	  req,
	  listFunc: c.List,
	  options:  options}, nil
}
func (e *exchangeBindings) Create(ctx context.Context, req *CreateExchangeBindingReq, options ...core.RequestOptionFunc) (*CreateExchangeBindingResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,e.service.config, http.MethodPost,
		"/open-apis/calendar/v4/exchange_bindings", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateExchangeBindingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exchangeBindings) Delete(ctx context.Context, req *DeleteExchangeBindingReq, options ...core.RequestOptionFunc) (*DeleteExchangeBindingResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,e.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &DeleteExchangeBindingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exchangeBindings) Get(ctx context.Context, req *GetExchangeBindingReq, options ...core.RequestOptionFunc) (*GetExchangeBindingResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,e.service.config, http.MethodGet,
		"/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &GetExchangeBindingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *freebusy) BatchGet(ctx context.Context, req *BatchGetFreebusyReq, options ...core.RequestOptionFunc) (*BatchGetFreebusyResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodPost,
		"/open-apis/calendar/v4/freebusy/batch_get", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &BatchGetFreebusyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *freebusy) List(ctx context.Context, req *ListFreebusyReq, options ...core.RequestOptionFunc) (*ListFreebusyResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,f.service.config, http.MethodPost,
		"/open-apis/calendar/v4/freebusy/list", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &ListFreebusyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *settings) GenerateCaldavConf(ctx context.Context, req *GenerateCaldavConfSettingReq, options ...core.RequestOptionFunc) (*GenerateCaldavConfSettingResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,s.service.config, http.MethodPost,
		"/open-apis/calendar/v4/settings/generate_caldav_conf", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &GenerateCaldavConfSettingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *timeoffEvents) Create(ctx context.Context, req *CreateTimeoffEventReq, options ...core.RequestOptionFunc) (*CreateTimeoffEventResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,t.service.config, http.MethodPost,
		"/open-apis/calendar/v4/timeoff_events", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &CreateTimeoffEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *timeoffEvents) Delete(ctx context.Context, req *DeleteTimeoffEventReq, options ...core.RequestOptionFunc) (*DeleteTimeoffEventResp, error) {

	// 发起请求
	rawResp, err := core.SendRequest(ctx,t.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/timeoff_events/:timeoff_event_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}

	// 反序列响应结果
	resp := &DeleteTimeoffEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}