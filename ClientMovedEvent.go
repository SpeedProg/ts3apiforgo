// ClientMovedEvent
package ts3api

import (
	"strconv"
	"strings"
)

//notifyclientmoved ctid=1 reasonid=4 invokerid=9 invokername=SpeedProg invokeruid=8qoQBRqUwyGvQfGc08OUFAarE6Q= reasonmsg=test clid=20
type ClientMovedEvent struct {
	chToId      int
	reasonId    int
	reasonMsg   string
	clIds       []int
	invokerName string
	invokerId   int
	invokerUID  string
	api         *TS3Api
}

var _ Event = (*ClientMovedEvent)(nil)

func (event *ClientMovedEvent) setParam(key string, val string) (err error) {
	switch key {
	case "ctid":
		event.chToId, err = strconv.Atoi(val)
	case "reasonid":
		event.reasonId, err = strconv.Atoi(val)
	case "invokerid":
		event.invokerId, err = strconv.Atoi(val)
	case "invokername":
		event.invokerName = val
	case "invokeruid":
		event.invokerUID = val
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
		logger.Fatalln(key + "=" + val + " is not valid!")
	}
	return
}

func (event *ClientMovedEvent) InvokerUID() string {
	return event.invokerUID
}

func (event *ClientMovedEvent) InvokerId() int {
	return event.invokerId
}

func (event *ClientMovedEvent) InvokerName() string {
	return event.invokerName
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

func (event *ClientMovedEvent) Api() *TS3Api {
	return event.api
}
func (event *ClientMovedEvent) setApi(api *TS3Api) {
	event.api = api
}
