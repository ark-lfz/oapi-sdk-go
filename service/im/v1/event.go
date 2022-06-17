// Package im code generated by oapi sdk gen
package im

import (
	"context"
)

/**
消息处理器定义
**/
type ChatDisbandedEventHandler struct {
	handler func(context.Context, *ChatDisbandedEvent) error
}

func NewChatDisbandedEventHandler(handler func(context.Context, *ChatDisbandedEvent) error) *ChatDisbandedEventHandler {
	h := &ChatDisbandedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ChatDisbandedEventHandler) Event() interface{} {
	return &ChatDisbandedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ChatDisbandedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ChatDisbandedEvent))
}

/**
消息处理器定义
**/
type ChatUpdatedEventHandler struct {
	handler func(context.Context, *ChatUpdatedEvent) error
}

func NewChatUpdatedEventHandler(handler func(context.Context, *ChatUpdatedEvent) error) *ChatUpdatedEventHandler {
	h := &ChatUpdatedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ChatUpdatedEventHandler) Event() interface{} {
	return &ChatUpdatedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ChatUpdatedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ChatUpdatedEvent))
}

/**
消息处理器定义
**/
type ChatMemberBotAddedEventHandler struct {
	handler func(context.Context, *ChatMemberBotAddedEvent) error
}

func NewChatMemberBotAddedEventHandler(handler func(context.Context, *ChatMemberBotAddedEvent) error) *ChatMemberBotAddedEventHandler {
	h := &ChatMemberBotAddedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ChatMemberBotAddedEventHandler) Event() interface{} {
	return &ChatMemberBotAddedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ChatMemberBotAddedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ChatMemberBotAddedEvent))
}

/**
消息处理器定义
**/
type ChatMemberBotDeletedEventHandler struct {
	handler func(context.Context, *ChatMemberBotDeletedEvent) error
}

func NewChatMemberBotDeletedEventHandler(handler func(context.Context, *ChatMemberBotDeletedEvent) error) *ChatMemberBotDeletedEventHandler {
	h := &ChatMemberBotDeletedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ChatMemberBotDeletedEventHandler) Event() interface{} {
	return &ChatMemberBotDeletedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ChatMemberBotDeletedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ChatMemberBotDeletedEvent))
}

/**
消息处理器定义
**/
type ChatMemberUserAddedEventHandler struct {
	handler func(context.Context, *ChatMemberUserAddedEvent) error
}

func NewChatMemberUserAddedEventHandler(handler func(context.Context, *ChatMemberUserAddedEvent) error) *ChatMemberUserAddedEventHandler {
	h := &ChatMemberUserAddedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ChatMemberUserAddedEventHandler) Event() interface{} {
	return &ChatMemberUserAddedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ChatMemberUserAddedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ChatMemberUserAddedEvent))
}

/**
消息处理器定义
**/
type ChatMemberUserDeletedEventHandler struct {
	handler func(context.Context, *ChatMemberUserDeletedEvent) error
}

func NewChatMemberUserDeletedEventHandler(handler func(context.Context, *ChatMemberUserDeletedEvent) error) *ChatMemberUserDeletedEventHandler {
	h := &ChatMemberUserDeletedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ChatMemberUserDeletedEventHandler) Event() interface{} {
	return &ChatMemberUserDeletedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ChatMemberUserDeletedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ChatMemberUserDeletedEvent))
}

/**
消息处理器定义
**/
type ChatMemberUserWithdrawnEventHandler struct {
	handler func(context.Context, *ChatMemberUserWithdrawnEvent) error
}

func NewChatMemberUserWithdrawnEventHandler(handler func(context.Context, *ChatMemberUserWithdrawnEvent) error) *ChatMemberUserWithdrawnEventHandler {
	h := &ChatMemberUserWithdrawnEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ChatMemberUserWithdrawnEventHandler) Event() interface{} {
	return &ChatMemberUserWithdrawnEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ChatMemberUserWithdrawnEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ChatMemberUserWithdrawnEvent))
}

/**
消息处理器定义
**/
type MessageMessageReadEventHandler struct {
	handler func(context.Context, *MessageMessageReadEvent) error
}

func NewMessageMessageReadEventHandler(handler func(context.Context, *MessageMessageReadEvent) error) *MessageMessageReadEventHandler {
	h := &MessageMessageReadEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *MessageMessageReadEventHandler) Event() interface{} {
	return &MessageMessageReadEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *MessageMessageReadEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MessageMessageReadEvent))
}

/**
消息处理器定义
**/
type MessageReceiveEventHandler struct {
	handler func(context.Context, *MessageReceiveEvent) error
}

func NewMessageReceiveEventHandler(handler func(context.Context, *MessageReceiveEvent) error) *MessageReceiveEventHandler {
	h := &MessageReceiveEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *MessageReceiveEventHandler) Event() interface{} {
	return &MessageReceiveEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *MessageReceiveEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MessageReceiveEvent))
}

/**
消息处理器定义
**/
type MessageReactionCreatedEventHandler struct {
	handler func(context.Context, *MessageReactionCreatedEvent) error
}

func NewMessageReactionCreatedEventHandler(handler func(context.Context, *MessageReactionCreatedEvent) error) *MessageReactionCreatedEventHandler {
	h := &MessageReactionCreatedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *MessageReactionCreatedEventHandler) Event() interface{} {
	return &MessageReactionCreatedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *MessageReactionCreatedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MessageReactionCreatedEvent))
}

/**
消息处理器定义
**/
type MessageReactionDeletedEventHandler struct {
	handler func(context.Context, *MessageReactionDeletedEvent) error
}

func NewMessageReactionDeletedEventHandler(handler func(context.Context, *MessageReactionDeletedEvent) error) *MessageReactionDeletedEventHandler {
	h := &MessageReactionDeletedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *MessageReactionDeletedEventHandler) Event() interface{} {
	return &MessageReactionDeletedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *MessageReactionDeletedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*MessageReactionDeletedEvent))
}
