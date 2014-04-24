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

var _ Event = (*ChannelMovedEvent)(nil)

type ChannelMovedEvent struct {
	*ApiHolder
	*InvokerHolder
	chId       int
	chParentId int
	order      int
	reasonId   int
}

func NewChannelMovedEvent() (event *ChannelMovedEvent) {
	event = &ChannelMovedEvent{}
	event.ApiHolder = &ApiHolder{}
	event.InvokerHolder = &InvokerHolder{}
	return
}

func (event *ChannelMovedEvent) setParam(key, val string) (err error) {
	switch key {
	case "cid":
		event.chId, err = strconv.Atoi(val)
	case "cpid":
		event.chParentId, err = strconv.Atoi(val)
	case "order":
		event.order, err = strconv.Atoi(val)
	case "reasonid":
		event.reasonId, err = strconv.Atoi(val)
	default:
		err = event.InvokerHolder.setParam(key, val)
	}
	return
}

func (event *ChannelMovedEvent) ChannelId() int {
	return event.chId
}
func (event *ChannelMovedEvent) ChannelParentId() int {
	return event.chParentId
}
func (event *ChannelMovedEvent) Order() int {
	return event.order
}
func (event *ChannelMovedEvent) ReasonId() int {
	return event.reasonId
}
