// TS3ApiExportedMethods
package ts3api

import (
	"container/list"
	"errors"
	"strconv"
	"strings"
)

// New creates a TS3Api, connecting to the given address.
// addr should looke like 127.0.0.1:10011.
// Starts the TS3 Query connection in an other routine.
// To wait for its ending you can use the given channel.
func New(addr string, ch chan<- bool) (api *TS3Api, err error) {
	ts3conn, err := newConnection("tcp", addr)
	if err != nil {
		return
	}
	api = &TS3Api{
		conn:         ts3conn,
		lineList:     list.New(),
		listenerList: list.New(),
	}
	api.conn.ReadLine()
	api.conn.ReadLine()
	go api.reader(ch)
	return
}

// Register a TS3Listener.
func (api TS3Api) RegisterTS3Listener(listener TS3Listener) {
	api.listenerList.PushBack(listener)
}

// Login as user with password.
func (api TS3Api) Login(user, password string) {
	cmd := "login " + user + " " + password
	api.doCommand(cmd)
}

// Logout.
// Logging out does not end the connection, you can login again afterwards.
func (api TS3Api) Logout() {
	cmd := "logout"
	api.doCommand(cmd)
}

// Send quit over the query connection.
// This causes the ts queryserver to end the connection.
// After using this you can not use this TS3Api object anymore.
func (api TS3Api) Quit() {
	cmd := "quit"
	api.doCommand(cmd)
	api.conn.Close()
}

// id is ignored for every event except channel
// id = 0 for channel, stands for all channels
// Events are: tokenused, textserver, textchannel, textprivate, channel, server
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

// Send a textmessage to the given targetmode and target.
//
// CLIENT = 1 : target is a client
// CHANNEL 2: target is a channel
// SERVER 3: target is a virtual server
func (api TS3Api) SendTextMessage(targetmode int, target int, msg string) (err error) {
	if targetmode < 1 || targetmode > 3 {
		err = errors.New("Targetmode out of range musst be > 1 and < 4")
		return
	}
	cmd := "sendtextmessage targetmode=" + strconv.Itoa(targetmode) + " target=" + strconv.Itoa(target) + " msg=" + msg
	api.doCommand(cmd)
	return
}

// Select a virtualserver by id.
func (api TS3Api) SelectVirtualServer(serverid int) {
	cmd := "use " + strconv.Itoa(serverid)
	api.doCommand(cmd)
}

// Get informations about your self, like your id.
func (api TS3Api) WhoAmI() (client *Me) {
	cmd := "whoami"
	answers, _ := api.doCommand(cmd)
	// TODO: error handling
	arr := strings.Split(answers.Front().Value.(string), " ")
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
	return
}

// Move client with id clid to channel with id cid.
func (api TS3Api) ClientMove(clid int, cid int) {
	cmd := "clientmove clid=" + strconv.Itoa(clid) + " cid=" + strconv.Itoa(cid)
	api.doCommand(cmd)
}

// Set your own nick.
func (api TS3Api) SetNick(nick string) {
	cmd := "clientupdate client_nickname=" + nick
	api.doCommand(cmd)
}

func (api TS3Api) Version() (version string, build uint64, platform string) {
	answers, _ := api.doCommand("version")
	var answer string = answers.Front().Value.(string)
	parts := strings.Split(answer, " ")
	var err error
	for i := 0; i < 3; i++ {
		var tparts []string = strings.SplitN(parts[i], "=", 2)
		switch i {
		case 0:
			version = tparts[1]
		case 1:
			build, err = strconv.ParseUint(tparts[1], 10, 64)
			if err != nil {
				logger.Error(err.Error())
			}
		case 2:
			platform = tparts[1]
		}
	}
	return
}

