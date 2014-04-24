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

// TextMessageEvent
package ts3api

import (
	"strconv"
)

var _ Event = (*TextMessageEvent)(nil)

type TextMessageEvent struct {
	*ApiHolder
	*InvokerHolder
	targetmode int
	target     int
	msg        string
}

func NewTextMessageEvent() (event *TextMessageEvent) {
	event = &TextMessageEvent{}
	event.ApiHolder = &ApiHolder{}
	event.InvokerHolder = &InvokerHolder{}
	return
}

func (event *TextMessageEvent) setParam(key, val string) (err error) {
	switch key {
	case "targetmode":
		event.targetmode, err = strconv.Atoi(val)
	case "msg":
		event.msg = val
	case "target":
		event.target, err = strconv.Atoi(val)
	default:
		err = event.InvokerHolder.setParam(key, val)
	}
	return
}

func (event *TextMessageEvent) Msg() string {
	return event.msg
}

func (event *TextMessageEvent) Targetmode() int {
	return event.targetmode
}

func (event *TextMessageEvent) Target() int {
	return event.target
}
