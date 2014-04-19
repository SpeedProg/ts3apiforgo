// ClientMovedEvent
package ts3api

import (
	"errors"
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
		if err != nil {
			logger.Error("%s=%s is not valid!", key, val)
			err = errors.New(key + "=" + val + " is not valid!")
		}
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
