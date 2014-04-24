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

var _ Event = (*ClientLeftEvent)(nil)

type ClientLeftEvent struct {
	*InvokerHolder
	*ApiHolder
	chFromId  int
	chToId    int
	reasonId  int
	reasonMsg string
	clId      int
	bantime   int
}

func NewClientLeftEvent() (event *ClientLeftEvent) {
	event = &ClientLeftEvent{}
	event.InvokerHolder = &InvokerHolder{}
	event.ApiHolder = &ApiHolder{}
	return
}

func (event *ClientLeftEvent) setParam(key string, val string) (err error) {
	switch key {
	case "cfid":
		event.chFromId, err = strconv.Atoi(val)
	case "ctid":
		event.chToId, err = strconv.Atoi(val)
	case "reasonid":
		event.reasonId, err = strconv.Atoi(val)
	case "reasonmsg":
		event.reasonMsg = val
	case "clid":
		event.clId, err = strconv.Atoi(val)
	case "bantime":
		event.bantime, err = strconv.Atoi(val)
	default:
		err = event.InvokerHolder.setParam(key, val)
	}
	return
}
func (event *ClientLeftEvent) ChannelFromId() int {
	return event.chFromId
}

func (event *ClientLeftEvent) ChannelToId() int {
	return event.chToId
}

func (event *ClientLeftEvent) ReasonId() int {
	return event.reasonId
}

func (event *ClientLeftEvent) Id() int {
	return event.clId
}

func (event *ClientLeftEvent) ReasonMsg() string {
	return event.reasonMsg
}

func (event *ClientLeftEvent) Bantime() int {
	return event.bantime
}
