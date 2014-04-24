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
