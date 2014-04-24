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
