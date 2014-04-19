package ts3api

// Implemts the

var _ TS3Listener = (*TS3Adapter)(nil)

// Implemts the TS3Listener interface
// All methods do nothing
// Can be embedded into own listeners to avoid needing to add all methods
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
