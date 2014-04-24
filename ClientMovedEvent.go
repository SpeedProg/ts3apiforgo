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
	"strings"
)

var _ Event = (*ClientMovedEvent)(nil)

//notifyclientmoved ctid=1 reasonid=4 invokerid=9 invokername=SpeedProg invokeruid=8qoQBRqUwyGvQfGc08OUFAarE6Q= reasonmsg=test clid=20
type ClientMovedEvent struct {
	*ApiHolder
	*InvokerHolder
	chToId    int
	reasonId  int
	reasonMsg string
	clIds     []int
}

func NewClientMovedEvent() (event *ClientMovedEvent) {
	event = &ClientMovedEvent{}
	event.InvokerHolder = &InvokerHolder{}
	event.ApiHolder = &ApiHolder{}
	return
}

func (event *ClientMovedEvent) setParam(key string, val string) (err error) {
	switch key {
	case "ctid":
		event.chToId, err = strconv.Atoi(val)
	case "reasonid":
		event.reasonId, err = strconv.Atoi(val)
	case "reasonmsg":
		event.reasonMsg = val
	case "clid":
		// this can be a list 1,2,3 ...
		if strings.Contains(val, ",") {
			ids := strings.Split(val, ",")
			event.clIds = make([]int, len(ids))
			for idx, id := range ids {
				event.clIds[idx], err = strconv.Atoi(id)
				// TODO: error check
			}
		} else { // and if not
			event.clIds = make([]int, 1)
			event.clIds[0], err = strconv.Atoi(val)
		}
	default:
		err = event.InvokerHolder.setParam(key, val)
	}
	return
}

func (event *ClientMovedEvent) ReasonMsg() string {
	return event.reasonMsg
}

func (event *ClientMovedEvent) ChannelToId() int {
	return event.chToId
}

func (event *ClientMovedEvent) ReasonId() int {
	return event.reasonId
}

func (event *ClientMovedEvent) Ids() []int {
	return event.clIds
}
