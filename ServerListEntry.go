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
	"errors"
	"strconv"
	"strings"
)

// VSMachinId == -1 => no machin id
type ServerListEntry struct {
	VSId        int
	VSPort      int
	VSStatus    string
	VSClOnline  int
	VSQClOnline int
	VSMaxCl     int
	VSUTime     uint64
	VSName      string
	VSAutostart bool
	VSMachinId  int
	VSUID       string
}

func (entry *ServerListEntry) parseMsg(s string) {
	params := strings.Split(s, " ")
	for _, param := range params {
		if strings.Contains(param, "=") {
			kvpair := strings.SplitN(param, "=", 2)
			entry.parseParam(kvpair[0], kvpair[1])
		} else {
			entry.parseParam(param, "")
		}
	}
}

func (entry *ServerListEntry) parseParam(k, v string) (err error) {
	switch k {
	case "virtualserver_id":
		entry.VSId, err = strconv.Atoi(v)
	case "virtualserver_port":
		entry.VSPort, err = strconv.Atoi(v)
	case "virtualserver_status":
		entry.VSStatus = v
	case "virtualserver_clientsonline":
		entry.VSClOnline, err = strconv.Atoi(v)
	case "virtualserver_queryclientsonline":
		entry.VSQClOnline, err = strconv.Atoi(v)
	case "virtualserver_maxclients":
		entry.VSMaxCl, err = strconv.Atoi(v)
	case "virtualserver_uptime":
		entry.VSUTime, err = strconv.ParseUint(v, 10, 64)
	case "virtualserver_name":
		entry.VSName = v
	case "virtualserver_autostart":
		entry.VSAutostart, err = getBoolFromString(v)
	case "virtualserver_machine_id":
		if v == "" {
			entry.VSMachinId = -1
		} else {
			entry.VSMachinId, err = strconv.Atoi(v)
		}
	case "virtualserver_unique_identifier":
		entry.VSUID = v
	default:
		logger.Error("%s=%s is not a valid param of serverlist.", k, v)
		err = errors.New(k + "=" + v + " is not a valid param of serverlist")
	}
	return
}
