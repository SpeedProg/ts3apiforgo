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

import (
	"strconv"
)

var _ Event = (*ChannelEditedEvent)(nil)

type ChannelEditedEvent struct {
	*InvokerHolder
	*ApiHolder
	chId     int
	reasonId int
}

func NewChannelEditedEvent() (event *ChannelEditedEvent) {
	event = &ChannelEditedEvent{}
	event.InvokerHolder = &InvokerHolder{}
	event.ApiHolder = &ApiHolder{}
	return
}

func (event *ChannelEditedEvent) setParam(key, val string) (err error) {
	if key == "cid" {
		event.chId, err = strconv.Atoi(val)
	} else {
		err = event.InvokerHolder.setParam(key, val)
	}
	return
}

func (event *ChannelEditedEvent) ChannelId() int {
	return event.chId
}
func (event *ChannelEditedEvent) ReasonId() int {
	return event.reasonId
}
