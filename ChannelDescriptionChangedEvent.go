// ChannelDescriptionChangedEvent
package ts3api

import (
	"errors"
	"strconv"
)

var _ Event = (*ChannelDescriptionChangedEvent)(nil)

type ChannelDescriptionChangedEvent struct {
	chId int
	*ApiHolder
}

func NewChannelDescripionChangedEvent() (event *ChannelDescriptionChangedEvent) {
	event = &ChannelDescriptionChangedEvent{}
	event.ApiHolder = &ApiHolder{}
	return
}

func (event *ChannelDescriptionChangedEvent) setParam(key, val string) (err error) {
	if key == "cid" {
		event.chId, err = strconv.Atoi(val)
	} else {
		logger.Error(key + "=" + val + " not valid!")
		err = errors.New(key + "=" + val + " not valid!")
	}
	return
}

func (event *ChannelDescriptionChangedEvent) ChannelId() int {
	return event.chId
}
