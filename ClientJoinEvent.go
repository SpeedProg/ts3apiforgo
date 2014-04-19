// ClientJoinEvent
package ts3api

import (
	"strconv"
	"strings"
)

var _ Event = (*ClientJoinEvent)(nil)

type ClientJoinEvent struct {
	*ApiHolder
	clMetaData                   string // client_meta_data
	chFromId                     int
	chToId                       int
	reasonId                     int
	clUID                        string
	clNick                       string
	clInputMuted                 bool
	clOutputMuted                bool
	clOutputOnlyMuted            bool
	clInputHW                    int
	clOutputHW                   int
	clIsRecording                bool
	clDbId                       int
	clChGroupId                  int
	clServerGroups               []int // int,int...
	clAway                       bool
	clAwayMsg                    string //
	clType                       int    // 0 Voice, 1 Query
	clFlagAvatar                 string // some id like: 2daca766255626e1c3f4f66c61989b06
	clTalkPwr                    int
	clTalkReq                    bool
	clTalkReqMsg                 string
	clDesc                       string
	clIsTalker                   bool
	clIsPrioSpeaker              bool
	clUnreadMsgs                 int
	clNickPhonetic               string
	clNeededServerqueryViewPower int
	clIconId                     int
	clIsChCommander              bool
	clCountry                    string // DE|US ...
	clChGroupInheritedChanId     int
	clBadages                    string // key=value... no sure how multiple a seperated
	clId                         int
}

func NewClientJoinEvent() (event *ClientJoinEvent) {
	event = &ClientJoinEvent{}
	event.ApiHolder = &ApiHolder{}
	return event
}

func (event *ClientJoinEvent) setParam(key string, val string) (err error) {
	switch key {
	case "clid":
		event.clId, err = strconv.Atoi(val)
	case "cfid":
		event.chFromId, err = strconv.Atoi(val)
	case "ctid":
		event.chToId, err = strconv.Atoi(val)
	case "reasonid":
		event.reasonId, err = strconv.Atoi(val)
	case "client_unique_identifier":
		event.clUID = val
	case "client_nickname":
		event.clNick = val
	case "client_input_muted":
		if val == "0" {
			event.clInputMuted = false
		} else if val == "1" {
			event.clInputMuted = true
		} else {
			logger.Error("client_input_muted was: " + val)
			event.clInputMuted = false
		}
	case "client_output_muted":
		if val == "0" {
			event.clOutputMuted = false
		} else if val == "1" {
			event.clOutputMuted = true
		} else {
			logger.Error("client_output_muted was: " + val)
			event.clOutputMuted = false
		}
	case "client_outputonly_muted":
		if val == "0" {
			event.clOutputOnlyMuted = false
		} else if val == "1" {
			event.clOutputOnlyMuted = true
		} else {
			logger.Error("client_outputonly_muted was: " + val)
			event.clOutputOnlyMuted = false
		}
	case "client_input_hardware":
		event.clInputHW, err = strconv.Atoi(val)
	case "client_output_hardware":
		event.clOutputHW, err = strconv.Atoi(val)
	case "client_meta_data":
		event.clMetaData = val
	case "client_is_recording":
		if val == "0" {
			event.clIsRecording = false
		} else if val == "1" {
			event.clIsRecording = true
		} else {
			logger.Error("client_is_recording was: " + val)
			event.clIsRecording = false
		}
	case "client_database_id":
		event.clDbId, err = strconv.Atoi(val)
	case "client_channel_group_id":
		event.clChGroupId, err = strconv.Atoi(val)
	case "client_servergroups":
		sgroups := strings.Split(val, ",")
		event.clServerGroups = make([]int, len(sgroups))
		for idx, gids := range sgroups {
			event.clServerGroups[idx], err = strconv.Atoi(gids)
		}
	case "client_away":
		if val == "0" {
			event.clAway = false
		} else if val == "1" {
			event.clAway = true
		} else {
			logger.Error("client_away was: " + val)
			event.clAway = false
		}
	case "client_away_message":
		event.clAwayMsg = val
	case "client_type":
		event.clType, err = strconv.Atoi(val)
	case "client_flag_avatar":
		event.clFlagAvatar = val
	case "client_talk_power":
		event.clTalkPwr, err = strconv.Atoi(val)
	case "client_talk_request":
		valb, err := getBoolFromString(val)
		if err != nil {
			logger.Error("client_talk_request was: " + val)
			event.clTalkReq = false
		} else {
			event.clTalkReq = valb
		}
	case "client_talk_request_msg":
		event.clTalkReqMsg = val
	case "client_description":
		event.clDesc = val
	case "client_is_talker":
		valb, err := getBoolFromString(val)
		if err != nil {
			logger.Error("client_is_talker was: " + val)
			event.clIsTalker = false
		} else {
			event.clIsTalker = valb
		}
	case "client_is_priority_speaker":
		valb, err := getBoolFromString(val)
		if err != nil {
			logger.Error("client_is_priority_speaker was: " + val)
			event.clIsPrioSpeaker = false
		} else {
			event.clIsPrioSpeaker = valb
		}
	case "client_nickname_phonetic":
		event.clNickPhonetic = val
	case "client_needed_serverquery_view_power":
		event.clNeededServerqueryViewPower, err = strconv.Atoi(val)
	case "client_icon_id":
		event.clIconId, err = strconv.Atoi(val)
	case "client_is_channel_commander":
		valb, err := getBoolFromString(val)
		if err != nil {
			logger.Error("client_is_channel_commander was: " + val)
			event.clIsChCommander = false
		} else {
			event.clIsChCommander = valb
		}
	case "client_country":
		event.clCountry = val
	case "client_channel_group_inherited_channel_id":
		event.clChGroupInheritedChanId, err = strconv.Atoi(val)
	case "client_badges":
		event.clBadages = val
	case "client_unread_messages":
		event.clUnreadMsgs, err = strconv.Atoi(val)
	default:
		logger.Error("Key: " + key + " Value:" + val + " where not valid!")
	}
	return
}

func (event *ClientJoinEvent) Id() int {
	return event.clId
}

func (event *ClientJoinEvent) ChannelFromId() int {
	return event.chFromId
}

func (event *ClientJoinEvent) ChannelToId() int {
	return event.chToId
}

func (event *ClientJoinEvent) ReasonId() int {
	return event.reasonId
}

func (event *ClientJoinEvent) UID() string {
	return event.clUID
}

func (event *ClientJoinEvent) Nick() string {
	return event.clNick
}

func (event *ClientJoinEvent) InputMuted() bool {
	return event.clInputMuted
}

func (event *ClientJoinEvent) OutputMuted() bool {
	return event.clOutputMuted
}

func (event *ClientJoinEvent) InputHardware() int {
	return event.clInputHW
}

func (event *ClientJoinEvent) OutputHardware() int {
	return event.clOutputHW
}

func (event *ClientJoinEvent) OutputOnlyMuted() bool {
	return event.clOutputOnlyMuted
}

func (event *ClientJoinEvent) IsRecording() bool {
	return event.clIsRecording
}

func (event *ClientJoinEvent) DatabaseId() int {
	return event.clDbId
}

func (event *ClientJoinEvent) ChannelGroupId() int {
	return event.clChGroupId
}

func (event *ClientJoinEvent) ServerGroups() []int {
	return event.clServerGroups
}

func (event *ClientJoinEvent) Away() bool {
	return event.clAway
}

func (event *ClientJoinEvent) AwayMsg() string {
	return event.clAwayMsg
}

// 0 = Voice
// 1 = Query
func (event *ClientJoinEvent) Type() int {
	return event.clType
}

func (event *ClientJoinEvent) FlagAvatar() string {
	return event.clFlagAvatar
}

func (event *ClientJoinEvent) TalkPower() int {
	return event.clTalkPwr
}

func (event *ClientJoinEvent) TalkRequested() bool {
	return event.clTalkReq
}

func (event *ClientJoinEvent) TalkRequestMsg() string {
	return event.clTalkReqMsg
}

func (event *ClientJoinEvent) Description() string {
	return event.clDesc
}

func (event *ClientJoinEvent) IsTalker() bool {
	return event.clIsTalker
}

func (event *ClientJoinEvent) IsPrioritySpeaker() bool {
	return event.clIsPrioSpeaker
}

func (event *ClientJoinEvent) UnreadMsgs() int {
	return event.clUnreadMsgs
}

func (event *ClientJoinEvent) NickPhonetic() string {
	return event.clNickPhonetic
}

func (event *ClientJoinEvent) NeededServerqueryViewPower() int {
	return event.clNeededServerqueryViewPower
}

func (event *ClientJoinEvent) IconId() int {
	return event.clIconId
}

func (event *ClientJoinEvent) IsChannelCommander() bool {
	return event.clIsChCommander
}

func (event *ClientJoinEvent) Country() string {
	return event.clCountry
}

func (event *ClientJoinEvent) ChannelGroupInheritedChannelId() int {
	return event.clChGroupInheritedChanId
}

func (event *ClientJoinEvent) Badages() string {
	return event.clBadages
}
