// InvokerHolder
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
