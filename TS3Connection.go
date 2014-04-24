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

// TS3Connection
package ts3api

import (
	"net"
	"net/textproto"
)

type ts3Connection struct {
	textprotoConn *textproto.Conn
}

func newConnection(network, addr string) (conn *ts3Connection, err error) {
	nconn, err := net.Dial(network, addr)
	if err != nil {
		return
	}
	conn = &ts3Connection{
		textprotoConn: textproto.NewConn(nconn),
	}
	return
}

func (conn ts3Connection) ReadString(b byte) (string, error) {
	msg, err := conn.textprotoConn.R.ReadString(b)
	return msg, err
}
func (conn ts3Connection) ReadLine() (line string, err error) {
	line, err = conn.textprotoConn.ReadLine()
	return
}
func (conn ts3Connection) Close() {
	conn.textprotoConn.Close()
}
func (conn ts3Connection) DoCommand(cmd string) {
	logger.Trace(cmd)
	conn.textprotoConn.W.WriteString(cmd + "\n")
	conn.textprotoConn.W.Flush()
	return
}
