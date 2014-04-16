// ts3api
package ts3api

import (
	"fmt"
	"io"
	"net"
	"net/textproto"
	"strconv"
	"strings"
)

type TS3Api struct {
	conn *textproto.Conn
}

func New(network, addr string) (ts3conn *TS3Api, err error) {
	nconn, err := net.Dial(network, addr)
	ts3conn = NewConn(nconn)
	return
}

func NewConn(nconn io.ReadWriteCloser) (conn *TS3Api) {
	conn = &TS3Api{
		conn: textproto.NewConn(nconn),
	}
	conn.readLineToStdo()
	conn.readLineToStdo()
	return
}

func (conn *TS3Api) doCommand(cmd string) (answer string) {
	fmt.Println("-->" + cmd + "<--")
	//conn.conn.PrintfLine(cmd)
	conn.conn.Writer.W.WriteString(cmd + "\n")
	conn.conn.Writer.W.Flush()
	answer = conn.readLineToStdo()
	return
}

func (conn *TS3Api) readLineToStdo() (msg string) {
	msg, err := conn.conn.ReadLine()
	msg = strings.TrimSpace(msg)
	if err != nil {
		fmt.Print("Error Happended!")
	} else {
		fmt.Println("-->" + msg + "<--")
	}
	return
}

func (conn *TS3Api) Login(user, password string) {
	cmd := "login " + user + " " + password
	conn.doCommand(cmd)
}

func (conn *TS3Api) Logout() {
	cmd := "logout"
	conn.doCommand(cmd)
}

func (conn *TS3Api) Quit() {
	cmd := "quit"
	conn.doCommand(cmd)
	conn.conn.Close()
}

/*
enum TextMessageTargetMode {
TextMessageTarget_CLIENT
= 1,
// 1:
target is a client
TextMessageTarget_CHANNEL,
// 2: target is a channel
TextMessageTarget_SERVER
// 3: target is a virtual server
};
*/
func (conn *TS3Api) SendTextMessage(targetmode int, target int, msg string) {
	cmd := "sendtextmessage targetmode=" + strconv.Itoa(targetmode) + " target=" + strconv.Itoa(target) + " msg=" + msg
	conn.doCommand(cmd)
}

func (conn *TS3Api) SelectVirtualServer(serverid int) {
	cmd := "use " + strconv.Itoa(serverid)
	conn.doCommand(cmd)
}

func (conn *TS3Api) WhoAmI() (client *Me) {
	cmd := "whoami"
	answer := conn.doCommand(cmd)
	//"virtualserver_status=online virtualserver_id=1 virtualserver_unique_identifier=Ee9hKUn3SzddLH\/nzUeQxevRLNo= virtualserver_port=9987 client_id=1 client_channel_id=1 client_nickname=TS3TriviaBot\sfrom\s92.194.104.232:65235 client_database_id=32 client_login_name=TS3ToIRCBridge client_unique_identifier=puCS36nEiC9WiSN2Yvp5dlft7wY= client_origin_server_id=1"
	arr := strings.Split(answer, " ")
	client = &Me{}
	for index, element := range arr {
		prop := strings.SplitN(element, "=", 2)
		fmt.Println("Index=" + strconv.Itoa(index) + " Key=" + prop[0] + " Value=" + prop[1])
		if prop[0] == "client_id" {
			cid, err := strconv.Atoi(prop[1])
			if err != nil {
				fmt.Println(err.Error())
			} else {
				client.cId = cid
			}
		}
	}
	conn.readLineToStdo()
	return
}

func (conn *TS3Api) ClientMove(clid int, cid int) {
	cmd := "clientmove clid=" + strconv.Itoa(clid) + " cid=" + strconv.Itoa(cid)
	conn.doCommand(cmd)
}

func (conn *TS3Api) SetNick(nick string) {
	cmd := "clientupdate client_nickname=" + nick
	conn.doCommand(cmd)
}

type Me struct {
	vsId            int
	vsUniqueId      string
	vsPort          int
	cId             int
	cChannelId      int
	cNick           string
	cDbId           int
	cLoginName      string
	cUniqueId       string
	cOriginServerId int
}

func (me *Me) VirtualServerId() int {
	return me.vsId
}

func (me *Me) ClientId() int {
	return me.cId
}
