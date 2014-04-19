// TS3ApiExportedMethods
package ts3api

import (
	"container/list"
	"errors"
	"strconv"
	"strings"
)

func New(network, addr string) (api *TS3Api, err error) {
	ts3conn, err := newConnection(network, addr)
	if err != nil {
		return
	}
	api = &TS3Api{
		conn:         ts3conn,
		lineList:     list.New(),
		listenerList: list.New(),
	}
	return
}

func (api TS3Api) Run(ch chan<- bool) {
	api.conn.ReadLine()
	api.conn.ReadLine()
	go api.reader(ch)
}

func (api TS3Api) RegisterTS3Listener(listener TS3Listener) {
	api.listenerList.PushBack(listener)
}

func (api TS3Api) Login(user, password string) {
	cmd := "login " + user + " " + password
	api.doCommand(cmd)
}

func (api TS3Api) Logout() {
	cmd := "logout"
	api.doCommand(cmd)
}

func (api TS3Api) Quit() {
	cmd := "quit"
	api.doCommand(cmd)
	api.conn.Close()
}

// id is ignored for every event except channel
// id = 0 for channel, stands for all channels
func (api TS3Api) RegisterEvent(event string, id int) {
	cmd := "servernotifyregister event=" + event
	if event == "channel" {
		api.doCommand(cmd + " id=" + strconv.Itoa(id))
	} else {
		api.doCommand(cmd)
	}
}

const (
	MESSAGE_TARGETMODE_CLIENT  = 1
	MESSAGE_TARGETMODE_CHANNEL = 2
	MESSAGE_TARGETMODE_SERVER  = 3
)

/*
CLIENT = 1 : target is a client
CHANNEL 2: target is a channel
SERVER 3: target is a virtual server
*/
func (api TS3Api) SendTextMessage(targetmode int, target int, msg string) (err error) {
	if targetmode < 1 || targetmode > 3 {
		err = errors.New("Targetmode out of range musst be > 1 and < 4")
		return
	}
	cmd := "sendtextmessage targetmode=" + strconv.Itoa(targetmode) + " target=" + strconv.Itoa(target) + " msg=" + msg
	api.doCommand(cmd)
	return
}

func (api TS3Api) SelectVirtualServer(serverid int) {
	cmd := "use " + strconv.Itoa(serverid)
	api.doCommand(cmd)
}

func (api TS3Api) WhoAmI() (client *Me) {
	cmd := "whoami"
	answer := api.doCommand(cmd)
	//"virtualserver_status=online virtualserver_id=1 virtualserver_unique_identifier=Ee9hKUn3SzddLH\/nzUeQxevRLNo= virtualserver_port=9987 client_id=1 client_channel_id=1 client_nickname=TS3TriviaBot\sfrom\s92.194.104.232:65235 client_database_id=32 client_login_name=TS3ToIRCBridge client_unique_identifier=puCS36nEiC9WiSN2Yvp5dlft7wY= client_origin_server_id=1"
	arr := strings.Split(answer, " ")
	client = &Me{}
	for index, element := range arr {
		prop := strings.SplitN(element, "=", 2)
		logger.Error("Index=" + strconv.Itoa(index) + " Key=" + prop[0] + " Value=" + prop[1])
		if prop[0] == "client_id" {
			cid, err := strconv.Atoi(prop[1])
			if err != nil {
				logger.Error(err.Error())
			} else {
				client.cId = cid
			}
		}
	}
	api.readLine()
	return
}

func (api TS3Api) ClientMove(clid int, cid int) {
	cmd := "clientmove clid=" + strconv.Itoa(clid) + " cid=" + strconv.Itoa(cid)
	api.doCommand(cmd)
}

func (api TS3Api) SetNick(nick string) {
	cmd := "clientupdate client_nickname=" + nick
	api.doCommand(cmd)
}
