package ts3api

import (
	"strconv"
)

var _ Event = (*ChannelCreatedEvent)(nil)

type ChannelCreatedEvent struct {
	*ApiHolder
	*InvokerHolder
	chId                        int
	chParentId                  int
	chName                      string
	chTopic                     string
	chCodec                     int
	chCodecQuality              int
	chMaxClients                int
	chMaxFamilyClients          int
	chOrder                     int
	chPermanent                 bool
	chSemiPermanent             bool
	chDefault                   bool
	chPassword                  bool
	chCodecLatFactor            int
	chCodecIsUnencrypted        bool
	chDeleteDelay               int
	chMaxClientsUnlimited       bool
	chMaxFamilyclientsUnlimited bool
	chMaxFamilyclientsInherited bool
	chNeededTalkPower           int
	chNamePhonectic             string
	chIconId                    string
}

func NewChannelCreatedEvent() (event *ChannelCreatedEvent) {
	event = &ChannelCreatedEvent{}
	event.ApiHolder = &ApiHolder{}
	event.InvokerHolder = &InvokerHolder{}
	return
}

func (event *ChannelCreatedEvent) setParam(key, val string) (err error) {
	switch key {
	case "cid":
		event.chId, err = strconv.Atoi(val)
	case "cpid":
		event.chParentId, err = strconv.Atoi(val)
	case "channel_name":
		event.chName = val
	case "channel_topic":
		event.chTopic = val
	case "channel_codec":
		event.chCodec, err = strconv.Atoi(val)
	case "channel_codec_quality":
		event.chCodecQuality, err = strconv.Atoi(val)
	case "channel_maxclients":
		event.chMaxClients, err = strconv.Atoi(val)
	case "channel_maxfamilyclients":
		event.chMaxFamilyClients, err = strconv.Atoi(val)
	case "channel_order":
		event.chOrder, err = strconv.Atoi(val)
	case "channel_flag_permanent":
		event.chPermanent, err = getBoolFromString(val)
	case "channel_flag_semi_permanent":
		event.chSemiPermanent, err = getBoolFromString(val)
	case "channel_flag_default":
		event.chDefault, err = getBoolFromString(val)
	case "channel_flag_password":
		event.chPassword, err = getBoolFromString(val)
	case "channel_codec_latency_factor":
		event.chCodecLatFactor, err = strconv.Atoi(val)
	case "channel_codec_is_unencrypted":
		event.chCodecIsUnencrypted, err = getBoolFromString(val)
	case "channel_delete_delay":
		event.chDeleteDelay, err = strconv.Atoi(val)
	case "channel_flag_maxclients_unlimited":
		event.chMaxClientsUnlimited, err = getBoolFromString(val)
	case "channel_flag_maxfamilyclients_unlimited":
		event.chMaxFamilyclientsUnlimited, err = getBoolFromString(val)
	case "channel_flag_maxfamilyclients_inherited":
		event.chMaxFamilyclientsInherited, err = getBoolFromString(val)
	case "channel_needed_talk_power":
		event.chNeededTalkPower, err = strconv.Atoi(val)
	case "channel_name_phonetic":
		event.chNamePhonectic = val
	case "channel_icon_id":
		event.chNamePhonectic = val
	default:
		err = event.InvokerHolder.setParam(key, val)
	}
	return
}

func (event *ChannelCreatedEvent) ChannelId() int {
	return event.chId
}

func (event *ChannelCreatedEvent) ParentId() int {
	return event.chParentId
}

func (event *ChannelCreatedEvent) Name() string {
	return event.chName
}

func (event *ChannelCreatedEvent) Topic() string {
	return event.chTopic
}

func (event *ChannelCreatedEvent) Codec() int {
	return event.chCodec
}

func (event *ChannelCreatedEvent) CodecQuality() int {
	return event.chCodecQuality
}

func (event *ChannelCreatedEvent) MaxClients() int {
	return event.chMaxClients
}

func (event *ChannelCreatedEvent) MaxFamilyClients() int {
	return event.chMaxFamilyClients
}

func (event *ChannelCreatedEvent) Order() int {
	return event.chOrder
}

func (event *ChannelCreatedEvent) Permanent() bool {
	return event.chPermanent
}

func (event *ChannelCreatedEvent) SemiPermanent() bool {
	return event.chSemiPermanent
}
func (event *ChannelCreatedEvent) Default() bool {
	return event.chDefault
}
func (event *ChannelCreatedEvent) Password() bool {
	return event.chPassword
}
func (event *ChannelCreatedEvent) CodecLatFactor() int {
	return event.chCodecLatFactor
}

func (event *ChannelCreatedEvent) CodecIsUnencrypted() bool {
	return event.chCodecIsUnencrypted
}

func (event *ChannelCreatedEvent) DeleteDelay() int {
	return event.chDeleteDelay
}

func (event *ChannelCreatedEvent) MaxClientsUnlimited() bool {
	return event.chMaxClientsUnlimited
}
func (event *ChannelCreatedEvent) MaxFamilyclientsUnlimited() bool {
	return event.chMaxFamilyclientsUnlimited
}
func (event *ChannelCreatedEvent) MaxFamilyClientsInherited() bool {
	return event.chMaxFamilyclientsInherited
}
func (event *ChannelCreatedEvent) NeededTalkPower() int {
	return event.chNeededTalkPower
}
func (event *ChannelCreatedEvent) NamePhonectic() string {
	return event.chNamePhonectic
}
func (event *ChannelCreatedEvent) IconId() string {
	return event.chIconId
}
