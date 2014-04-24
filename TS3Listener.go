/*
This file is part of TS3QueryApi For Go.
TS3QueryApi For Go is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
