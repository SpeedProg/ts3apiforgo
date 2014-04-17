// ClientLeaveEvent
package ts3api

import ()

type ClientLeaveEvent struct {
	chFromId  int
	chToId    int
	reasonId  int
	reasonMsg string
	clId      int
	api       *TS3Api
}

func (event ClientLeaveEvent) ChannelFromId() int {
	return event.chFromId
}
func (event ClientLeaveEvent) ChannelToId() int {
	return event.chToId
}
func (event ClientLeaveEvent) ReasonId() int {
	return event.reasonId
}
func (event ClientLeaveEvent) ClientId() int {
	return event.clId
}
func (event ClientLeaveEvent) ReasonMsg() string {
	return event.reasonMsg
}
func (event ClientLeaveEvent) Api() *TS3Api {
	return event.api
}
