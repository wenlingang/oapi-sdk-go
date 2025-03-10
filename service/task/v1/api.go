// Package task code generated by oapi sdk gen
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

package larktask

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *TaskService {
	t := &TaskService{config: config}
	t.Task = &task{service: t}
	t.TaskCollaborator = &taskCollaborator{service: t}
	t.TaskComment = &taskComment{service: t}
	t.TaskFollower = &taskFollower{service: t}
	t.TaskReminder = &taskReminder{service: t}
	return t
}

type TaskService struct {
	config           *larkcore.Config
	Task             *task             // 任务
	TaskCollaborator *taskCollaborator // 执行者
	TaskComment      *taskComment      // 评论
	TaskFollower     *taskFollower     // 关注人
	TaskReminder     *taskReminder     // 提醒
}

type task struct {
	service *TaskService
}
type taskCollaborator struct {
	service *TaskService
}
type taskComment struct {
	service *TaskService
}
type taskFollower struct {
	service *TaskService
}
type taskReminder struct {
	service *TaskService
}

// 批量删除执行者
//
// - 该接口用于批量删除执行者
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/batch_delete_collaborator
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/batchDeleteCollaborator_task.go
func (t *task) BatchDeleteCollaborator(ctx context.Context, req *BatchDeleteCollaboratorTaskReq, options ...larkcore.RequestOptionFunc) (*BatchDeleteCollaboratorTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/batch_delete_collaborator"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchDeleteCollaboratorTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 批量删除关注人
//
// - 该接口用于批量删除关注人
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/batch_delete_follower
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/batchDeleteFollower_task.go
func (t *task) BatchDeleteFollower(ctx context.Context, req *BatchDeleteFollowerTaskReq, options ...larkcore.RequestOptionFunc) (*BatchDeleteFollowerTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/batch_delete_follower"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchDeleteFollowerTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 完成任务
//
// - 该接口用于将任务状态修改为 “已完成”。;完成任务是指整个任务全部完成，而不支持执行者分别完成任务，执行成功后，任务对所有关联用户都变为完成状态。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/complete
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/complete_task.go
func (t *task) Complete(ctx context.Context, req *CompleteTaskReq, options ...larkcore.RequestOptionFunc) (*CompleteTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/complete"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CompleteTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建任务
//
// - 该接口可以创建一个任务，支持填写任务的基本信息，包括任务的标题，描述及协作者等。;在此基础上，创建任务时可以设置截止时间和重复规则，将任务设置为定期执行的重复任务。通过添加协作者，则可以让其他用户协同完成该任务。;此外，接口也提供了一些支持自定义内容的字段，调用方可以实现定制化效果，如完成任务后跳转到指定结束界面。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/create
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/create_task.go
func (t *task) Create(ctx context.Context, req *CreateTaskReq, options ...larkcore.RequestOptionFunc) (*CreateTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除任务
//
// - 该接口用于删除任务
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/delete
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/delete_task.go
func (t *task) Delete(ctx context.Context, req *DeleteTaskReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取任务详情
//
// - 该接口用于获取任务详情，包括任务标题、描述、时间、来源等信息
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/get
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/get_task.go
func (t *task) Get(ctx context.Context, req *GetTaskReq, options ...larkcore.RequestOptionFunc) (*GetTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取任务列表
//
// - 以分页的方式获取任务列表。当使用 user_access_token 时，获取与该用户身份相关的所有任务。当使用 tenant_access_token 时，获取以该应用身份通过 “创建任务 “接口创建的所有任务（并非获取该应用所在租户下所有用户创建的任务）。;本接口支持通过任务创建时间以及任务的完成状态对任务进行过滤。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/list
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/list_task.go
func (t *task) List(ctx context.Context, req *ListTaskReq, options ...larkcore.RequestOptionFunc) (*ListTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) ListByIterator(ctx context.Context, req *ListTaskReq, options ...larkcore.RequestOptionFunc) (*ListTaskIterator, error) {
	return &ListTaskIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 更新任务
//
// - 该接口用于修改任务的标题、描述、时间、来源等相关信息
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/patch
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/patch_task.go
func (t *task) Patch(ctx context.Context, req *PatchTaskReq, options ...larkcore.RequestOptionFunc) (*PatchTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 取消完成任务
//
// - 该接口用于取消任务的已完成状态
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/uncomplete
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/uncomplete_task.go
func (t *task) Uncomplete(ctx context.Context, req *UncompleteTaskReq, options ...larkcore.RequestOptionFunc) (*UncompleteTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/uncomplete"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UncompleteTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 新增执行者
//
// - 该接口用于新增任务执行者，一次性可以添加多个执行者。;只有任务的创建者和执行者才能添加执行者，关注人无权限添加。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/create
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/create_taskCollaborator.go
func (t *taskCollaborator) Create(ctx context.Context, req *CreateTaskCollaboratorReq, options ...larkcore.RequestOptionFunc) (*CreateTaskCollaboratorResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/collaborators"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskCollaboratorResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除执行者
//
// - 该接口用于删除任务执行者
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/delete
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/delete_taskCollaborator.go
func (t *taskCollaborator) Delete(ctx context.Context, req *DeleteTaskCollaboratorReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskCollaboratorResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/collaborators/:collaborator_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskCollaboratorResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取一个任务的执行者列表
//
// - 该接口用于查询任务执行者列表，支持分页，最大值为 50
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/list
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/list_taskCollaborator.go
func (t *taskCollaborator) List(ctx context.Context, req *ListTaskCollaboratorReq, options ...larkcore.RequestOptionFunc) (*ListTaskCollaboratorResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/collaborators"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskCollaboratorResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskCollaborator) ListByIterator(ctx context.Context, req *ListTaskCollaboratorReq, options ...larkcore.RequestOptionFunc) (*ListTaskCollaboratorIterator, error) {
	return &ListTaskCollaboratorIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 创建评论
//
// - 该接口用于创建和回复任务的评论。当 parent_id 字段为 0 时，为创建评论；当 parent_id 不为 0 时，为回复某条评论
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/create
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/create_taskComment.go
func (t *taskComment) Create(ctx context.Context, req *CreateTaskCommentReq, options ...larkcore.RequestOptionFunc) (*CreateTaskCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除评论
//
// - 该接口用于通过评论 ID 删除评论
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/delete
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/delete_taskComment.go
func (t *taskComment) Delete(ctx context.Context, req *DeleteTaskCommentReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments/:comment_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取评论详情
//
// - 该接口用于通过评论 ID 获取评论详情
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/get
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/get_taskComment.go
func (t *taskComment) Get(ctx context.Context, req *GetTaskCommentReq, options ...larkcore.RequestOptionFunc) (*GetTaskCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments/:comment_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTaskCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取评论列表
//
// - 该接口用于查询任务评论列表，支持分页，最大值为 100
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/list
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/list_taskComment.go
func (t *taskComment) List(ctx context.Context, req *ListTaskCommentReq, options ...larkcore.RequestOptionFunc) (*ListTaskCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskComment) ListByIterator(ctx context.Context, req *ListTaskCommentReq, options ...larkcore.RequestOptionFunc) (*ListTaskCommentIterator, error) {
	return &ListTaskCommentIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 更新评论
//
// - 该接口用于更新评论内容
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/update
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/update_taskComment.go
func (t *taskComment) Update(ctx context.Context, req *UpdateTaskCommentReq, options ...larkcore.RequestOptionFunc) (*UpdateTaskCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments/:comment_id"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateTaskCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 新增关注人
//
// - 该接口用于创建任务关注人。可以一次性添加多位关注人。关注人 ID 要使用表示用户的 ID。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/create
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/create_taskFollower.go
func (t *taskFollower) Create(ctx context.Context, req *CreateTaskFollowerReq, options ...larkcore.RequestOptionFunc) (*CreateTaskFollowerResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/followers"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskFollowerResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除关注人
//
// - 该接口用于删除任务关注人
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/delete
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/delete_taskFollower.go
func (t *taskFollower) Delete(ctx context.Context, req *DeleteTaskFollowerReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskFollowerResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/followers/:follower_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskFollowerResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取任务关注人列表
//
// - 该接口用于查询任务关注人列表，支持分页，最大值为 50
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/list
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/list_taskFollower.go
func (t *taskFollower) List(ctx context.Context, req *ListTaskFollowerReq, options ...larkcore.RequestOptionFunc) (*ListTaskFollowerResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/followers"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskFollowerResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskFollower) ListByIterator(ctx context.Context, req *ListTaskFollowerReq, options ...larkcore.RequestOptionFunc) (*ListTaskFollowerIterator, error) {
	return &ListTaskFollowerIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 新增提醒时间
//
// - 该接口用于创建任务的提醒时间。提醒时间在截止时间基础上做偏移，但是偏移后的结果不能早于当前时间。
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/create
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/create_taskReminder.go
func (t *taskReminder) Create(ctx context.Context, req *CreateTaskReminderReq, options ...larkcore.RequestOptionFunc) (*CreateTaskReminderResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/reminders"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskReminderResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除提醒时间
//
// - 删除提醒时间，返回结果状态
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/delete
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/delete_taskReminder.go
func (t *taskReminder) Delete(ctx context.Context, req *DeleteTaskReminderReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskReminderResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/reminders/:reminder_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskReminderResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询提醒时间列表
//
// - 返回提醒时间列表，支持分页，最大值为 50
//
// - 官网 API 文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/list
//
// - 使用 Demo 链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/taskv1/list_taskReminder.go
func (t *taskReminder) List(ctx context.Context, req *ListTaskReminderReq, options ...larkcore.RequestOptionFunc) (*ListTaskReminderResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/reminders"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskReminderResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskReminder) ListByIterator(ctx context.Context, req *ListTaskReminderReq, options ...larkcore.RequestOptionFunc) (*ListTaskReminderIterator, error) {
	return &ListTaskReminderIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
