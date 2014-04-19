// TextMessageEvent
package ts3api

import (
	"errors"
	"strconv"
)

var _ Event = (*TextMessageEvent)(nil)

type TextMessageEvent struct {
	*ApiHolder
	*InvokerHolder
	targetmode int
	msg        string
}

func (event *TextMessageEvent) setParam(key, val string) (err error) {
	switch key {
	case "targetmode":
		event.targetmode, err = strconv.Atoi(val)
	case "msg":
		event.msg = val
	default:
		err = event.InvokerHolder.setParam(key, val)
		if err != nil {
			logger.Error(key + "=" + val + " not valid!")
			err = errors.New(key + "=" + val + " not valid!")
		}
	}
	return
}

func (event *TextMessageEvent) Msg() string {
	return event.msg
}

func (event *TextMessageEvent) Targetmode() int {
	return event.targetmode
}
