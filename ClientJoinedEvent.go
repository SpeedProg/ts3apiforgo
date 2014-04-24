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

package ts3api

import (
	"strconv"
	"strings"
)

var _ Event = (*ClientJoinedEvent)(nil)

type ClientJoinedEvent struct {
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

func NewClientJoinedEvent() (event *ClientJoinedEvent) {
	event = &ClientJoinedEvent{}
	event.ApiHolder = &ApiHolder{}
	return event
}

func (event *ClientJoinedEvent) setParam(key string, val string) (err error) {
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

func (event *ClientJoinedEvent) Id() int {
	return event.clId
}

func (event *ClientJoinedEvent) ChannelFromId() int {
	return event.chFromId
}

func (event *ClientJoinedEvent) ChannelToId() int {
	return event.chToId
}

func (event *ClientJoinedEvent) ReasonId() int {
	return event.reasonId
}

func (event *ClientJoinedEvent) UID() string {
	return event.clUID
}

func (event *ClientJoinedEvent) Nick() string {
	return event.clNick
}

func (event *ClientJoinedEvent) InputMuted() bool {
	return event.clInputMuted
}

func (event *ClientJoinedEvent) OutputMuted() bool {
	return event.clOutputMuted
}

func (event *ClientJoinedEvent) InputHardware() int {
	return event.clInputHW
}

func (event *ClientJoinedEvent) OutputHardware() int {
	return event.clOutputHW
}

func (event *ClientJoinedEvent) OutputOnlyMuted() bool {
	return event.clOutputOnlyMuted
}

func (event *ClientJoinedEvent) IsRecording() bool {
	return event.clIsRecording
}

func (event *ClientJoinedEvent) DatabaseId() int {
	return event.clDbId
}

func (event *ClientJoinedEvent) ChannelGroupId() int {
	return event.clChGroupId
}

func (event *ClientJoinedEvent) ServerGroups() []int {
	return event.clServerGroups
}

func (event *ClientJoinedEvent) Away() bool {
	return event.clAway
}

func (event *ClientJoinedEvent) AwayMsg() string {
	return event.clAwayMsg
}

// 0 = Voice
// 1 = Query
func (event *ClientJoinedEvent) Type() int {
	return event.clType
}

func (event *ClientJoinedEvent) FlagAvatar() string {
	return event.clFlagAvatar
}

func (event *ClientJoinedEvent) TalkPower() int {
	return event.clTalkPwr
}

func (event *ClientJoinedEvent) TalkRequested() bool {
	return event.clTalkReq
}

func (event *ClientJoinedEvent) TalkRequestMsg() string {
	return event.clTalkReqMsg
}

func (event *ClientJoinedEvent) Description() string {
	return event.clDesc
}

func (event *ClientJoinedEvent) IsTalker() bool {
	return event.clIsTalker
}

func (event *ClientJoinedEvent) IsPrioritySpeaker() bool {
	return event.clIsPrioSpeaker
}

func (event *ClientJoinedEvent) UnreadMsgs() int {
	return event.clUnreadMsgs
}

func (event *ClientJoinedEvent) NickPhonetic() string {
	return event.clNickPhonetic
}

func (event *ClientJoinedEvent) NeededServerqueryViewPower() int {
	return event.clNeededServerqueryViewPower
}

func (event *ClientJoinedEvent) IconId() int {
	return event.clIconId
}

func (event *ClientJoinedEvent) IsChannelCommander() bool {
	return event.clIsChCommander
}

func (event *ClientJoinedEvent) Country() string {
	return event.clCountry
}

func (event *ClientJoinedEvent) ChannelGroupInheritedChannelId() int {
	return event.clChGroupInheritedChanId
}

func (event *ClientJoinedEvent) Badages() string {
	return event.clBadages
}
