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

var _ Event = (*TokenUsedEvent)(nil)

type TokenUsedEvent struct {
	*ApiHolder
	clId           int
	clDbId         int
	clUID          string
	token          string
	tokencustomset string
	token1         int
	token2         int
}

func NewTokenUsedEvent() (event *TokenUsedEvent) {
	event = &TokenUsedEvent{}
	event.ApiHolder = &ApiHolder{}
	return event
}

func (event *TokenUsedEvent) setParam(key, val string) (err error) {
	switch key {
	case "clid":
		event.clId, err = strconv.Atoi(val)
	case "cldbid":
		event.clDbId, err = strconv.Atoi(val)
	case "cluid":
		event.clUID = val
	case "token":
		event.token = val
	case "tokenconstomset":
		event.tokencustomset = val
	case "token1":
		event.token1, err = strconv.Atoi(val)
	case "token2":
		event.token2, err = strconv.Atoi(val)
	default:
		logger.Trace("%s=%s is not valid.", key, val)
		err = errors.New(key + "=" + val + " is not valid.")
	}
	return
}

func (event *TokenUsedEvent) ClientId() int {
	return event.clId
}

func (event *TokenUsedEvent) ClientDBId() int {
	return event.clDbId
}

func (event *TokenUsedEvent) ClientUID() string {
	return event.clUID
}

func (event *TokenUsedEvent) Token() string {
	return event.token
}

func (event *TokenUsedEvent) TokenGroup() int {
	return event.token1
}

func (event *TokenUsedEvent) TokenChannel() int {
	return event.token2
}

// 0 Server Group, 1 Channel Group
func (event *TokenUsedEvent) TokenType() int {
	if event.token2 == 0 {
		return 0
	}
	return 1
}
