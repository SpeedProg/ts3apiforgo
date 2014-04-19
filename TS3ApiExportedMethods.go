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
func (api TS3Api) RegisterEvent(event string, id int) (err error) {
	if event != "tokenused" && event != "textserver" && event != "textchannel" && event != "textprivate" && event != "channel" && event != "server" {
		err = errors.New("Event type " + event + " is not valid!")
		return
	}
	cmd := "servernotifyregister event=" + event
	if event == "channel" {
		api.doCommand(cmd + " id=" + strconv.Itoa(id))
	} else {
		api.doCommand(cmd)
	}
	return
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
	cmd := "sendtextmessage targetmode=" + strconv.Itoa(targetmode) +
		" target=" + strconv.Itoa(target) + " msg=" + encodeValue(msg)
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
		prop[1] = decodeValue(prop[1])
		logger.Trace("Index=" + strconv.Itoa(index) + " Key=" + prop[0] + " Value=" + prop[1])
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
	cmd := "clientupdate client_nickname=" + encodeValue(nick)
	api.doCommand(cmd)
}

func (api TS3Api) Version() (version string, build uint64, platform string) {
	answers, _ := api.doCommand("version")
	var answer string = answers.Front().Value.(string)
	parts := strings.Split(answer, " ")
	var err error
	for i := 0; i < 3; i++ {
		var tparts []string = strings.SplitN(parts[i], "=", 2)
		tparts[1] = decodeValue(tparts[1])
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

type HostInfo struct {
	Uptime                           uint64
	TimestampUTC                     uint64
	VirtualserverCount               uint
	VirtualserverTotalMaxClients     uint
	VirtualserverTotalClientsOnline  uint
	VirtualserverTotalChannelsOnline uint
	FiletransferBandwidthSent        uint64
	FiletransferBandwidthReceived    uint64
	FiletransferBytesSentTotal       uint64
	FiletransferBytesReceivedTotal   uint64
	PacketsSentTotal                 uint64
	BytesSentTotal                   uint64
	PacketsReceivedTotal             uint64
	BytesReceivedTotal               uint64
	BandwidthSendLastSecond          uint64
	BandwidthSendLastMinute          uint64
	BandwidthReceivedLastSecond      uint64
	BandwidthReceivedLastMinute      uint64
}

func (api TS3Api) Hostinfo() (info *HostInfo, qerr QueryError, err error) {
	info = &HostInfo{}
	answerList, qerr := api.doCommand("hostinfo")
	// TODO: handle errors
	answer := answerList.Front().Value.(string)
	params := strings.Split(answer, " ")
	for _, param := range params {
		func(p string) {
			var ival uint64
			parts := strings.SplitN(p, "=", 2)
			parts[1] = decodeValue(parts[1])
			switch parts[0] {
			case "instance_uptime":
				info.Uptime, err = strconv.ParseUint(parts[1], 10, 64)
			case "host_timestamp_utc":
				info.TimestampUTC, err = strconv.ParseUint(parts[1], 10, 64)
			case "virtualservers_running_total":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.VirtualserverCount = uint(ival)
			case "virtualservers_total_maxclients":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.VirtualserverTotalMaxClients = uint(ival)
			case "virtualservers_total_clients_online":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.VirtualserverTotalClientsOnline = uint(ival)
			case "virtualservers_total_channels_online":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.VirtualserverTotalChannelsOnline = uint(ival)
			case "connection_filetransfer_bandwidth_sent":
				info.FiletransferBandwidthSent, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_filetransfer_bandwidth_received":
				info.FiletransferBandwidthReceived, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_filetransfer_bytes_sent_total":
				info.FiletransferBytesSentTotal, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_filetransfer_bytes_received_total":
				info.FiletransferBytesReceivedTotal, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_packets_received_total":
				info.PacketsReceivedTotal, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_packets_sent_total":
				info.PacketsSentTotal, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_bytes_sent_total":
				info.BytesSentTotal, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_bytes_received_total":
				info.BytesReceivedTotal, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_bandwidth_sent_last_second_total":
				info.BandwidthSendLastSecond, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_bandwidth_sent_last_minute_total":
				info.BandwidthSendLastMinute, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_bandwidth_received_last_second_total":
				info.BandwidthReceivedLastSecond, err = strconv.ParseUint(parts[1], 10, 64)
			case "connection_bandwidth_received_last_minute_total":
				info.BandwidthReceivedLastMinute, err = strconv.ParseUint(parts[1], 10, 64)
			default:
				logger.Error("%s=%s is invalid for hostinfo.", parts[0], parts[1])
				err = errors.New(parts[0] + "=" + parts[1] + " is invalid for hostinfo.")
			}
		}(param)
	}
	return
}

type InstanceInfo struct {
	DBVersion             uint
	FtPort                uint
	MaxDlBandwidth        uint64
	MaxUpTotalBandwidth   uint64
	GuestServerQueryGroup uint
	QueryFloodCommands    uint
	QueryFloodTime        uint
	QueryBanTime          uint
	TmplSvrAdminGroup     uint
	TmplSvrDefaultGroup   uint
	TmplChAdminGroup      uint
	TmplChDefaultGroup    uint
	PermVersion           uint
	PendingConPerIP       uint
}

func (api TS3Api) Instanceinfo() (info *InstanceInfo, qerr QueryError, err error) {
	/*
		serverinstance_pending_connections_per_ip=0
	*/
	info = &InstanceInfo{}
	answerList, qerr := api.doCommand("instanceinfo")
	// TODO: handle errors
	answer := answerList.Front().Value.(string)
	params := strings.Split(answer, " ")
	for _, param := range params {
		func(p string) {
			var ival uint64
			parts := strings.SplitN(p, "=", 2)
			parts[1] = decodeValue(parts[1])
			switch parts[0] {
			case "serverinstance_database_version":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.DBVersion = uint(ival)
			case "serverinstance_filetransfer_port":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.FtPort = uint(ival)
			case "serverinstance_max_download_total_bandwidth":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.MaxDlBandwidth = ival
			case "serverinstance_max_upload_total_bandwidth":
				ival, err = strconv.ParseUint(parts[1], 10, 64)
				info.MaxUpTotalBandwidth = ival
			case "serverinstance_guest_serverquery_group":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.GuestServerQueryGroup = uint(ival)
			case "serverinstance_serverquery_flood_commands":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.QueryFloodCommands = uint(ival)
			case "serverinstance_serverquery_flood_time":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.QueryFloodTime = uint(ival)
			case "serverinstance_serverquery_ban_time":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.QueryBanTime = uint(ival)
			case "serverinstance_template_serveradmin_group":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.TmplSvrAdminGroup = uint(ival)
			case "serverinstance_template_serverdefault_group":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.TmplSvrDefaultGroup = uint(ival)
			case "serverinstance_template_channeladmin_group":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.TmplChAdminGroup = uint(ival)
			case "serverinstance_template_channeldefault_group":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.TmplChDefaultGroup = uint(ival)
			case "serverinstance_permissions_version":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.PermVersion = uint(ival)
			case "serverinstance_pending_connections_per_ip":
				ival, err = strconv.ParseUint(parts[1], 10, 32)
				info.PendingConPerIP = uint(ival)
			default:
				logger.Error("%s=%s is invalid for instanceinfo.", parts[0], parts[1])
				err = errors.New(parts[0] + "=" + parts[1] + " is invalid for instanceinfo.")
			}
		}(param)
	}
	return
}

// Takes an array of 2 dimensinal arrays representing Serverinstance properties
// that are editable and the value to set.
// Valid Properties are found as constants starting with SERVERINSTANCE_ in ts3const.
func (api TS3Api) Instanceedit(properties [][]string) (qerr QueryError) {
	cmd := "instanceedit"
	props := ""
	for _, pel := range properties {
		props += " " + strings.ToLower(pel[0]) + "=" + encodeValue(pel[1])
	}
	cmd += props
	_, qerr = api.doCommand(cmd)
	return
}

type ServerListEntry struct {
	VSId        int
	VSPort      int
	VSStatus    string
	VSClOnline  int
	VSQClOnline int
	VSMaxCl     int
	VSUTime     int64
	VSName      string
	VSAutostrt  bool
	VSMachinId  int
	VSUID       string
}

