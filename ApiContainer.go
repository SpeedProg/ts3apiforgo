// ApiContainer
package ts3api

type ApiContainer interface {
	setApi(api *TS3Api)
	Api() *TS3Api
}
