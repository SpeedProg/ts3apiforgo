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
