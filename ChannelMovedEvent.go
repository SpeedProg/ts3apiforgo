// ChannelMovedEvent
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
