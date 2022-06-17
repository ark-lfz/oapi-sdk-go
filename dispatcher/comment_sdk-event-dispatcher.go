// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"
	"github.com/feishu/oapi-sdk-go/service/comment_sdk/v1"
)

func (dispatcher *EventReqDispatcher) EntityCommentAddV1(handler func(ctx context.Context, event *comment_sdk.EntityCommentAddEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["comment_sdk.entity.comment.add_v1"] = comment_sdk.NewEntityCommentAddEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) EntityCommentFinishV1(handler func(ctx context.Context, event *comment_sdk.EntityCommentFinishEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["comment_sdk.entity.comment.finish_v1"] = comment_sdk.NewEntityCommentFinishEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) EntityCommentReopenV1(handler func(ctx context.Context, event *comment_sdk.EntityCommentReopenEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["comment_sdk.entity.comment.reopen_v1"] = comment_sdk.NewEntityCommentReopenEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) EntityCommentReplyAddV1(handler func(ctx context.Context, event *comment_sdk.EntityCommentReplyAddEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["comment_sdk.entity.comment.reply.add_v1"] = comment_sdk.NewEntityCommentReplyAddEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) EntityCommentReplyDeleteV1(handler func(ctx context.Context, event *comment_sdk.EntityCommentReplyDeleteEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["comment_sdk.entity.comment.reply.delete_v1"] = comment_sdk.NewEntityCommentReplyDeleteEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) EntityCommentReplyUpdateV1(handler func(ctx context.Context, event *comment_sdk.EntityCommentReplyUpdateEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["comment_sdk.entity.comment.reply.update_v1"] = comment_sdk.NewEntityCommentReplyUpdateEventHandler(handler)
	return dispatcher
}
