package main

import (
	"code.google.com/p/log4go"
	"flag"
	"fmt"
	"speedprog.de/ts3api"
	"speedprog.de/ts3api/ts3const"
	"strconv"
)

const APP_VERSION = "0.1"

var lName string = "myloginname"
var lPass string = "myloginpass"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

var logger log4go.Logger = log4go.NewDefaultLogger(log4go.DEBUG)

func init() {
	logger.LoadConfiguration("ts3bot.xml")
}

var _ ts3api.TS3Listener = (*TS3Obs)(nil)

type TS3Obs struct {
	*ts3api.TS3Adapter
	api *ts3api.TS3Api
}

func (obs *TS3Obs) TextMessage(msg *ts3api.TextMessageEvent) {
	fmt.Println(msg.Msg)
	if msg.Msg() == "ping" {
		msg.Api().SendTextMessage(msg.Targetmode(), 1, "pong")
	} else if msg.Msg() == "!shutdown" {
		obs.api.Logout()
		obs.api.Quit()
	} else {
		logger.Trace(msg.Msg())
	}
}
func (obs *TS3Obs) ClientJoined(event *ts3api.ClientJoinedEvent) {
	fmt.Println("Client With ID: " + strconv.Itoa(event.Id()))
}
func (obs *TS3Obs) ClientMoved(event *ts3api.ClientMovedEvent) {
	ids := event.Ids()
	for _, id := range ids {
		fmt.Println("Client Moved With ID: " + strconv.Itoa(id))
	}
}

func (obs *TS3Obs) ServerEdited(event *ts3api.ServerEditedEvent) {
	fmt.Println("Server Edited: " + event.Values())
}

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}
	var client *ts3api.Me

	var conn *ts3api.TS3Api
	var err error
	ch := make(chan bool)
	conn, err = ts3api.New("hostname:10011", ch)
	if err != nil {
		fmt.Println("Error creating Connection!")
		fmt.Println(err.Error())
	}
	var obs = &TS3Obs{&ts3api.TS3Adapter{}, conn}
	/*
		for i := 0; i < 20; i++ {
			go connectToServer()
		}
	*/
	conn.Login(lName, lPass)
	conn.SelectVirtualServer(1)
	conn.SetNick("TestClient")
	client = conn.WhoAmI()
	conn.ClientMove(client.ClientId(), 1)

	conn.RegisterEvent("textchannel", 0)
	conn.RegisterEvent("textserver", 0)
	conn.RegisterEvent("server", 0)
	conn.RegisterEvent("channel", 0)
	conn.RegisterEvent("textprivate", 0)
	conn.RegisterEvent("tokenused", 0)
	conn.RegisterTS3Listener(obs)
	version, build, platform := conn.Version()
	logger.Trace("%s %d %s", version, build, platform)
	hostinfo, _, _ := conn.Hostinfo()
	logger.Trace("Hostinfo: %+v", hostinfo)
	instanceinfo, _, _ := conn.Instanceinfo()
	logger.Trace("Instanceinfo: %+v", instanceinfo)
	conn.SendTextMessage(2, 1, "Hallo")
	conn.SendTextMessage(2, 1, "Das muss man mal versuche ZUB\n+BELL\b+WHat ever a is \a.   VOID STUFF////\\\\\\\\")
	props := make([][]string, 1)
	props[0] = make([]string, 2)
	props[0][0] = ts3const.SERVERINSTANCE_SERVERQUERY_FLOOD_COMMANDS
	props[0][1] = "1500"
	conn.Instanceedit(props)
	instanceinfo, _, _ = conn.Instanceinfo()
	logger.Trace("Instanceinfo: %+v", instanceinfo)
	<-ch
}

func connectToServer() {
	var conn *ts3api.TS3Api
	var client *ts3api.Me
	var err error
	ch := make(chan bool)
	conn, err = ts3api.New("hostname:10011", ch)
	if err != nil {
		fmt.Println("Error creating Connection!")
		fmt.Println(err.Error())
	}
	conn.Login(lName, lPass)
	conn.SelectVirtualServer(1)
	conn.SetNick("TestClient")
	client = conn.WhoAmI()
	conn.ClientMove(client.ClientId(), 1)
	conn.SendTextMessage(2, 1, "Hallo")
	conn.RegisterEvent("textchannel", 0)
	conn.RegisterEvent("textserver", 0)
	conn.RegisterEvent("server", 0)
	conn.RegisterEvent("channel", 1)
	conn.RegisterEvent("textprivate", 0)
	<-ch
}
