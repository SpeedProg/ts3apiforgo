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
