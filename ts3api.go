// ts3api
package ts3api

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var logger *log.Logger = log.New(bufio.NewWriter(os.Stdout), "TS3Api: ", log.Ldate|log.Ltime|log.Lshortfile)

type TS3Api struct {
	conn         *ts3Connection
	lineList     *list.List
	listenerList *list.List
}

func (api TS3Api) reader(ch chan<- bool) {
	for {

		msg, err := api.conn.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		msg = strings.TrimSpace(msg)
		prefixSplits := strings.SplitN(msg, " ", 2)
		if len(prefixSplits) < 2 {
			// TODO: handle weired stuff
		} else {
			switch prefixSplits[0] {
			case "notifytextmessage":
				api.dispatchTextMessage(prefixSplits[1])
			case "notifycliententerview":
				api.dispatchClientJoinMessage(prefixSplits[1])
			case "notifyclientmoved":
				api.dispatchClientMovedMessage(prefixSplits[1])
			}
		}
		/*
			Read: notifyclientleftview cfid=4 ctid=0 reasonid=3 reasonmsg=connection\slost clid=16
			Read:
			Read: notifyclientleftview cfid=1 ctid=0 reasonid=8 reasonmsg=deselected\svirtualserver clid=18
			Read: notifycliententerview cfid=0 ctid=1 reasonid=0 clid=19 client_unique_identifier=a8DECwLONmPE4kNW0W2C3xDiRIA= client_nickname=Manfred\sfrom\s192.168.0.1:57329 client_input_muted=0 client_output_muted=0 client_outputonly_muted=0 client_input_hardware=0 client_output_hardware=0 client_meta_data client_is_recording=0 client_database_id=4 client_channel_group_id=8 client_servergroups=6,11 client_away=0 client_away_message client_type=1 client_flag_avatar=84ce0342c818723a75c51af66a723d2d client_talk_power=99 client_talk_request=0 client_talk_request_msg client_description=BETTER\sTHAN\sYUOF client_is_talker=0 client_is_priority_speaker=0 client_unread_messages=0 client_nickname_phonetic client_needed_serverquery_view_power=75 client_icon_id=0 client_is_channel_commander=0 client_country client_channel_group_inherited_channel_id=1 client_badges

		*/
		api.lineList.PushBack(msg)

		logger.Println("Read: " + msg)
	}
	ch <- true
}

func (api TS3Api) dispatchClientMovedMessage(msg string) {
	clientMovedEv := &ClientMovedEvent{}
	params := strings.Split(msg, " ")
	for _, message := range params {
		if strings.Contains(message, "=") {
			keyval := strings.SplitN(message, "=", 2)
			switch keyval[0] {
			case "ctid":
				clientMovedEv.chTargetId, _ = strconv.Atoi(keyval[1])
			case "reasonid":
				clientMovedEv.reasonId, _ = strconv.Atoi(keyval[1])
			case "clid":
				clientMovedEv.cId, _ = strconv.Atoi(keyval[1])
			}
		}
	}
	clientMovedEv.api = &api
	for element := api.listenerList.Front(); element != nil; element = element.Next() {
		listener := element.Value.(TS3Listener)
		go listener.ClientMoved(clientMovedEv)
	}
}

func (api TS3Api) dispatchClientJoinMessage(msg string) {
	clientJoinEv := &ClientJoinEvent{}
	params := strings.Split(msg, " ")
	for _, message := range params {
		if strings.Contains(message, "=") {
			keyval := strings.SplitN(message, "=", 2)
			switch keyval[0] {
			case "clid":
				clientJoinEv.cId, _ = strconv.Atoi(keyval[1])
			}
		}
	}
	clientJoinEv.api = &api
	for element := api.listenerList.Front(); element != nil; element = element.Next() {
		listener := element.Value.(TS3Listener)
		go listener.ClientJoined(clientJoinEv)
	}
}

func (api TS3Api) dispatchClientLeaveMessage(msg string) {
	clientLeaveEv := &ClientLeaveEvent{}
	params := strings.Split(msg, " ")
	for _, message := range params {
		if strings.Contains(message, "=") {
			keyval := strings.SplitN(message, "=", 2)
			switch keyval[0] {
			case "cfid":
				clientLeaveEv.chFromId, _ = strconv.Atoi(keyval[1])
			case "clid":
				clientLeaveEv.clId, _ = strconv.Atoi(keyval[1])
			case "ctid":
				clientLeaveEv.chToId, _ = strconv.Atoi(keyval[1])
			case "reasonid":
				clientLeaveEv.reasonId, _ = strconv.Atoi(keyval[1])
			case "reasonmsg":
				clientLeaveEv.reasonMsg = keyval[1]
			}
		}
	}
	clientLeaveEv.api = &api
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
