// Event
package ts3api

type Event interface {
	setParam(key, val string) (err error)
	ApiContainer
}
