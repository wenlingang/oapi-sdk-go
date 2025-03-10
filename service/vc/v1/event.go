// Package vc code generated by oapi sdk gen
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

package larkvc

import (
	"context"
)

// 消息处理器定义
type P2MeetingAllMeetingEndedV1Handler struct {
	handler func(context.Context, *P2MeetingAllMeetingEndedV1) error
}

func NewP2MeetingAllMeetingEndedV1Handler(handler func(context.Context, *P2MeetingAllMeetingEndedV1) error) *P2MeetingAllMeetingEndedV1Handler {
	h := &P2MeetingAllMeetingEndedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingAllMeetingEndedV1Handler) Event() interface{} {
	return &P2MeetingAllMeetingEndedV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingAllMeetingEndedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingAllMeetingEndedV1))
}

// 消息处理器定义
type P2MeetingAllMeetingStartedV1Handler struct {
	handler func(context.Context, *P2MeetingAllMeetingStartedV1) error
}

func NewP2MeetingAllMeetingStartedV1Handler(handler func(context.Context, *P2MeetingAllMeetingStartedV1) error) *P2MeetingAllMeetingStartedV1Handler {
	h := &P2MeetingAllMeetingStartedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingAllMeetingStartedV1Handler) Event() interface{} {
	return &P2MeetingAllMeetingStartedV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingAllMeetingStartedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingAllMeetingStartedV1))
}

// 消息处理器定义
type P2MeetingJoinMeetingV1Handler struct {
	handler func(context.Context, *P2MeetingJoinMeetingV1) error
}

func NewP2MeetingJoinMeetingV1Handler(handler func(context.Context, *P2MeetingJoinMeetingV1) error) *P2MeetingJoinMeetingV1Handler {
	h := &P2MeetingJoinMeetingV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingJoinMeetingV1Handler) Event() interface{} {
	return &P2MeetingJoinMeetingV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingJoinMeetingV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingJoinMeetingV1))
}

// 消息处理器定义
type P2MeetingLeaveMeetingV1Handler struct {
	handler func(context.Context, *P2MeetingLeaveMeetingV1) error
}

func NewP2MeetingLeaveMeetingV1Handler(handler func(context.Context, *P2MeetingLeaveMeetingV1) error) *P2MeetingLeaveMeetingV1Handler {
	h := &P2MeetingLeaveMeetingV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingLeaveMeetingV1Handler) Event() interface{} {
	return &P2MeetingLeaveMeetingV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingLeaveMeetingV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingLeaveMeetingV1))
}

// 消息处理器定义
type P2MeetingEndedV1Handler struct {
	handler func(context.Context, *P2MeetingEndedV1) error
}

func NewP2MeetingEndedV1Handler(handler func(context.Context, *P2MeetingEndedV1) error) *P2MeetingEndedV1Handler {
	h := &P2MeetingEndedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingEndedV1Handler) Event() interface{} {
	return &P2MeetingEndedV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingEndedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingEndedV1))
}

// 消息处理器定义
type P2MeetingStartedV1Handler struct {
	handler func(context.Context, *P2MeetingStartedV1) error
}

func NewP2MeetingStartedV1Handler(handler func(context.Context, *P2MeetingStartedV1) error) *P2MeetingStartedV1Handler {
	h := &P2MeetingStartedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingStartedV1Handler) Event() interface{} {
	return &P2MeetingStartedV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingStartedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingStartedV1))
}

// 消息处理器定义
type P2MeetingRecordingEndedV1Handler struct {
	handler func(context.Context, *P2MeetingRecordingEndedV1) error
}

func NewP2MeetingRecordingEndedV1Handler(handler func(context.Context, *P2MeetingRecordingEndedV1) error) *P2MeetingRecordingEndedV1Handler {
	h := &P2MeetingRecordingEndedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingRecordingEndedV1Handler) Event() interface{} {
	return &P2MeetingRecordingEndedV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingRecordingEndedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingRecordingEndedV1))
}

// 消息处理器定义
type P2MeetingRecordingReadyV1Handler struct {
	handler func(context.Context, *P2MeetingRecordingReadyV1) error
}

func NewP2MeetingRecordingReadyV1Handler(handler func(context.Context, *P2MeetingRecordingReadyV1) error) *P2MeetingRecordingReadyV1Handler {
	h := &P2MeetingRecordingReadyV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingRecordingReadyV1Handler) Event() interface{} {
	return &P2MeetingRecordingReadyV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingRecordingReadyV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingRecordingReadyV1))
}

// 消息处理器定义
type P2MeetingRecordingStartedV1Handler struct {
	handler func(context.Context, *P2MeetingRecordingStartedV1) error
}

func NewP2MeetingRecordingStartedV1Handler(handler func(context.Context, *P2MeetingRecordingStartedV1) error) *P2MeetingRecordingStartedV1Handler {
	h := &P2MeetingRecordingStartedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingRecordingStartedV1Handler) Event() interface{} {
	return &P2MeetingRecordingStartedV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingRecordingStartedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingRecordingStartedV1))
}

// 消息处理器定义
type P2MeetingShareEndedV1Handler struct {
	handler func(context.Context, *P2MeetingShareEndedV1) error
}

func NewP2MeetingShareEndedV1Handler(handler func(context.Context, *P2MeetingShareEndedV1) error) *P2MeetingShareEndedV1Handler {
	h := &P2MeetingShareEndedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingShareEndedV1Handler) Event() interface{} {
	return &P2MeetingShareEndedV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingShareEndedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingShareEndedV1))
}

// 消息处理器定义
type P2MeetingShareStartedV1Handler struct {
	handler func(context.Context, *P2MeetingShareStartedV1) error
}

func NewP2MeetingShareStartedV1Handler(handler func(context.Context, *P2MeetingShareStartedV1) error) *P2MeetingShareStartedV1Handler {
	h := &P2MeetingShareStartedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2MeetingShareStartedV1Handler) Event() interface{} {
	return &P2MeetingShareStartedV1{}
}

// 回调开发者注册的 handle
func (h *P2MeetingShareStartedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2MeetingShareStartedV1))
}

// 消息处理器定义
type P2ReserveConfigUpdatedV1Handler struct {
	handler func(context.Context, *P2ReserveConfigUpdatedV1) error
}

func NewP2ReserveConfigUpdatedV1Handler(handler func(context.Context, *P2ReserveConfigUpdatedV1) error) *P2ReserveConfigUpdatedV1Handler {
	h := &P2ReserveConfigUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2ReserveConfigUpdatedV1Handler) Event() interface{} {
	return &P2ReserveConfigUpdatedV1{}
}

// 回调开发者注册的 handle
func (h *P2ReserveConfigUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2ReserveConfigUpdatedV1))
}

// 消息处理器定义
type P2RoomCreatedV1Handler struct {
	handler func(context.Context, *P2RoomCreatedV1) error
}

func NewP2RoomCreatedV1Handler(handler func(context.Context, *P2RoomCreatedV1) error) *P2RoomCreatedV1Handler {
	h := &P2RoomCreatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2RoomCreatedV1Handler) Event() interface{} {
	return &P2RoomCreatedV1{}
}

// 回调开发者注册的 handle
func (h *P2RoomCreatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2RoomCreatedV1))
}

// 消息处理器定义
type P2RoomDeletedV1Handler struct {
	handler func(context.Context, *P2RoomDeletedV1) error
}

func NewP2RoomDeletedV1Handler(handler func(context.Context, *P2RoomDeletedV1) error) *P2RoomDeletedV1Handler {
	h := &P2RoomDeletedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2RoomDeletedV1Handler) Event() interface{} {
	return &P2RoomDeletedV1{}
}

// 回调开发者注册的 handle
func (h *P2RoomDeletedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2RoomDeletedV1))
}

// 消息处理器定义
type P2RoomUpdatedV1Handler struct {
	handler func(context.Context, *P2RoomUpdatedV1) error
}

func NewP2RoomUpdatedV1Handler(handler func(context.Context, *P2RoomUpdatedV1) error) *P2RoomUpdatedV1Handler {
	h := &P2RoomUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2RoomUpdatedV1Handler) Event() interface{} {
	return &P2RoomUpdatedV1{}
}

// 回调开发者注册的 handle
func (h *P2RoomUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2RoomUpdatedV1))
}

// 消息处理器定义
type P2RoomLevelCreatedV1Handler struct {
	handler func(context.Context, *P2RoomLevelCreatedV1) error
}

func NewP2RoomLevelCreatedV1Handler(handler func(context.Context, *P2RoomLevelCreatedV1) error) *P2RoomLevelCreatedV1Handler {
	h := &P2RoomLevelCreatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2RoomLevelCreatedV1Handler) Event() interface{} {
	return &P2RoomLevelCreatedV1{}
}

// 回调开发者注册的 handle
func (h *P2RoomLevelCreatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2RoomLevelCreatedV1))
}

// 消息处理器定义
type P2RoomLevelDeletedV1Handler struct {
	handler func(context.Context, *P2RoomLevelDeletedV1) error
}

func NewP2RoomLevelDeletedV1Handler(handler func(context.Context, *P2RoomLevelDeletedV1) error) *P2RoomLevelDeletedV1Handler {
	h := &P2RoomLevelDeletedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2RoomLevelDeletedV1Handler) Event() interface{} {
	return &P2RoomLevelDeletedV1{}
}

// 回调开发者注册的 handle
func (h *P2RoomLevelDeletedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2RoomLevelDeletedV1))
}

// 消息处理器定义
type P2RoomLevelUpdatedV1Handler struct {
	handler func(context.Context, *P2RoomLevelUpdatedV1) error
}

func NewP2RoomLevelUpdatedV1Handler(handler func(context.Context, *P2RoomLevelUpdatedV1) error) *P2RoomLevelUpdatedV1Handler {
	h := &P2RoomLevelUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2RoomLevelUpdatedV1Handler) Event() interface{} {
	return &P2RoomLevelUpdatedV1{}
}

// 回调开发者注册的 handle
func (h *P2RoomLevelUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2RoomLevelUpdatedV1))
}
