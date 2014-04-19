// ClientLeaveEvent
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
