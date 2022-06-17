// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"
	"github.com/feishu/oapi-sdk-go/service/drive/v1"
)

func (dispatcher *EventReqDispatcher) FileDeletedV1(handler func(ctx context.Context, event *drive.FileDeletedEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["drive.file.deleted_v1"] = drive.NewFileDeletedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) FileEditV1(handler func(ctx context.Context, event *drive.FileEditEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["drive.file.edit_v1"] = drive.NewFileEditEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) FilePermissionMemberAddedV1(handler func(ctx context.Context, event *drive.FilePermissionMemberAddedEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["drive.file.permission_member_added_v1"] = drive.NewFilePermissionMemberAddedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) FilePermissionMemberRemovedV1(handler func(ctx context.Context, event *drive.FilePermissionMemberRemovedEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["drive.file.permission_member_removed_v1"] = drive.NewFilePermissionMemberRemovedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) FileReadV1(handler func(ctx context.Context, event *drive.FileReadEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["drive.file.read_v1"] = drive.NewFileReadEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) FileTitleUpdatedV1(handler func(ctx context.Context, event *drive.FileTitleUpdatedEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["drive.file.title_updated_v1"] = drive.NewFileTitleUpdatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventReqDispatcher) FileTrashedV1(handler func(ctx context.Context, event *drive.FileTrashedEvent) error) *EventReqDispatcher {
	dispatcher.eventType2EventHandler["drive.file.trashed_v1"] = drive.NewFileTrashedEventHandler(handler)
	return dispatcher
}
