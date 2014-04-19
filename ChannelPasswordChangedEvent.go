// ChannelPasswordChangedEvent
package ts3api

import (
	"errors"
	"strconv"
)

var _ Event = (*ChannelPasswordChangedEvent)(nil)

type ChannelPasswordChangedEvent struct {
	chId int
	*ApiHolder
}

func NewChannelPasswordChangedEvent() (event *ChannelPasswordChangedEvent) {
	event = &ChannelPasswordChangedEvent{}
	event.ApiHolder = &ApiHolder{}
	return
}

func (event *ChannelPasswordChangedEvent) setParam(key, val string) (err error) {
	if key == "cid" {
		event.chId, err = strconv.Atoi(val)
	} else {
		logger.Error("%s=%s is not valid!", key, val)
		err = errors.New(key + "=" + val + " not valid!")
	}
	return
}

func (event *ChannelPasswordChangedEvent) Id() int {
	return event.chId
}
