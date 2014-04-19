// TS3Listener
package ts3api

import ()

type TS3Listener interface {
	TextMessage(msg *TextMessageEvent)
	ClientJoined(event *ClientJoinEvent)
	ClientMoved(event *ClientMovedEvent)
	ClientLeft(event *ClientLeftEvent)
	ServerEdited(event *ServerEditedEvent)
	ChannelEdited(event *ChannelEditedEvent)
	ChannelCreated(event *ChannelCreatedEvent)
	ChannelDescriptionChanged(event *ChannelDescriptionChangedEvent)
	ChannelPasswordChanged(event *ChannelPasswordChangedEvent)
	ChannelMoved(event *ChannelMovedEvent)
}
