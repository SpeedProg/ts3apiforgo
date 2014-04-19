// Event
package ts3api

type Event interface {
	setParam(key, val string) (err error)
	setApi(api *TS3Api)
	Api() *TS3Api
}
