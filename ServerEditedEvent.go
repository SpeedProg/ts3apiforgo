package ts3api

import ()

type ServerEditedEvent struct {
	reasonId    int
	invokerId   int
	invokerName string
	invokerUid  string
}
