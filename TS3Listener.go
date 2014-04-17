// TS3Listener
package ts3api

import ()

type TS3Listener interface {
	TextMessage(msg *TextMessageEvent)
	ClientJoined(event *ClientJoinEvent)
	ClientMoved(event *ClientMovedEvent)
	ClientLeft(event *ClientLeaveEvent)
	ServerEdited(event *ServerEditedEvent)
}
