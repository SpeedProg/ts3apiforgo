// ServerEditedEvent
package ts3api

import (
	"errors"
	"strconv"
)

var _ Event = (*ServerEditedEvent)(nil)

type ServerEditedEvent struct {
	*InvokerHolder
	*ApiHolder
	reasonId int
	values   string
}

func NewServerEditedEvent() (event *ServerEditedEvent) {
	event = &ServerEditedEvent{}
	event.InvokerHolder = &InvokerHolder{}
	event.ApiHolder = &ApiHolder{}
	return
}

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
			err = event.InvokerHolder.setParam(key, val)
			if err != nil {
				logger.Error(key + "=" + val + " not valid!")
				err = errors.New(key + "=" + val + " not valid!")
			}

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

func (event *ServerEditedEvent) Values() string {
	return event.values
}
