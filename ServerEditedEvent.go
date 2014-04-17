// ServerEditedEvent
package ts3api

import (
	"strconv"
)

type ServerEditedEvent struct {
	reasonId    int
	invokerId   int
	invokerName string
	invokerUID  string
	values      string
	api         *TS3Api
}

var _ Event = (*ServerEditedEvent)(nil)

func (event *ServerEditedEvent) setParam(key string, val string) (err error) {
	switch key {
	case "reasonid":
		event.reasonId, err = strconv.Atoi(val)
	case "invokerid":
		event.invokerId, err = strconv.Atoi(val)
	case "invokername":
		event.invokerName = val
	case "invokeruid":
		event.invokerUID = val
	default:
		if event.isValidValue(key) {
			event.values += key + "=" + val
		} else {
			logger.Fatalln(key + "=" + val + " is not valid!")
		}
	}
	return
}

func (event *ServerEditedEvent) isValidValue(key string) (isValid bool) {
	switch key {
	case "virtualserver_name":
		fallthrough
	case "virtualserver_codec_encryption_mode":
		fallthrough
	case "virtualserver_default_server_group":
		fallthrough
	case "virtualserver_default_channel_group":
		fallthrough
	case "virtualserver_hostbanner_url":
		fallthrough
	case "virtualserver_hostbanner_gfx_url":
		fallthrough
	case "virtualserver_hostbanner_gfx_interval":
		fallthrough
	case "virtualserver_priority_speaker_dimm_modificator":
		fallthrough
	case "virtualserver_hostbutton_tooltip":
		fallthrough
	case "virtualserver_hostbutton_url":
		fallthrough
	case "virtualserver_hostbutton_gfx_url":
		fallthrough
	case "virtualserver_name_phonetic":
		fallthrough
	case "virtualserver_icon_id":
		fallthrough
	case "virtualserver_hostbanner_mode":
		fallthrough
	case "virtualserver_channel_temp_delete_delay_default":
		isValid = true
	default:
		isValid = false

	}
	return
}

func (event *ServerEditedEvent) ReasonId() int {
	return event.reasonId
}

func (event *ServerEditedEvent) InvokerId() int {
	return event.invokerId
}

func (event *ServerEditedEvent) InvokerName() string {
	return event.invokerName
}

func (event *ServerEditedEvent) InvokerUID() string {
	return event.invokerUID
}

func (event *ServerEditedEvent) Values() string {
	return event.values
}

func (event *ServerEditedEvent) Api() *TS3Api {
	return event.api
}

func (event *ServerEditedEvent) setApi(api *TS3Api) {
	event.api = api
}
