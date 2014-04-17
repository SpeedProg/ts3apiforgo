// ClientLeaveEvent
package ts3api

import (
	"strconv"
)

type ClientLeaveEvent struct {
	chFromId    int
	chToId      int
	reasonId    int
	reasonMsg   string
	clId        int
	invokerId   int
	invokerName string
	invokerUId  string
	bantime     int
	api         *TS3Api
}

func (event *ClientLeaveEvent) setParam(key string, val string) (err error) {
	switch key {
	case "cfid":
		event.chFromId, err = strconv.Atoi(val)
	case "ctid":
		event.chToId, err = strconv.Atoi(val)
	case "reasonid":
		event.reasonId, err = strconv.Atoi(val)
	case "reasonmsg":
		event.reasonMsg = val
	case "invokerid":
		event.invokerId, err = strconv.Atoi(val)
	case "invokername":
		event.invokerName = val
	case "invokeruid":
		event.invokerUId = val
	case "clid":
		event.clId, err = strconv.Atoi(val)
	case "bantime":
		event.bantime, err = strconv.Atoi(val)
	default:
		logger.Fatalln(key + "=" + val + " is not valid!")
	}
	return
}

func (event *ClientLeaveEvent) ChannelFromId() int {
	return event.chFromId
}

func (event *ClientLeaveEvent) ChannelToId() int {
	return event.chToId
}

func (event *ClientLeaveEvent) ReasonId() int {
	return event.reasonId
}

func (event *ClientLeaveEvent) Id() int {
	return event.clId
}

func (event *ClientLeaveEvent) ReasonMsg() string {
	return event.reasonMsg
}

func (event *ClientLeaveEvent) InvokerId() int {
	return event.invokerId
}

func (event *ClientLeaveEvent) InvokerName() string {
	return event.invokerName
}

func (event *ClientLeaveEvent) InvokerUId() string {
	return event.invokerUId
}

func (event *ClientLeaveEvent) Bantime() int {
	return event.bantime
}

func (event *ClientLeaveEvent) Api() *TS3Api {
	return event.api
}
