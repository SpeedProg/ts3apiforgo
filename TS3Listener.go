// TS3Listener
package ts3api

import ()

type TS3Listener interface {
	ClientJoined(event *ClientJoinedEvent)
	ClientLeft(event *ClientLeftEvent)
	ServerEdited(event *ServerEditedEvent)
	ChannelDescriptionChanged(event *ChannelDescriptionChangedEvent)
	ChannelPasswordChanged(event *ChannelPasswordChangedEvent)
	ChannelMoved(event *ChannelMovedEvent)
	ChannelEdited(event *ChannelEditedEvent)
	ChannelCreated(event *ChannelCreatedEvent)
	ChannelDeletedEvent(event *ChannelDeletedEvent)
	ClientMoved(event *ClientMovedEvent)
	TextMessage(event *TextMessageEvent)
	TokenUsed(event *TokenUsedEvent)
}
