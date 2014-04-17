// ts3api
package ts3api

import (
	"container/list"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var logger *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

type TS3Api struct {
	conn         *ts3Connection
	lineList     *list.List
	listenerList *list.List
}

func (api TS3Api) reader(ch chan<- bool) {
	for {
		logger.Println("Waiting for line...")
		msg, err := api.conn.ReadString('\n')
		if err != nil {
			logger.Fatalln(err.Error())
			break
		}
		msg = strings.TrimSpace(msg)
		logger.Println(msg)
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
			default:
				logger.Println("Add To lineList")
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

func (api TS3Api) dispatchClientMovedMessage(msg string) {
	clientMovedEv := api.clientMovedEventFromString(msg)
	for element := api.listenerList.Front(); element != nil; element = element.Next() {
		listener := element.Value.(TS3Listener)
		go listener.ClientMoved(clientMovedEv)
	}
}

func (api TS3Api) dispatchClientJoinMessage(msg string) {
	clientJoinEv := api.clientJoinEventFromString(msg)
	for element := api.listenerList.Front(); element != nil; element = element.Next() {
		listener := element.Value.(TS3Listener)
		go listener.ClientJoined(clientJoinEv)
	}
}

func (api TS3Api) dispatchClientLeaveMessage(msg string) {
	clientLeaveEv := api.clientLeaveEventFromString(msg)
	for element := api.listenerList.Front(); element != nil; element = element.Next() {
		listener := element.Value.(TS3Listener)
		go listener.ClientLeft(clientLeaveEv)
	}
}

func (api TS3Api) dispatchTextMessage(msg string) {
	textMessageEvent := &TextMessageEvent{}
	pairs := strings.Split(msg, " ")
	for index, message := range pairs {
		keyval := strings.SplitN(message, "=", 2)
		index += index
		switch keyval[0] {
		case "targetmode":
			textMessageEvent.Targetmode, _ = strconv.Atoi(keyval[1])
		case "invokerid":
			textMessageEvent.InvokerId, _ = strconv.Atoi(keyval[1])
		case "msg":
			textMessageEvent.Msg = keyval[1]
		case "invokername":
			textMessageEvent.InvokerName = keyval[1]
		case "invokeruid":
			textMessageEvent.InvokerUid = keyval[1]
		}
	}
	textMessageEvent.Api = &api
	for element := api.listenerList.Front(); element != nil; element = element.Next() {
		listener := element.Value.(TS3Listener)
		go listener.TextMessage(textMessageEvent)
	}
}

func (api TS3Api) readLine() (msg string) {
	for api.lineList.Len() < 1 {
		time.Sleep(100 * time.Millisecond)
	}
	element := api.lineList.Front()
	api.lineList.Remove(element)
	msg = element.Value.(string)
	logger.Println("-->" + msg + "<--")
	return
}

func (api TS3Api) doCommand(cmd string) (answer string) {
	api.conn.DoCommand(cmd)
	answer = api.readLine()
	return
}

/*
	Expects the "notifycliententerview " to be allready removed!
*/
func (api TS3Api) clientJoinEventFromString(msg string) *ClientJoinEvent {
	event := ClientJoinEvent{}
	params := strings.Split(msg, " ")
	for _, message := range params {
		if strings.Contains(message, "=") {
			keyval := strings.SplitN(message, "=", 2)
			event.setParam(keyval[0], keyval[1])
		} else {
			event.setParam(message, "")
		}
	}
	event.api = &api
	return &event
}

/*
	Expects the "notifyclientleftview " to be allready removed!
*/
func (api TS3Api) clientLeaveEventFromString(msg string) *ClientLeaveEvent {
	event := ClientLeaveEvent{}
	params := strings.Split(msg, " ")
	for _, message := range params {
		if strings.Contains(message, "=") {
			keyval := strings.SplitN(message, "=", 2)
			event.setParam(keyval[0], keyval[1])
		} else {
			event.setParam(message, "")
		}
	}
	event.api = &api
	return &event
}

/*
	Expects the " " to be allready removed!
*/
func (api TS3Api) clientMovedEventFromString(msg string) *ClientMovedEvent {
	event := ClientMovedEvent{}
	params := strings.Split(msg, " ")
	for _, message := range params {
		if strings.Contains(message, "=") {
			keyval := strings.SplitN(message, "=", 2)
			event.setParam(keyval[0], keyval[1])
		} else {
			event.setParam(message, "")
		}
	}
	event.api = &api
	return &event
}
