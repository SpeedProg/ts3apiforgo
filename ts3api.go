/*
This file is part of TS3QueryApi For Go.
TS3QueryApi For Go is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

// A layer for the TS3Query interface.
// There is a subpackage speedprog.de/ts3api/ts3const
// that contains constants to use for setabel properties.
// TODO: BINDINGLIST, find out how it looks with more then one binding
package ts3api

import (
	"code.google.com/p/log4go"
	"container/list"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"
)

var logger log4go.Logger

type TS3Api struct {
	conn         *ts3Connection
	lineList     *list.List
	listenerList *list.List
	commandMutex *sync.Mutex
}

type QueryError struct {
	Id  int
	Msg string
}

func init() {
	logger = log4go.NewLogger()
	logger.LoadConfiguration("ts3api_log4go.xml")
}

func (api TS3Api) reader(ch chan<- bool) {
	for {
		logger.Trace("Waiting for Message Queue...")
		msg, err := api.conn.ReadString('\n')
		if err != nil {
			logger.Error(err.Error())
			break
		}
		msg = strings.TrimSpace(msg)
		logger.Trace("Processing Message: %s", msg)
		msgp := strings.SplitN(msg, " ", 2)
		if strings.HasPrefix(msgp[0], "notify") {
			switch msgp[0] {
			case "notifycliententerview":
				api.dispatchClientJoinMessage(msgp[1])
			case "notifyclientleftview":
				api.dispatchClientLeftMessage(msgp[1])
			case "notifyserveredited":
				api.dispatchServerEditedMessage(msgp[1])
			case "notifychanneldescriptionchanged":
				api.dispatchChDescChangedMessage(msgp[1])
			case "notifychannelpasswordchanged":
				api.dispatchChannelPasswordChangedMessage(msgp[1])
			case "notifychannelmoved":
				api.dispatchChannelMovedMessage(msgp[1])
			case "notifychanneledited":
				api.dispatchChannelEditedMessage(msgp[1])
			case "notifychannelcreated":
				api.dispatchChannelCreatedMessage(msgp[1])
			case "notifychanneldeleted":
				api.dispatchChannelDeleted(msgp[1])
			case "notifyclientmoved":
				api.dispatchClientMovedMessage(msgp[1])
			case "notifytextmessage":
				api.dispatchTextMessage(msgp[1])
			case "notifytokenused":
				api.dispatchTokenUsedMesage(msgp[1])
			default:

			}
		} else {
			api.lineList.PushBack(msg)
			logger.Trace("Added To Message Queue: %s", msg)
		}

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
	event := NewClientJoinedEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchClientLeftMessage(msg string) {
	event := NewClientLeftEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchTextMessage(msg string) {
	event := NewTextMessageEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchTokenUsedMesage(msg string) {
	event := NewTokenUsedEvent()
	api.initEventFromString(event, msg)
	api.callListeners(event)
}

func (api TS3Api) dispatchChannelDeleted(msg string) {
	event := NewChannelDeletedEvent()
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
	logger.Trace("Taken From Message Queue: %s", msg)
	return
}

func parseQueryError(s string) (err QueryError) {
	var props []string = strings.Split(s, " ")
	for _, prop := range props {
		if strings.Contains(prop, "=") {
			var kvPair []string = strings.Split(prop, "=")
			switch kvPair[0] {
			case "id":
				err.Id, _ = strconv.Atoi(kvPair[1])
			case "msg":
				err.Msg = decodeValue(kvPair[1])
			}
		}
	}
	return
}

func (api TS3Api) initEventFromString(event Event, msg string) {
	params := strings.Split(msg, " ")
	for _, message := range params {
		if strings.Contains(message, "=") {
			keyval := strings.SplitN(message, "=", 2)
			event.setParam(keyval[0], decodeValue(keyval[1]))
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
	case *ClientJoinedEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.ClientJoined(event.(*ClientJoinedEvent))
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
	case *TextMessageEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.TextMessage(event.(*TextMessageEvent))
		}
	case *TokenUsedEvent:
		for element := api.listenerList.Front(); element != nil; element = element.Next() {
			listener := element.Value.(TS3Listener)
			go listener.TokenUsed(event.(*TokenUsedEvent))
		}
	default:
		logger.Error("Event of Type: %+v could not be handled.", t)
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

func encodeValue(s string) string {
	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "/", "\\/", -1)
	s = strings.Replace(s, " ", "\\s", -1)
	s = strings.Replace(s, "|", "\\p", -1)
	s = strings.Replace(s, "\a", "\\a", -1)
	s = strings.Replace(s, "\b", "\\b", -1)
	s = strings.Replace(s, "\f", "\\f", -1)
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\r", "\\r", -1)
	s = strings.Replace(s, "\t", "\\t", -1)
	s = strings.Replace(s, "\v", "\\v", -1)
	return s
}

func decodeValue(s string) string {
	s = strings.Replace(s, "\\\\", "\\[$]", -1)
	s = strings.Replace(s, "\\/", "/", -1)
	s = strings.Replace(s, "\\s", " ", -1)
	s = strings.Replace(s, "\\p", "|", -1)
	s = strings.Replace(s, "\\a", "\a", -1)
	s = strings.Replace(s, "\\b", "\b", -1)
	s = strings.Replace(s, "\\f", "\f", -1)
	s = strings.Replace(s, "\\n", "\n", -1)
	s = strings.Replace(s, "\\r", "\r", -1)
	s = strings.Replace(s, "\\t", "\t", -1)
	s = strings.Replace(s, "\\v", "\v", -1)
	s = strings.Replace(s, "\\[$]", "\\", -1)
	return s
}

// Creats a command string with a leading space
func cmdStringFromProperties(props [][]string) (cmd string) {
	for _, pel := range props {
		cmd += " " + strings.ToLower(pel[0]) + "=" + encodeValue(pel[1])
	}
	return
}
