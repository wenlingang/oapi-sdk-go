// Package dispatcher code generated by oapi sdk gen
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

package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

// 成员字段变更
//
// - 通过该事件订阅成员字段变更。old_object 展示更新字段的原始值。
//
// - 触发事件的动作有「打开/关闭」开关、「增加/删除」成员字段。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/custom_attr_event/events/updated
func (dispatcher *EventDispatcher) OnP2CustomAttrEventUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2CustomAttrEventUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.custom_attr_event.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.custom_attr_event.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.custom_attr_event.updated_v3"] = larkcontact.NewP2CustomAttrEventUpdatedV3Handler(handler)
	return dispatcher
}

// 部门被创建
//
// - 创建通讯录部门时发送该事件给订阅应用。
//
// - 只有当应用拥有被改动字段的数据权限时，才会接收到事件。具体的数据权限与字段的关系请参考 [应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看事件体参数列表的字段描述。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/events/created
func (dispatcher *EventDispatcher) OnP2DepartmentCreatedV3(handler func(ctx context.Context, event *larkcontact.P2DepartmentCreatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.department.created_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.department.created_v3")
	}
	dispatcher.eventType2EventHandler["contact.department.created_v3"] = larkcontact.NewP2DepartmentCreatedV3Handler(handler)
	return dispatcher
}

// 部门被删除
//
// - 订阅这一事件可以获得被删除部门的信息。
//
// - 只有当应用拥有被改动字段的数据权限时，才会接收到事件。具体的数据权限与字段的关系请参考 [应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看事件体参数列表的字段描述。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/events/deleted
func (dispatcher *EventDispatcher) OnP2DepartmentDeletedV3(handler func(ctx context.Context, event *larkcontact.P2DepartmentDeletedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.department.deleted_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.department.deleted_v3")
	}
	dispatcher.eventType2EventHandler["contact.department.deleted_v3"] = larkcontact.NewP2DepartmentDeletedV3Handler(handler)
	return dispatcher
}

// 部门信息被修改
//
// - 通过该事件订阅部门更新。`old_object`只展示被更新字段的原始值。应用身份访问通讯录的权限为历史版本，不推荐申请。
//
// - 只有当应用拥有被改动字段的数据权限时，才会接收到事件。具体的数据权限与字段的关系请参考 [应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看事件体参数列表的字段描述。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/events/updated
func (dispatcher *EventDispatcher) OnP2DepartmentUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2DepartmentUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.department.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.department.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.department.updated_v3"] = larkcontact.NewP2DepartmentUpdatedV3Handler(handler)
	return dispatcher
}

// 启用人员类型事件
//
// - 启用人员类型会发出对应事件。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/events/actived
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumActivedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumActivedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.actived_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.actived_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.actived_v3"] = larkcontact.NewP2EmployeeTypeEnumActivedV3Handler(handler)
	return dispatcher
}

// 新建人员类型事件
//
// - 新建人员类型会发出对应事件。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/events/created
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumCreatedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumCreatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.created_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.created_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.created_v3"] = larkcontact.NewP2EmployeeTypeEnumCreatedV3Handler(handler)
	return dispatcher
}

// 停用人员类型事件
//
// - 停用人员类型会发出对应事件。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/events/deactivated
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumDeactivatedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumDeactivatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.deactivated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.deactivated_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.deactivated_v3"] = larkcontact.NewP2EmployeeTypeEnumDeactivatedV3Handler(handler)
	return dispatcher
}

// 删除人员类型事件
//
// - 删除人员类型会发出对应事件。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/events/deleted
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumDeletedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumDeletedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.deleted_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.deleted_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.deleted_v3"] = larkcontact.NewP2EmployeeTypeEnumDeletedV3Handler(handler)
	return dispatcher
}

// 修改人员类型名称事件
//
// - 修改人员类型名称会发出对应事件。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/events/updated
func (dispatcher *EventDispatcher) OnP2EmployeeTypeEnumUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2EmployeeTypeEnumUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.updated_v3"] = larkcontact.NewP2EmployeeTypeEnumUpdatedV3Handler(handler)
	return dispatcher
}

// 通讯录范围权限被更新
//
// - 当应用通讯录范围权限发生变更时，订阅这个事件的应用会收到事件。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/scope/events/updated
func (dispatcher *EventDispatcher) OnP2ScopeUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2ScopeUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.scope.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.scope.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.scope.updated_v3"] = larkcontact.NewP2ScopeUpdatedV3Handler(handler)
	return dispatcher
}

// 员工入职
//
// - 通过该事件订阅员工入职。
//
// - 只有当应用拥有被改动字段的数据权限时，才会接收到事件。具体的数据权限与字段的关系请参考 [应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看事件体参数列表的字段描述。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/events/created
func (dispatcher *EventDispatcher) OnP2UserCreatedV3(handler func(ctx context.Context, event *larkcontact.P2UserCreatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.user.created_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.user.created_v3")
	}
	dispatcher.eventType2EventHandler["contact.user.created_v3"] = larkcontact.NewP2UserCreatedV3Handler(handler)
	return dispatcher
}

// 员工离职
//
// - 通过该事件订阅员工离职。应用身份访问通讯录的权限为历史版本，不推荐申请。
//
// - 只有当应用拥有被改动字段的数据权限时，才会接收到事件。具体的数据权限与字段的关系请参考 [应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看事件体参数列表的字段描述。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/events/deleted
func (dispatcher *EventDispatcher) OnP2UserDeletedV3(handler func(ctx context.Context, event *larkcontact.P2UserDeletedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.user.deleted_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.user.deleted_v3")
	}
	dispatcher.eventType2EventHandler["contact.user.deleted_v3"] = larkcontact.NewP2UserDeletedV3Handler(handler)
	return dispatcher
}

// 员工变更
//
// - 通过该事件订阅员工变更。old_object 中只展示更新的字段的原始值。
//
// - 只有当应用拥有被改动字段的数据权限时，才会接收到事件。具体的数据权限与字段的关系请参考 [应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看事件体参数列表的字段描述。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/events/updated
func (dispatcher *EventDispatcher) OnP2UserUpdatedV3(handler func(ctx context.Context, event *larkcontact.P2UserUpdatedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.user.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.user.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.user.updated_v3"] = larkcontact.NewP2UserUpdatedV3Handler(handler)
	return dispatcher
}
