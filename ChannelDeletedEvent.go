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
