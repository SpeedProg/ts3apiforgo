// me
package ts3api

import ()

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

func (me *Me) VirtualServerUniqueId() string {
	return me.vsUniqueId
}

func (me *Me) VirtualServerPort() int {
	return me.vsPort
}

func (me *Me) ChannelId() int {
	return me.cChannelId
}

func (me *Me) Nick() string {
	return me.cNick
}

func (me *Me) DatabaseId() int {
	return me.cDbId
}
func (me *Me) LoginName() string {
	return me.cLoginName
}
func (me *Me) UniqueId() string {
	return me.cUniqueId
}

func (me *Me) OriginServerId() int {
	return me.cOriginServerId
}
