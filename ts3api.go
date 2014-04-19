// ts3api
package ts3api

import (
	"code.google.com/p/log4go"
	"container/list"
	"errors"
	"strings"
	"time"
)

var logger log4go.Logger

type TS3Api struct {
	conn         *ts3Connection
	lineList     *list.List
	listenerList *list.List
}

func init() {
	logger = log4go.NewLogger()
	logger.LoadConfiguration("log4go.xml")
}

func (api TS3Api) reader(ch chan<- bool) {
	for {
		logger.Error("Waiting for line...")
		msg, err := api.conn.ReadString('\n')
		if err != nil {
			logger.Error(err.Error())
			break
		}
		msg = strings.TrimSpace(msg)
		logger.Trace("Processing Message: %s", msg)
		prefixSplits := strings.SplitN(msg, " ", 2)
		if len(prefixSplits) < 2 {
			api.lineList.PushBack(msg)
		} else {

			switch prefixSplits[0] {
			case "notifytextmessage":
				api.dispatchTextMessage(prefixSplits[1])
			case "notifycliententerview":
				api.dispatchClientJoinMessage(prefixSplits[1])
			case "notifyclientmoved":
				api.dispatchClientMovedMessage(prefixSplits[1])
			case "notifyserveredited":
				api.dispatchServerEditedMessage(prefixSplits[1])
			default:
				logger.Error("Add To lineList")
				api.lineList.PushBack(msg)
			}
		}
		/*
			Read: notifyclientleftview cfid=4 ctid=0 reasonid=3 reasonmsg=connection\slost clid=16
			Read:
			Read: notifyclientleftview cfid=1 ctid=0 reasonid=8 reasonmsg=deselected\svirtualserver clid=18
			Read: notifycliententerview cfid=0 ctid=1 reasonid=0 clid=19 client_unique_identifier=a8DECwLONmPE4kNW0W2C3xDiRIA= client_nickname=Manfred\sfrom\s192.168.0.1:57329 client_input_muted=0 client_output_muted=0 client_outputonly_muted=0 client_input_hardware=0 client_output_hardware=0 client_meta_data client_is_recording=0 client_database_id=4 client_channel_group_id=8 client_servergroups=6,11 client_away=0 client_away_message client_type=1 client_flag_avatar=84ce0342c818723a75c51af66a723d2d client_talk_power=99 client_talk_request=0 client_talk_request_msg client_description=BETTER\sTHAN\sYUOF client_is_talker=0 client_is_priority_speaker=0 client_unread_messages=0 client_nickname_phonetic client_needed_serverquery_view_power=75 client_icon_id=0 client_is_channel_commander=0 client_country client_channel_group_inherited_channel_id=1 client_badges

		*/

	}
	ch <- true
}

func (api TS3Api) dispatchChannelPasswordChangedMessage(msg string) {
	event := NewChannelPasswordChangedEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchChannelMovedMessage(msg string) {
	event := NewChannelMovedEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchChannelEditedMessage(msg string) {
	event := NewChannelEditedEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchChannelCreatedMessage(msg string) {
	event := NewChannelCreatedEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchServerEditedMessage(msg string) {
	event := NewServerEditedEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchClientMovedMessage(msg string) {
	event := NewClientMovedEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchChDescChangedMessage(msg string) {
	event := NewChannelDescripionChangedEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchClientJoinMessage(msg string) {
	event := NewClientJoinEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchClientLeftMessage(msg string) {
	event := NewClientLeftEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchTextMessage(msg string) {
	event := &TextMessageEvent{}
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) readLine() (msg string) {
	for api.lineList.Len() < 1 {
		time.Sleep(100 * time.Millisecond)
	}
	element := api.lineList.Front()
	api.lineList.Remove(element)
	msg = element.Value.(string)
	logger.Error("-->" + msg + "<--")
	return
}

func (api TS3Api) doCommand(cmd string) (answer string) {
	api.conn.DoCommand(cmd)
	answer = api.readLine()
	return
}

func (api TS3Api) initEventFromString(event Event, msg string) {
	params := strings.Split(msg, " ")
	for _, message := range params {
		if strings.Contains(message, "=") {
			keyval := strings.SplitN(message, "=", 2)
			event.setParam(keyval[0], keyval[1])
		} else {
			event.setParam(message, "")
		}
	}
	event.setApi(&api)

}

func (api TS3Api) callListeners(event Event) {
	switch t := event.(type) {
	case *ChannelCreatedEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ChannelCreated(event.(*ChannelCreatedEvent))
		}
	case *ChannelDescriptionChangedEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ChannelDescriptionChanged(event.(*ChannelDescriptionChangedEvent))
		}
	case *ChannelEditedEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ChannelEdited(event.(*ChannelEditedEvent))
		}
	case *ChannelMovedEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ChannelMoved(event.(*ChannelMovedEvent))
		}
	case *ChannelPasswordChangedEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ChannelPasswordChanged(event.(*ChannelPasswordChangedEvent))
		}
	case *ClientJoinEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ClientJoined(event.(*ClientJoinEvent))
		}
	case *ClientLeftEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ClientLeft(event.(*ClientLeftEvent))
		}
	case *ClientMovedEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ClientMoved(event.(*ClientMovedEvent))
		}
	case *ServerEditedEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ServerEdited(event.(*ServerEditedEvent))
		}
	/*case *TextMessageEvent:
	for element := api.listenerList.Front(); element != nil; element = element.Next() {
		listener := element.Value.(TS3Listener)
		go listener.TextMessage(event.(*TextMessageEvent))
	}*/
	default:
		logger.Error("Event of Type: %s could not be handled.", t)
	}

}

/*
Maps 0 to false and 1 to true
Everything else trurns false and sets error
*/
func getBoolFromString(s string) (bool, error) {
	if s == "0" {
		return false, nil
	}
	if s == "1" {
		return true, nil
	}
	return false, errors.New(s + " is not valid!")
}
