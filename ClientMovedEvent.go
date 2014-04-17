// ClientMovedEvent
package ts3api

type ClientMovedEvent struct {
	chTargetId int
	reasonId   int
	cId        int
	api        *TS3Api
}

func (event ClientMovedEvent) ChannelTargetId() int {
	return event.chTargetId
}

func (event ClientMovedEvent) ReasonId() int {
	return event.reasonId
}

func (event ClientMovedEvent) ClientId() int {
	return event.cId
}

func (event ClientMovedEvent) Api() *TS3Api {
	return event.api
}
