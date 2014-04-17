// TextMessageEvent
package ts3api

import ()

type TextMessageEvent struct {
	Targetmode  int
	Msg         string
	InvokerId   int
	InvokerName string
	InvokerUid  string
	Api         *TS3Api
}
