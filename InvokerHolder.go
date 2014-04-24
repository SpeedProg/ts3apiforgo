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
)

type InvokerHolder struct {
	invokerId   int
	invokerName string
	invokerUID  string
}

func (event *InvokerHolder) setParam(key, val string) (err error) {
	switch key {
	case "invokerid":
		event.invokerId, err = strconv.Atoi(val)
	case "invokername":
		event.invokerName = val
	case "invokeruid":
		event.invokerUID = val
	default:
		logger.Trace("%s=%s is not valid!", key, val)
		err = errors.New(key + "=" + val + " is not valid!")
	}
	return
}

func (event *InvokerHolder) InvokerId() int {
	return event.invokerId
}

func (event *InvokerHolder) InvokerName() string {
	return event.invokerName
}

func (event *InvokerHolder) InvokerUID() string {
	return event.invokerUID
}
