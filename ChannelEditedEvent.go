// ChannelEditedEvent
package ts3api

import (
	"errors"
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
		if err != nil {
			logger.Error(key + "=" + val + " not valid!")
			err = errors.New(key + "=" + val + " not valid!")
		}
	}
	return
}

func (event *ChannelEditedEvent) ChannelId() int {
	return event.chId
}
func (event *ChannelEditedEvent) ReasonId() int {
	return event.reasonId
}
