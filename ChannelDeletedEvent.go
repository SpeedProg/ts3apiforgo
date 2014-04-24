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

var _ Event = (*ChannelDeletedEvent)(nil)

type ChannelDeletedEvent struct {
	*InvokerHolder
	*ApiHolder
	chId int
}

func NewChannelDeletedEvent() (event *ChannelDeletedEvent) {
	event = &ChannelDeletedEvent{}
	event.InvokerHolder = &InvokerHolder{}
	event.ApiHolder = &ApiHolder{}
	return
}

func (event *ChannelDeletedEvent) setParam(key, val string) (err error) {
	switch key {
	case "cid":
		event.chId, err = strconv.Atoi(val)
	default:
		err = event.InvokerHolder.setParam(key, val)
	}
	return
}

func (event *ChannelDeletedEvent) ChannelId() int {
	return event.chId
}
