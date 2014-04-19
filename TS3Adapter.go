package ts3api

type TS3Adapter struct {
}

func (obs *TS3Adapter) TextMessage(msg *TextMessageEvent)                               {}
func (obs *TS3Adapter) ClientJoined(event *ClientJoinedEvent)                           {}
func (obs *TS3Adapter) ClientMoved(event *ClientMovedEvent)                             {}
func (obs *TS3Adapter) ServerEdited(event *ServerEditedEvent)                           {}
func (obs *TS3Adapter) ClientLeft(event *ClientLeftEvent)                               {}
func (obs *TS3Adapter) ChannelCreated(event *ChannelCreatedEvent)                       {}
func (obs *TS3Adapter) ChannelDescriptionChanged(event *ChannelDescriptionChangedEvent) {}
func (obs *TS3Adapter) ChannelEdited(event *ChannelEditedEvent)                         {}
func (obs *TS3Adapter) ChannelMoved(event *ChannelMovedEvent)                           {}
func (obs *TS3Adapter) ChannelPasswordChanged(event *ChannelPasswordChangedEvent)       {}
func (obs *TS3Adapter) ChannelDeletedEvent(event *ChannelDeletedEvent)                  {}
func (obs *TS3Adapter) TokenUsed(event *TokenUsedEvent)                                 {}
